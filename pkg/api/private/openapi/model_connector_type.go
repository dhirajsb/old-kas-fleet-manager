/*
 * Managed Service API
 *
 * Managed Service API
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ConnectorType Represents a connector type supported by the API
type ConnectorType struct {
	Id   string `json:"id,omitempty"`
	Kind string `json:"kind,omitempty"`
	Href string `json:"href,omitempty"`
	// Name of the connector type.
	Name string `json:"name"`
	// Version of the connector type.
	Version string `json:"version"`
	// A description of the connector.
	Description string `json:"description,omitempty"`
	// A json schema that can be used to validate a connectors connector_spec field.
	JsonSchema map[string]interface{} `json:"json_schema,omitempty"`
}
