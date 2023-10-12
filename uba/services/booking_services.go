package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Hotel struct {
	//completar
}

type AmadeusCredentials struct {
	ClientID     string
	ClientSecret string
}

func getAccessToken(credentials AmadeusCredentials) (string, error) {
	u, err := url.Parse("https://test.api.amadeus.com/v3/security/oauth2/token")

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", credentials.ClientID)
	data.Set("client_secret", credentials.ClientSecret)

	resp, err := http.PostForm(u.String(), data)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get access token, status code: %d", resp.StatusCode)
	}

	var tokenResponse map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&tokenResponse)
	if err != nil {
		return "", err
	}

	accessToken, ok := tokenResponse["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("failed to parse access token from response")
	}

	return accessToken, nil
}

func getHotelAvailability(accessToken string, hotels []Hotel) error {
	// Endpoint para obtener disponibilidad de hoteles
	baseURL := "https://test.api.amadeus.com/v3/shopping/hotel-offers"

	for _, hotel := range hotels {
		queryParams := url.Values{}
		queryParams.Set("city", hotel.city)
		queryParams.Set("startDate", booking.startDate) // Ejemplo de fecha de inicio
		queryParams.Set("endDate", booking.endDate)     // Ejemplo de fecha de fin

		fullURL := baseURL + "?" + queryParams.Encode()

		// Crea una solicitud HTTP con el método GET
		req, err := http.NewRequest("GET", fullURL, nil)
		if err != nil {
			return err
		}

		// Agrega el token de acceso en el encabezado de autorización
		req.Header.Set("Authorization", "Bearer "+accessToken)

		// Realiza la solicitud
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Procesa la respuesta aquí según tus necesidades
		// Puedes leer el cuerpo de la respuesta usando resp.Body
	}

	return nil
}

func main() {
	credentials := AmadeusCredentials{
		ClientID:     "8O9tG0NdPxxEA90mbxbAgRrJxGU002Yb",
		ClientSecret: "qfurhUwj1OB0AjlE",
	}

	// Obtén el token de acceso
	accessToken, err := getAccessToken(credentials)
	if err != nil {
		fmt.Printf("Error obteniendo el token: %v\n", err)
		return
	}
	//-----------------------------------------------------------------------
	// Realiza la solicitud de disponibilidad de hoteles

	otherServiceURL := "https://otro_microservicio.com/hoteles"
	resp, err := http.Get(otherServiceURL)
	if err != nil {
		fmt.Printf("Error obteniendo datos de hoteles: %v\n", err)
		return
	}
	defer resp.Body.Close()

	var hotels []Hotel
	if err := json.NewDecoder(resp.Body).Decode(&hotels); err != nil {
		fmt.Printf("Error decodificando los datos de hoteles: %v\n", err)
		return
	}

	err = getHotelAvailability(accessToken, hotels)
	if err != nil {
		fmt.Printf("Error obteniendo disponibilidad de hoteles: %v\n", err)
		return
	}
}
