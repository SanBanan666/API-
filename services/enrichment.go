package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type EnrichmentService struct {
	client *http.Client
}

func NewEnrichmentService() *EnrichmentService {
	return &EnrichmentService{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

type AgeResponse struct {
	Age int `json:"age"`
}

type GenderResponse struct {
	Gender string `json:"gender"`
}

type NationalityResponse struct {
	Country []struct {
		CountryID   string  `json:"country_id"`
		Probability float64 `json:"probability"`
	} `json:"country"`
}

func (s *EnrichmentService) GetAge(name string) (int, error) {
	resp, err := s.client.Get(fmt.Sprintf("https://api.agify.io/?name=%s", name))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var ageResp AgeResponse
	if err := json.Unmarshal(body, &ageResp); err != nil {
		return 0, err
	}

	return ageResp.Age, nil
}

func (s *EnrichmentService) GetGender(name string) (string, error) {
	resp, err := s.client.Get(fmt.Sprintf("https://api.genderize.io/?name=%s", name))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var genderResp GenderResponse
	if err := json.Unmarshal(body, &genderResp); err != nil {
		return "", err
	}

	return genderResp.Gender, nil
}

func (s *EnrichmentService) GetNationality(name string) (string, error) {
	resp, err := s.client.Get(fmt.Sprintf("https://api.nationalize.io/?name=%s", name))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var nationalityResp NationalityResponse
	if err := json.Unmarshal(body, &nationalityResp); err != nil {
		return "", err
	}

	if len(nationalityResp.Country) > 0 {
		return nationalityResp.Country[0].CountryID, nil
	}

	return "", nil
}
