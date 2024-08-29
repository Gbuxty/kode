package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)


type YandexSpellerResponse []struct {
	Code int      `json:"code"`
	Word string   `json:"word"`
	S    []string `json:"s"`
}


func validateSpelling(text string) error {
	url := "https://speller.yandex.net/services/spellservice.json/checkText"
	payload := bytes.NewBufferString(fmt.Sprintf("text=%s", text))
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var spellerResponse YandexSpellerResponse
	if err := json.Unmarshal(body, &spellerResponse); err != nil {
		return err
	}

	if len(spellerResponse) > 0 {
		return fmt.Errorf("spelling errors found: %v", spellerResponse)
	}

	return nil
}