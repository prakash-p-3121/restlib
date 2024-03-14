package model

type IDResponse struct {
	ResourceID string `json:"resource-id"`
}

type MsConnectionCfg struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}
