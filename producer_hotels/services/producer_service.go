package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"producer_hotels/config"
	client "producer_hotels/services/repositories"

	log "github.com/sirupsen/logrus"
)

type ProducerService struct {
	queue *client.QueueClient
}

func NewProducer(
	queue *client.QueueClient,
) *ProducerService {
	return &ProducerService{
		queue: queue,
	}
}

func (s *ProducerService) TopicProducer(topic string, action string, data interface{}) {
	err := s.queue.ProcessMessages(config.EXCHANGE, topic, func(id string) {
		client := &http.Client{}
		var req *http.Request
		var err error

		if action == "CREATE" {
			// Crear una solicitud HTTP de tipo POST para la creación de un hotel
			req, err = http.NewRequest("POST", fmt.Sprintf("http://%s:%d/hotels", config.HOTELSHOST, config.HOTELSPORT), nil)
		} else if action == "UPDATE" {
			// Crear una solicitud HTTP de tipo PUT para la modificación de un hotel
			req, err = http.NewRequest("PUT", fmt.Sprintf("http://%s:%d/hotels/%s", config.HOTELSHOST, config.HOTELSPORT, id), nil)
		} else {
			// Manejar otras acciones según sea necesario
			log.Error("Acción no válida:", action)
		}

		// Verificar si se creó correctamente la solicitud
		if err != nil {
			log.Error("Error al crear la solicitud HTTP:", err)
			return
		}

		// Agregar datos adicionales al cuerpo de la solicitud si es necesario
		if data != nil {
			// Serializar los datos a formato JSON
			dataBytes, err := json.Marshal(data)
			if err != nil {
				log.Error("Error al serializar datos a JSON:", err)
				return
			}

			// Configurar el cuerpo de la solicitud con los datos serializados
			req.Body = ioutil.NopCloser(bytes.NewReader(dataBytes))
			req.Header.Set("Content-Type", "application/json")
		}

		log.Debug("Solicitud " + action + " enviada para el hotel con ID " + id)

		// Realizar la solicitud HTTP
		_, err = client.Do(req)

		// Verificar si hubo errores en la solicitud
		if err != nil {
			log.Error("Error en la solicitud HTTP:", err)
		}
	})

	if err != nil {
		log.Error("Error al iniciar el procesamiento del trabajador:", err)
	}
}

/*
func (s *ProducerService) TopicProducer(topic string) {
	err := s.queue.ProcessMessages(config.EXCHANGE, topic, func(id string) {
		client := &http.Client{}
		//Revisar http
		req, err := http.NewRequest("DELETE", fmt.Sprintf("http://%s:%d/users/%s/items", config.ITEMSHOST, config.ITEMSPORT, id), nil)
		log.Debug("Item delete sent " + id)

		if err != nil {
			log.Debug("error in delete request")
		}

		_, err = client.Do(req)
		if err != nil {
			log.Error(err)
		}

	})
	if err != nil {
		log.Error("Error starting worker processing", err)
	}
}
*/
