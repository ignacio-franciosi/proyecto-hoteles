package user

import (
	"time"
	"uba/model"

	"github.com/karlseguin/ccache/v3"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Db *gorm.DB
var cache *ccache.Cache[model.User]

func init() {
	cache = ccache.New(ccache.Configure[model.User]())
}

func GetUserById(id int) model.User {
	var user model.User

	Db.Where("id_user = ?", id).First(&user)
	log.Debug("User: ", user)

	return user
}

func GetUserByEmail(email string) model.User {
	//intento agarrar del cache con un get y la clave (email)
	item := cache.Get(email)
	if item != nil {
		//si encontro algo extiende el tiempo de vida a 24hs mas
		item.Extend(24 * time.Hour)
		log.Info("Traído del cache :)")
		return item.Value()
	}

	var user model.User

	result := Db.Where("email = ?", email).First(&user)
	log.Debug("User: ", user)

	if result.Error == nil {
		//Si la consulta a la base de datos es exitosa, el usuario recuperado se almacena en el caché
		cache.Set(email, user, 24*time.Hour)
	}

	return user
}

func InsertUser(user model.User) model.User {

	result := Db.Create(&user)

	if result.Error != nil {
		//TO DO Manage Errors
		log.Error("Couldn't create user")
		return model.User{}
	}
	log.Debug("User Created: ", user.IdUser)
	//Si la inserción es exitosa, el usuario recién creado se almacena
	//en la caché con su correo electrónico como clave y un tiempo de vida de 24 horas

	cache.Set(user.Email, user, 24*time.Hour)

	return user
}
