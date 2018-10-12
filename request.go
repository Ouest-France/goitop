package goitop

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func Request(client *Client, payload map[string]interface{}) (APIResponse, error) {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return APIResponse{}, err
	}

	form := url.Values{
		"json_data": []string{string(jsonPayload)},
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/webservices/rest.php?version=1.3", client.Address), strings.NewReader(form.Encode()))
	if err != nil {
		return APIResponse{}, err
	}
	req.SetBasicAuth(client.User, client.Password)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Client.Do(req)
	if err != nil {
		return APIResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return APIResponse{}, fmt.Errorf("Request failed with status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return APIResponse{}, err
	}

	var result APIResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return APIResponse{}, err
	}

	return result, nil
}
