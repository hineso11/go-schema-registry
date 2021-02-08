package pkg

type SchemaCompatibilityLevel string

const (
	Backward SchemaCompatibilityLevel = "BACKWARD"
	BackwardTransitive	SchemaCompatibilityLevel = "BACKWARD_TRANSITIVE"
	Forward	SchemaCompatibilityLevel = "FORWARD"
	ForwardTransitive SchemaCompatibilityLevel = "FORWARD_TRANSITIVE"
	Full	SchemaCompatibilityLevel = "FULL"
	FullTransitive	SchemaCompatibilityLevel = "FULL_TRANSITIVE"
	None	SchemaCompatibilityLevel = "NONE"
)

type SchemaRegistryConfig struct {
	Compatibility SchemaCompatibilityLevel `json:"compatibility"`
}

type SchemaRegistryClient interface {
	GetConfig() (*SchemaRegistryConfig, error)
}

type schemaRegistryClient struct {
	baseUrl string
	apiKey *string
	apiSecret *string
}

func (c schemaRegistryClient) GetConfig() (*SchemaRegistryConfig, error) {

	


	return &SchemaRegistryConfig{
		Compatibility: Backward,
	}, nil
}

func NewSchemaRegistryClientWithAuth(baseUrl string, apiKey string, apiSecret string) SchemaRegistryClient {

	return schemaRegistryClient{
		baseUrl: 	baseUrl,
		apiKey:    &apiKey,
		apiSecret: &apiSecret,
	}
}

func NewSchemaRegistryClient(baseUrl string) SchemaRegistryClient {

	return schemaRegistryClient{
		baseUrl: baseUrl,
	}
}