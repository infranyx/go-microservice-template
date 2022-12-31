package healthCheckDto

type HealthCheckUnit struct {
	Unit string `json:"unit"`
	Up   bool   `json:"up"`
}

type HealthCheckResponseDto struct {
	Status bool              `json:"status"`
	Units  []HealthCheckUnit `json:"units"`
}
