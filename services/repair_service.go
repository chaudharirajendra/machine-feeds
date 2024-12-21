package services

import "machine-feeds/models"

var mockRepairs = []models.Repair{
	{ID: "1", MachineID: "12345", Details: "Repair #1", Timestamp: "2024-01-02T19:35:00Z"},
	{ID: "2", MachineID: "12345", Details: "Repair #2", Timestamp: "2024-01-02T23:52:00Z"},
}

func FetchRepairs(machineID string) ([]models.Repair, error) {
	var repairs []models.Repair
	for _, repair := range mockRepairs {
		if repair.MachineID == machineID {
			repairs = append(repairs, repair)
		}
	}
	return repairs, nil
}
