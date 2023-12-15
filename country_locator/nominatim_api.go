package countrylocator

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetCountryByCoordinates makes a request to the Nominatim API and returns the country name
func GetCountryByCoordinates(lat, lon float64) (string, error) {
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?lat=%f&lon=%f&zoom=3&format=json", lat, lon)

	// Make HTTP request
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("Error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP request failed with status code: %v", resp.StatusCode)
	}

	// Decode JSON response
	var locationInfo map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&locationInfo)
	if err != nil {
		return "", fmt.Errorf("Error decoding Nominatim JSON response: %v", err)
	}
	countryCode := (locationInfo["address"].(map[string]any))["country_code"].(string)
	return countryCode, nil
}
