/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage kafka instances and connectors.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// MetricsInstantQueryList struct for MetricsInstantQueryList
type MetricsInstantQueryList struct {
	Kind  string         `json:"kind,omitempty"`
	Id    string         `json:"id,omitempty"`
	Items []InstantQuery `json:"items,omitempty"`
}
