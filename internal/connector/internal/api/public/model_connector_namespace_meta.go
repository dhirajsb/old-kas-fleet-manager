/*
 * Connector Service Fleet Manager
 *
 * Connector Service Fleet Manager is a Rest API to manage connectors.
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package public

import (
	"time"
)

// ConnectorNamespaceMeta struct for ConnectorNamespaceMeta
type ConnectorNamespaceMeta struct {
	Owner       string                                     `json:"owner,omitempty"`
	CreatedAt   time.Time                                  `json:"created_at,omitempty"`
	ModifiedAt  time.Time                                  `json:"modified_at,omitempty"`
	Name        string                                     `json:"name"`
	Annotations []ConnectorNamespaceRequestMetaAnnotations `json:"annotations,omitempty"`
}
