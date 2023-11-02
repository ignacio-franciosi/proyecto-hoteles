package services

import (
	"consumer/config"
	client "consumer/services/repositories"
	"fmt"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

type ConsumerService struct {
	queue *client.QueueClient
}

func NewConsumer(
	queue *client.QueueClient,
) *ConsumerService {
	return &ConsumerService{
		queue: queue,
	}
}

func (s *ConsumerService) TopicConsumer(topic string) {
	err := s.queue.ProcessMessages(config.EXCHANGE, topic, func(message string) {
		var resp *http.Response
		//var err error
		cli := &http.Client{}

		// Analiza el contenido del mensaje para determinar la acción
		action, id, err := parseMessage(message)

		if action == "GET" {
			// Realiza una solicitud HTTP de tipo GET para obtener un elemento
			resp, err = http.Get(fmt.Sprintf("http://%s:%d/hotels/%s", config.LBHOST, config.LBPORT, id))
		} else if action == "CREATE" {
			// Realiza una solicitud HTTP de tipo POST para crear un elemento
			req, err := http.NewRequest("POST", fmt.Sprintf("http://%s:%d/hotels", config.LBHOST, config.LBPORT), strings.NewReader(id))
			req.Header.Set("Content-Type", "application/json")
			if err != nil {
				log.Error(err)
			}
			resp, err = cli.Do(req)
		} else if action == "UPDATE" {
			// Realiza una solicitud HTTP de tipo PUT para modificar un elemento
			req, err := http.NewRequest("PUT", fmt.Sprintf("http://%s:%d/hotels/%s", config.LBHOST, config.LBPORT, id), strings.NewReader(id))
			req.Header.Set("Content-Type", "application/json")
			if err != nil {
				log.Error(err)
			}
			resp, err = cli.Do(req)
		}

		log.Debug("Hotel action " + action + " for ID " + id)

		if err != nil {
			log.Error("Error in HTTP request: " + err.Error())
			log.Debug(resp)
		}
	})

	if err != nil {
		log.Error("Error starting consumer processing", err)
	}
}

func parseMessage(message string) (action string, id string, err error) {
	// Divide el mensaje en función del carácter '|'
	parts := strings.Split(message, "|")

	// Verifica si el mensaje tiene el formato correcto
	if len(parts) != 2 {
		return "", "", fmt.Errorf("El formato del mensaje es incorrecto: %s", message)
	}

	// La primera parte del mensaje es la acción y la segunda parte es el ID
	action = parts[0]
	id = parts[1]

	// Devuelve la acción y el ID extraídos del mensaje
	return action, id, nil
}

/*
func parseMessage(message string) (action string, id string) {
    // Analiza el contenido del mensaje para determinar la acción (GET, CREATE, MODIFY)
    // y el ID del elemento
    // Puedes implementar la lógica de análisis del mensaje según tu formato de mensajes
    // Por ejemplo, si el mensaje es "GET|12345", se puede dividir en acción "GET" e ID "12345".
    parts := strings.Split(message, "|")

    if len(parts) != 2 {
        // Mensaje mal formateado
        return "", ""
    }

    return parts[0], parts[1]
}
*/

/*
func (s *ConsumerService) TopicConsumer(topic string) {
	err := s.queue.ProcessMessages(config.EXCHANGE, topic, func(id string) {
		var resp *http.Response
		var err error
		cli := &http.Client{}
		strs := strings.Split(id, ".")
		if len(strs) < 2 {
			resp, err = http.Get(fmt.Sprintf("http://%s:%d/items/%s", config.LBHOST, config.LBPORT, id))
		} else {
			if strs[1] == "delete" {
				req, err := http.NewRequest("DELETE", fmt.Sprintf("http://%s:%d/items/%s", config.LBHOST, config.LBPORT, strs[0]), nil)
				if err != nil {
					log.Error(err)
				}
				resp, err = cli.Do(req)
				if err != nil {
					log.Error(err)
					log.Debug(resp)
				}
			}
		}
		log.Debug("Item sent " + id)
		if err != nil {
			log.Debug("error in get request")
			log.Debug(resp)
		}
	})
	if err != nil {
		log.Error("Error starting worker processing", err)
	}
}
*/
