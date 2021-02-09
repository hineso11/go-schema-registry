package pkg

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type SchemaVersion struct {
	Subject	string `json:"subject"`
	Id	int	`json:"id"`
	Version	int	`json:"version"`
	Schema	string	`json:"schema"`
}

func (c schemaRegistryClient) SubjectVersionSchema(subject string, version int) (string, error) {

	req, err := c.baseGetRequest("/subjects/" + subject + "/versions/" + strconv.Itoa(version))

	if err != nil {
		return "", err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		var apiError SchemaRegistryApiError
		err = json.Unmarshal(body, &apiError)

		if err != nil {
			return "", err
		}

		return "", apiError
	}

	return string(body), nil
}

func (c schemaRegistryClient) SubjectVersion(subject string, version int) (*SchemaVersion, error) {

	req, err := c.baseGetRequest("/subjects/" + subject + "/versions/" + strconv.Itoa(version))

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

	var schemaVersion SchemaVersion
	err = json.Unmarshal(body, &schemaVersion)

	if err != nil {
		return nil, err
	}

	return &schemaVersion, nil
}

func (c schemaRegistryClient) SubjectVersions(subject string) ([]int, error) {

	req, err := c.baseGetRequest("/subjects/" + subject + "/versions")

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

	var versionNumbers []int
	err = json.Unmarshal(body, &versionNumbers)

	if err != nil {
		return nil, err
	}

	return versionNumbers, nil
}

func (c schemaRegistryClient) Subjects(subject *string, deleted *bool) ([]string, error) {

	req, err := c.baseGetRequest("/subjects")

	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	if subject != nil {
		query.Add("subject", *subject)
	}
	if deleted != nil {
		query.Add("deleted", strconv.FormatBool(*deleted))
	}
	req.URL.RawQuery = query.Encode()

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


	var schemaNames []string
	err = json.Unmarshal(body, &schemaNames)

	if err != nil {
		return nil, err
	}

	return schemaNames, nil
}
