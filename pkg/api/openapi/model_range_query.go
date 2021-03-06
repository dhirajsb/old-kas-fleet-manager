/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage kafka instances and connectors.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// RangeQuery struct for RangeQuery
type RangeQuery struct {
	Metric map[string]string `json:"metric,omitempty"`
	Values []Values          `json:"values,omitempty"`
}
