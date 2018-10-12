package goitop

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func (c *Client) GetEnvironment(name string) (string, error) {

	payload := map[string]interface{}{
		"operation": "core/get",
		"class":     "Environment",
		"key": map[string]string{
			"name": name,
		},
		"output_fields": "id",
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	form := url.Values{
		"json_data": []string{string(jsonPayload)},
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/webservices/rest.php?version=1.3", c.Address), strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(c.User, c.Password)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Get environment request with http status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result APIResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	if result.Code != 0 {
		return "", fmt.Errorf("Get environment request failed with code %d: %s", result.Code, result.Message)
	}

	if len(result.Objects) > 1 {
		return "", fmt.Errorf("Too many objects in get environment response")
	}

	if len(result.Objects) < 1 {
		return "", fmt.Errorf("Environment not found")
	}

	var firstObject APIObject
	for _, object := range result.Objects {
		firstObject = object
		break
	}

	return firstObject.Fields["id"].(string), nil
}
