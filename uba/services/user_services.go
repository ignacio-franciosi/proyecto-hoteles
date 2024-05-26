package services

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	userClient "uba/clients/user"
	dto2 "uba/dto"
	"uba/model"

	"github.com/dgrijalva/jwt-go"

	log "github.com/sirupsen/logrus"
)

type userService struct{}

type userServiceInterface interface {
	GetUserById(id int) (dto2.UserDto, error)
	LoginUser(loginDto dto2.LoginDto) (dto2.TokenDto, error)
	InsertUser(userDto dto2.UserDto) (dto2.TokenDto, error)
	GetUserByEmail(email string) (dto2.UserDto, error)
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (s *userService) GetUserById(id int) (dto2.UserDto, error) {

	var user model.User = userClient.GetUserById(id)
	var userDto dto2.UserDto

	if user.Id == 0 {
		return userDto, errors.New("usuario no encontrado")
	}
	userDto.Id = user.Id
	userDto.Name = user.Name
	userDto.LastName = user.LastName
	userDto.Email = user.Email
	userDto.Password = user.Password
	userDto.UserType = user.UserType
	userDto.Dni = user.Dni

	return userDto, nil
}

func (s *userService) GetUserByEmail(email string) (dto2.UserDto, error) {

	var user model.User = userClient.GetUserByEmail(email)
	var userDto dto2.UserDto

	if user.Email == "" {
		return userDto, errors.New("usuario no encontrado")
	}
	userDto.Id = user.Id
	userDto.Name = user.Name
	userDto.LastName = user.LastName
	userDto.Email = user.Email
	userDto.Password = user.Password
	userDto.UserType = user.UserType
	userDto.Dni = user.Dni

	return userDto, nil
}

//login

var jwtKey = []byte("secret_key")

func (s *userService) LoginUser(loginDto dto2.LoginDto) (dto2.TokenDto, error) {

	log.Debug(loginDto) //para registrar el contenido de loginDto
	var user model.User = userClient.GetUserByEmail(loginDto.Email)
	var tokenDto dto2.TokenDto

	if user.Id == 0 {
		return tokenDto, errors.New("user not found")
	}

	//pasamos password como slice de bytes
	//hashea con md5.sum
	var pswMd5 = md5.Sum([]byte(loginDto.Password))
	//convertir a cadena hexadecimal
	pswMd5String := hex.EncodeToString(pswMd5[:])

	if pswMd5String == user.Password {
		//se firma el token para verificar autenticidad
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id_user": user.Id,
		})
		tokenString, _ := token.SignedString(jwtKey)
		tokenDto.Token = tokenString
		tokenDto.IdUser = user.Id
		tokenDto.UserType = user.UserType

		return tokenDto, nil
	} else {
		return tokenDto, errors.New("contraseña incorrecta")
	}

}

func (s *userService) InsertUser(userDto dto2.UserDto) (dto2.TokenDto, error) {
	log.Debug(userDto) // Para registrar el contenido de userDto

	var user model.User
	var tokenDto dto2.TokenDto

	if user.Id == 0 { // El usuario no está registrado y puedo crear uno nuevo
		// Pasamos la contraseña como slice de bytes
		// Hash con md5.Sum
		var pswMd5 = md5.Sum([]byte(userDto.Password))
		// Convertir a cadena hexadecimal
		pswMd5String := hex.EncodeToString(pswMd5[:])

		// Asignamos valores al usuario antes de generas el token
		user.Id = userDto.Id
		user.Name = userDto.Name
		user.LastName = userDto.LastName
		user.Email = userDto.Email
		user.Password = pswMd5String
		user.UserType = userDto.UserType
		user.Dni = userDto.Dni

		// Insertamos el usuario en la base de datos
		user = userClient.InsertUser(user)

		// Ahora, después de asignar el ID del usuario, generamos el token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id_user": user.Id,
			"tipo":    user.UserType,
		})

		// Firmamos el token
		tokenString, _ := token.SignedString(jwtKey)
		tokenDto.Token = tokenString
		tokenDto.IdUser = user.Id
		tokenDto.UserType = user.UserType

		return tokenDto, nil

	} else { // El usuario ya existe
		return tokenDto, errors.New("usuario ya existe")
	}
}

