/*
 * Managed Service API
 *
 * Managed Service API
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ConnectorClusterUpdateStatusResizeInfo struct for ConnectorClusterUpdateStatusResizeInfo
type ConnectorClusterUpdateStatusResizeInfo struct {
	NodeDelta int32                                       `json:"nodeDelta,omitempty"`
	Delta     ConnectorClusterUpdateStatusResizeInfoDelta `json:"delta,omitempty"`
}