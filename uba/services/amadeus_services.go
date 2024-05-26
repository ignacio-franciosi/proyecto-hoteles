package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

type amadeusService struct{}

type amadeusServiceInterface interface {
	CheckAvailabilityAmadeus(hotelId string, startDate time.Time, endDate time.Time) (bool, error)
}

var AmadeusService amadeusServiceInterface

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

var (
	amadeusToken    string
	mutex           sync.Mutex
	lastTokenUpdate time.Time
)

func init() {
	AmadeusService = &amadeusService{}
}

func GetAmadeusToken() (string, error) {
	mutex.Lock()
	defer mutex.Unlock()

	// Verificar si el token necesita actualización (más de 30 minutos desde la última actualización)
	if time.Since(lastTokenUpdate) > 30*time.Minute {
		err := updateAmadeusToken()
		if err != nil {
			return "", err
		}
	}

	// Devolver el token actual
	return amadeusToken, nil
}

func updateAmadeusToken() error {
	url := "https://test.api.amadeus.com/v1/security/oauth2/token"
	method := "POST"

	payload := strings.NewReader("grant_type=client_credentials&client_id=rZ5NyDeG74mzhfaFqk2V7VaM2yUQtZux&client_secret=IoO2Wocmr1Aqrw2X")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var response TokenResponse

	// Decodificar la respuesta JSON en la estructura TokenResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	// Actualizar el token y la hora de la última actualización
	amadeusToken = response.AccessToken
	lastTokenUpdate = time.Now()

	fmt.Println("Token actualizado:", amadeusToken)
	return nil
}

func (s *amadeusService) CheckAvailabilityAmadeus(hotelId string, startDate time.Time, endDate time.Time) (bool, error) {
	// Construir la URL concatenando los parámetros

	formatedStartDate := startDate.Format("2006-01-02")
	formatedEndDate := endDate.Format("2006-01-02")

	apiUrl := "https://test.api.amadeus.com/v3/shopping/hotel-offers"
	apiUrl += "?hotelIds=" + hotelId
	apiUrl += "&checkInDate=" + formatedStartDate
	apiUrl += "&checkOutDate=" + formatedEndDate

	// Crear la request HTTP GET
	request, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return false, fmt.Errorf("error al crear la request: %v", err)
	}

	// Obtener el token de Amadeus
	token, err := GetAmadeusToken()
	if err != nil {
		return false, fmt.Errorf("error al obtener el token: %v", err)
	}
	request.Header.Set("Authorization", "Bearer "+token)

	// Realizar la request HTTP
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return false, fmt.Errorf("error al realizar la request: %v", err)
	}
	defer response.Body.Close()

	// Verificar el código de estado de la response
	if response.StatusCode != http.StatusOK {
		// Leer el cuerpo de la response para más información
		responseBody, _ := io.ReadAll(response.Body)
		return false, fmt.Errorf("no hay disponibilidad. Código: %d. Respuesta: %s", response.StatusCode, string(responseBody))
	}

	// Leer el cuerpo de la response
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return false, fmt.Errorf("error al leer la response: %v", err)
	}

	// Crear una estructura para deserializar el JSON de la response
	var responseStruct struct {
		Data []map[string]interface{} `json:"data"`
	}

	// Decodificar el JSON
	if err := json.Unmarshal(responseBody, &responseStruct); err != nil {
		return false, fmt.Errorf("error al decodificar el JSON de la response: %v", err)
	}

	// Verificar si hay datos disponibles
	if len(responseStruct.Data) > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