/*package services

import (
	"crypto/md5"
	"encoding/hex"
	userClient "uba/clients/user"
	dto2 "uba/dto"
	"uba/model"
	e "uba/utils"

	"github.com/dgrijalva/jwt-go"

	log "github.com/sirupsen/logrus"
)

type userService struct{}

type userServiceInterface interface {
	GetUserById(id int) (dto2.UserDto, e.ApiError)
	LoginUser(loginDto dto2.LoginDto) (dto2.TokenDto, e.ApiError)
	InsertUser(userDto dto2.UserDto) (dto2.TokenDto, e.ApiError)
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (s *userService) GetUserById(id int) (dto2.UserDto, e.ApiError) {

	var user model.User = userClient.GetUserById(id)
	var userDto dto2.UserDto

	if user.Id == 0 {
		return userDto, e.NewBadRequestApiError("user not found")
	}
	userDto.Id = user.Id
	userDto.Name = user.Name
	userDto.LastName = user.LastName
	userDto.Email = user.Email
	userDto.Password = user.Password
	userDto.UserType = user.UserType
	userDto.Dni = user.Dni

	return userDto, nil
}

//login

var jwtKey = []byte("secret_key")

func (s *userService) LoginUser(loginDto dto2.LoginDto) (dto2.TokenDto, e.ApiError) {

	log.Debug(loginDto) //para registrar el contenido de loginDto
	var user model.User = userClient.GetUserByEmail(loginDto.Email)

	var tokenDto dto2.TokenDto

	if user.Id == 0 {
		return tokenDto, e.NewBadRequestApiError("user not found")
	}

	//pasamos password como slice de bytes
	//hashea con md5.sum
	var pswMd5 = md5.Sum([]byte(loginDto.Password))
	//convertir a cadena hexadecimal
	pswMd5String := hex.EncodeToString(pswMd5[:])

	if pswMd5String == user.Password {
		//se firma el token para verificar autenticidad
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id_user": user.Id,
		})
		tokenString, _ := token.SignedString(jwtKey)
		tokenDto.Token = tokenString
		tokenDto.IdUser = user.Id

		return tokenDto, nil
	} else {
		return tokenDto, e.NewBadRequestApiError("contraseña incorrecta")
	}

}

func (s *userService) InsertUser(userDto dto2.UserDto) (dto2.TokenDto, e.ApiError) {

	log.Debug(userDto) //para registrar el contenido de userDto
	var user model.User = userClient.GetUserByEmail(userDto.Email)

	var tokenDto dto2.TokenDto

	if user.Id == 0 { //el usuario no esta registrado y puedo crear uno nuevo

		//pasamos password como slice de bytes
		//hashea con md5.sum
		var pswMd5 = md5.Sum([]byte(userDto.Password))
		//convertir a cadena hexadecimal
		pswMd5String := hex.EncodeToString(pswMd5[:])
		//se firma el token para verificar autenticidad
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id_user": user.Id,
		})
		tokenString, _ := token.SignedString(jwtKey)
		tokenDto.Token = tokenString
		tokenDto.IdUser = user.Id

		user.Id = userDto.Id
		user.Name = userDto.Name
		user.LastName = userDto.LastName
		user.Email = userDto.Email
		user.Password = pswMd5String
		user.UserType = userDto.UserType
		user.Dni = userDto.Dni

		user = userClient.InsertUser(user)

		return tokenDto, nil

	} else { //el usuario ya existe
		return tokenDto, e.NewBadRequestApiError("user already exists")
	}

}
*/
