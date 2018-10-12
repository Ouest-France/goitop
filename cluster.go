package goitop

import (
	"fmt"
)

func (c *Client) GetCluster(name string) (string, error) {

	payload := map[string]interface{}{
		"operation": "core/get",
		"class":     "Cluster",
		"key": map[string]string{
			"name": name,
		},
		"output_fields": "id",
	}

	result, err := Request(c, payload)
	if err != nil {
		return "", err
	}

	if result.Code != 0 {
		return "", fmt.Errorf("Get cluster request failed with code %d: %s", result.Code, result.Message)
	}

	if len(result.Objects) > 1 {
		return "", fmt.Errorf("Too many objects in get cluster response")
	}

	if len(result.Objects) < 1 {
		return "", fmt.Errorf("Cluster not found")
	}

	var firstObject APIObject
	for _, object := range result.Objects {
		firstObject = object
		break
	}

	return firstObject.Fields["id"].(string), nil
}
