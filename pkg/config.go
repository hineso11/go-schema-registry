package pkg

import (
	"encoding/json"
	"io/ioutil"
)

type SchemaRegistryConfig struct {
	Compatibility SchemaCompatibilityLevel `json:"compatibilityLevel"`
}

const (
	Backward SchemaCompatibilityLevel = "BACKWARD"
	BackwardTransitive	SchemaCompatibilityLevel = "BACKWARD_TRANSITIVE"
	Forward	SchemaCompatibilityLevel = "FORWARD"
	ForwardTransitive SchemaCompatibilityLevel = "FORWARD_TRANSITIVE"
	Full	SchemaCompatibilityLevel = "FULL"
	FullTransitive	SchemaCompatibilityLevel = "FULL_TRANSITIVE"
	None	SchemaCompatibilityLevel = "NONE"
)

func (c schemaRegistryClient) Config() (*SchemaRegistryConfig, error) {

	req, err := c.baseGetRequest("/config")

	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		var apiError SchemaRegistryApiError
		err = json.Unmarshal(body, &apiError)

		if err != nil {
			return nil, err
		}

		return nil, apiError
	}

	var config SchemaRegistryConfig
	err = json.Unmarshal(body, &config)

	if err != nil {
		return nil, err
	}

	return &config, nil
}
