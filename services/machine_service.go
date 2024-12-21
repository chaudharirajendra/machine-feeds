package services

import (
	"fmt"
	"machine-feeds/models"
)

var mockMachines = []models.Machine{
	{ID: "12345", Name: "Machine A", Location: "Floor 1", Details: "High-performance machine."},
	{ID: "67890", Name: "Machine B", Location: "Floor 2", Details: "Standard efficiency machine."},
}

func FetchMachine(machineID string) (*models.Machine, error) {
	for _, machine := range mockMachines {
		if machine.ID == machineID {
			return &machine, nil
		}
	}
	return nil, fmt.Errorf("machine not found")
}
