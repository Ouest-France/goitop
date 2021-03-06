package goitop

import (
	"fmt"
)

type VM struct {
	Name                  string
	OrgID                 string
	EnvID                 string
	ClusterID             string
	ExploitationServiceID string
	ID                    string
	Backup                string
	BackupID              string
	Description           string
}

func (c *Client) CreateVM(name, org_id, env_id, cluster_id, exploitationservice_id, backup, backup_id, description string) (string, error) {

	if backup != "yes" && backup != "no" {
		return "", fmt.Errorf("backup parameter must be yes or no, got %q", backup)
	}

	payload := map[string]interface{}{
		"operation":     "core/create",
		"comment":       "create new VM",
		"class":         "VirtualMachine",
		"output_fields": "id",
		"fields": map[string]string{
			"name":                   name,
			"org_id":                 org_id,
			"env_id":                 env_id,
			"cluster_id":             cluster_id,
			"exploitationservice_id": exploitationservice_id,
			"backup":                 backup,
			"backup_id":              backup_id,
			"description":            description,
		},
	}

	result, err := Request(c, payload)
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

func (c *Client) UpdateVM(id, name, org_id, env_id, cluster_id, exploitationservice_id, backup, backup_id, description string) error {

	if backup != "yes" && backup != "no" {
		return fmt.Errorf("backup parameter must be yes or no, got %q", backup)
	}

	payload := map[string]interface{}{
		"operation":     "core/update",
		"comment":       "update VM",
		"class":         "VirtualMachine",
		"key":           id,
		"output_fields": "id",
		"fields": map[string]string{
			"name":                   name,
			"org_id":                 org_id,
			"env_id":                 env_id,
			"cluster_id":             cluster_id,
			"exploitationservice_id": exploitationservice_id,
			"backup":                 backup,
			"backup_id":              backup_id,
			"description":            description,
		},
	}

	result, err := Request(c, payload)
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

	payload := map[string]interface{}{
		"operation":     "core/get",
		"class":         "VirtualMachine",
		"key":           id,
		"output_fields": "id,name,org_id,env_id,cluster_id,exploitationservice_id,backup,backup_id,description",
	}

	result, err := Request(c, payload)
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
		ID:                    firstObject.Fields["id"].(string),
		Name:                  firstObject.Fields["name"].(string),
		OrgID:                 firstObject.Fields["org_id"].(string),
		EnvID:                 firstObject.Fields["env_id"].(string),
		ClusterID:             firstObject.Fields["cluster_id"].(string),
		ExploitationServiceID: firstObject.Fields["exploitationservice_id"].(string),
		Backup:                firstObject.Fields["backup"].(string),
		BackupID:              firstObject.Fields["backup_id"].(string),
		Description:           firstObject.Fields["description"].(string),
	}

	return vm, nil
}

func (c *Client) GetAllVM() ([]VM, error) {

	payload := map[string]interface{}{
		"operation":     "core/get",
		"class":         "VirtualMachine",
		"key":           "SELECT VirtualMachine",
		"output_fields": "id,name,org_id,env_id,cluster_id,exploitationservice_id,backup,backup_id,description",
	}

	result, err := Request(c, payload)
	if err != nil {
		return []VM{}, err
	}

	if len(result.Objects) < 1 {
		return []VM{}, fmt.Errorf("No VM found")
	}

	vms := []VM{}
	for _, object := range result.Objects {
		vms = append(vms, VM{
			ID:                    object.Fields["id"].(string),
			Name:                  object.Fields["name"].(string),
			OrgID:                 object.Fields["org_id"].(string),
			EnvID:                 object.Fields["env_id"].(string),
			ClusterID:             object.Fields["cluster_id"].(string),
			ExploitationServiceID: object.Fields["exploitationservice_id"].(string),
			Backup:                object.Fields["backup"].(string),
			BackupID:              object.Fields["backup_id"].(string),
			Description:           object.Fields["description"].(string),
		})
	}

	return vms, nil
}

func (c *Client) DeleteVM(id string) error {

	payload := map[string]interface{}{
		"operation": "core/delete",
		"comment":   "delete VM",
		"class":     "VirtualMachine",
		"key":       id,
	}

	result, err := Request(c, payload)
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
