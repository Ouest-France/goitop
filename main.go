package goitop

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	Client   *http.Client
	Address  string
	User     string
	Password string
}

type APIResponse struct {
	Code    int
	Message string
	Objects map[string]APIObject
}

type APIObject struct {
	Code    int
	Message string
	Class   string
	Key     string
	Fields  map[string]interface{}
}

type VM struct {
	Name  string
	OrgID string
	EnvID string
	ID    string
}

func NewClient(address, user, password string) *Client {
	return &Client{
		Client:   &http.Client{},
		Address:  address,
		User:     user,
		Password: password,
	}
}

func (c *Client) CreateVM(name, org, env string) (string, error) {

	payload := map[string]interface{}{
		"operation":     "core/create",
		"comment":       "create new VM",
		"class":         "VirtualMachine",
		"output_fields": "id",
		"fields": map[string]string{
			"name":   name,
			"org_id": org,
			"env_id": env,
		},
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
		return "", fmt.Errorf("Create VM request failed with status code %d", resp.StatusCode)
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
		return "", fmt.Errorf("Create VM request failed with code %d: %s", result.Code, result.Message)
	}

	if len(result.Objects) != 1 {
		return "", fmt.Errorf("Too few or to many objects in create VM response")
	}

	var firstObject APIObject
	for _, object := range result.Objects {
		firstObject = object
		break
	}

	return firstObject.Key, nil
}

func (c *Client) UpdateVM(id, name, org, env string) error {

	payload := map[string]interface{}{
		"operation":     "core/update",
		"comment":       "update VM",
		"class":         "VirtualMachine",
		"key":           id,
		"output_fields": "id",
		"fields": map[string]string{
			"name":   name,
			"org_id": org,
			"env_id": env,
		},
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	form := url.Values{
		"json_data": []string{string(jsonPayload)},
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/webservices/rest.php?version=1.3", c.Address), strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.User, c.Password)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Update VM request failed with status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var result APIResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return err
	}

	if result.Code != 0 {
		return fmt.Errorf("Update VM request failed with code %d: %s", result.Code, result.Message)
	}

	if len(result.Objects) != 1 {
		return fmt.Errorf("Too few or to many objects in create VM response")
	}

	return nil
}

func (c *Client) GetVM(id string) (VM, error) {

	payload := map[string]string{
		"operation":     "core/get",
		"class":         "VirtualMachine",
		"key":           id,
		"output_fields": "id,name,org_id,env_id",
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return VM{}, err
	}

	form := url.Values{
		"json_data": []string{string(jsonPayload)},
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/webservices/rest.php?version=1.3", c.Address), strings.NewReader(form.Encode()))
	if err != nil {
		return VM{}, err
	}
	req.SetBasicAuth(c.User, c.Password)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.Client.Do(req)
	if err != nil {
		return VM{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return VM{}, err
	}

	var result APIResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return VM{}, err
	}

	if len(result.Objects) != 1 {
		return VM{}, nil
	}

	var firstObject APIObject
	for _, object := range result.Objects {
		firstObject = object
		break
	}

	vm := VM{
		ID:    firstObject.Fields["id"].(string),
		Name:  firstObject.Fields["name"].(string),
		OrgID: firstObject.Fields["org_id"].(string),
		EnvID: firstObject.Fields["env_id"].(string),
	}

	return vm, nil
}

func (c *Client) DeleteVM(id string) error {

	payload := map[string]interface{}{
		"operation": "core/delete",
		"comment":   "delete VM",
		"class":     "VirtualMachine",
		"key":       id,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	form := url.Values{
		"json_data": []string{string(jsonPayload)},
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/webservices/rest.php?version=1.3", c.Address), strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.User, c.Password)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("Delete VM request failed with status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil

	var result APIResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return err
	}

	if result.Code != 0 {
		return fmt.Errorf("Delete VM request failed with code %d: %s", result.Code, result.Message)
	}

	if len(result.Objects) != 1 {
		return fmt.Errorf("Too few or to many objects in delete VM response")
	}

	return nil
}
