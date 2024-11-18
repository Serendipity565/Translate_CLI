package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

type Config struct {
	MicrosoftAzure struct {
		Key      string `yaml:"key"`
		Location string `yaml:"location"`
	} `yaml:"Microsoft_Azure"`
}

func OpenYaml() (Config, error) {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("failed to get executable path: %v", err)
	}
	configFilePath := filepath.Join(filepath.Dir(exePath), "config", "user.yaml")

	//configFilePathForTest := "../config/user.yaml"

	file, err := os.Open(configFilePath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	var config Config

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}

type Translation struct {
	Text string `json:"text"`
	To   string `json:"to"`
}

type TranslateResponse struct {
	Translations []Translation `json:"translations"`
}

func Translate(text, from, to string) (string, error) {
	config, err := OpenYaml()
	if err != nil {
		return "", err
	}
	key := config.MicrosoftAzure.Key
	endpoint := "https://api.cognitive.microsofttranslator.com/"
	uri := endpoint + "/translate?api-version=3.0"

	// location, also known as region.
	// required if you're using a multi-service or regional (not global) resource. It can be found in the Azure portal on the Keys and Endpoint page.
	location := config.MicrosoftAzure.Location

	// Build the request URL. See: https://go.dev/pkg/net/url/#example_URL_Parse
	u, _ := url.Parse(uri)
	q := u.Query()
	q.Add("from", from)
	q.Add("to", to)

	u.RawQuery = q.Encode()

	// Create an anonymous struct for your request body and encode it to JSON
	body := []struct {
		Text string
	}{
		{Text: text},
	}
	b, _ := json.Marshal(body)

	// Build the HTTP POST request
	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}
	// Add required headers to the request
	req.Header.Add("Ocp-Apim-Subscription-Key", key)
	// location required if you're using a multi-service or regional (not global) resource.
	req.Header.Add("Ocp-Apim-Subscription-Region", location)
	req.Header.Add("Content-Type", "application/json")

	// Call the Translator API
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	// Decode the JSON response
	var result []struct {
		Translations []struct {
			Text string `json:"text"`
			To   string `json:"to"`
		} `json:"translations"`
	}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode JSON response: %w", err)
	}

	if len(result) > 0 && len(result[0].Translations) > 0 {
		return result[0].Translations[0].Text, nil
	}

	return "", fmt.Errorf("no translations found in response")
}
