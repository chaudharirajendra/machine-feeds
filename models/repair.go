package models

type Repair struct {
	ID        string `json:"id"`
	MachineID string `json:"machine_id"`
	Details   string `json:"details"`
	Timestamp string `json:"timestamp"`
}
