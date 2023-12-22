package model

// A HealthzResponse expresses health check message.
type HealthzResponse struct {
	//シリアライズ
	Message string `json:"message"`
}
