package synchronization

type Labels map[string][]string

type SpecFormat string

const (
	SpecFormatYaml SpecFormat = "YAML"
	SpecFormatJSON SpecFormat = "JSON"
)

type EventAPISpecType string

const (
	EventAPISpecTypeAsyncAPI EventAPISpecType = "ASYNC_API"
)

type APISpecType string

const (
	APISpecTypeOdata   APISpecType = "ODATA"
	APISpecTypeOpenAPI APISpecType = "OPEN_API"
)

type DocumentFormat string

const (
	DocumentFormatMarkdown DocumentFormat = "MARKDOWN"
)

type Application struct {
	ID          string
	Name        string
	Description *string
	Labels      Labels
	APIs        []APIDefinition
	EventAPIs   []EventAPIDefinition
	Documents   []Document
}

// APIDefinition contains API data such as URL, credentials and spec
type APIDefinition struct {
	ID                string
	Name              string
	Description       string
	TargetUrl         string
	APIType           APISpecType
	RequestParameters RequestParameters
	Credentials       *Credentials
	APISpec           *APISpec
}

// EventAPIDefinition contains Event API details such
type EventAPIDefinition struct {
	ID          string
	Name        string
	Description string
	APISpec     *APISpec
}

// Document contains data of document stored in the Asset Store
type Document struct {
	ID            string
	ApplicationID string
	Title         string
	DisplayName   string
	Description   string
	Format        DocumentFormat
	Kind          *string
	Data          []byte
}

// APISpec contains spec BLOB and
type APISpec struct {
	Data []byte
	Type APISpecType
}

// Credentials contains OAuth, BasicAuth or Certificates configuration along with optional CSRF data.
type Credentials struct {
	// OAuth configuration
	Oauth *Oauth
	// BasicAuth configuration
	Basic *Basic

	// Optional CSRF Data
	CSRFInfo *CSRFInfo
}

// Oauth contains data for performing Oauth token request
type Oauth struct {
	// URL to OAuth token provider.
	URL string
	// ClientID to use for authentication.
	ClientID string
	// ClientSecret to use for authentication.
	ClientSecret string
}

// Basic contains user and password for Basic Auth
type Basic struct {
	// Username to use for authentication.
	Username string
	// Password to use for authentication.
	Password string
}

// CSRFInfo contains data for performing CSRF token request
type CSRFInfo struct {
	TokenEndpointURL string
}

// RequestParameters contains additional headers and query parameters
type RequestParameters struct {
	// Additional headers
	Headers *map[string][]string `json:"headers"`
	// Additional query parameters
	QueryParameters *map[string][]string `json:"queryParameters"`
}
