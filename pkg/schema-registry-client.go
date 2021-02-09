package pkg

import (
	"net/http"
	"strconv"
)

type SchemaCompatibilityLevel string

type SchemaRegistryApiError struct {
	ErrorCode	int	`json:"error_code"`
	Message		string	`json:"message"`
}

func (s SchemaRegistryApiError) Error() string {

	return "schema registry API: " + s.Message + "- error code (" + strconv.Itoa(s.ErrorCode) + ")"
}

type SchemaRegistryClient interface {
	Config() (*SchemaRegistryConfig, error)
	Subjects(subject *string, deleted *bool) ([]string, error)
	SubjectVersions(subject string) ([]int, error)
	SubjectVersion(subject string, version int) (*SchemaVersion, error)
	SubjectVersionSchema(subject string, version int) (string, error)
}

type schemaRegistryClient struct {
	baseUrl string
	apiKey *string
	apiSecret *string
	httpClient http.Client
}

func (c schemaRegistryClient) baseGetRequest(endpoint string) (*http.Request, error) {

	req, err := http.NewRequest(http.MethodGet, c.baseUrl + endpoint, nil)

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(*c.apiKey, *c.apiSecret)
	return req, nil
}

func NewSchemaRegistryClientWithAuth(baseUrl string, apiKey string, apiSecret string) SchemaRegistryClient {

	return schemaRegistryClient{
		baseUrl: 	baseUrl,
		apiKey:    &apiKey,
		apiSecret: &apiSecret,
		httpClient: http.Client{},
	}
}

func NewSchemaRegistryClient(baseUrl string) SchemaRegistryClient {

	return schemaRegistryClient{
		baseUrl: baseUrl,
		httpClient: http.Client{},
	}
}