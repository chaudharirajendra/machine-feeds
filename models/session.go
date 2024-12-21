package models

type Session struct {
	ID        string `json:"id"`
	MachineID string `json:"machine_id"`
	Details   string `json:"details"`
	Timestamp string `json:"timestamp"`
}
