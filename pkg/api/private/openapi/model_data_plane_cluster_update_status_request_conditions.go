/*
 * Managed Service API
 *
 * Managed Service API
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// DataPlaneClusterUpdateStatusRequestConditions struct for DataPlaneClusterUpdateStatusRequestConditions
type DataPlaneClusterUpdateStatusRequestConditions struct {
	Type string `json:"type,omitempty"`
	Reason string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
	Status string `json:"status,omitempty"`
	LastTransitionTime string `json:"lastTransitionTime,omitempty"`
}
