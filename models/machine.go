package models

type Machine struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Details  string `json:"details"`
}
