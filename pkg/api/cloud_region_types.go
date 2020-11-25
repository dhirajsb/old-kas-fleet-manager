package api

type CloudRegion struct {
	Kind          string `json:"kind"`
	Id            string `json:"id"`
	DisplayName   string `json:"display_name"`
	CloudProvider string `json:"cloud_provider"`
	Enabled       bool   `json:"enabled"`
}

type CloudRegionList *[]CloudRegion
