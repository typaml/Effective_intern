package utils

import (
	"Effective_intern/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func FetchPeopleInfo(serie, number string) (*models.User, error) {
	url := fmt.Sprintf("http://api.example.com/info?passportSerie=%d&passportNumber=%d", serie, number) //запрос к внешнему api, для получения доп инфы
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error response from API: %s", resp.Status)
	}

	var people models.User
	if err := json.NewDecoder(resp.Body).Decode(&people); err != nil {
		return nil, err
	}

	return &people, nil
}
