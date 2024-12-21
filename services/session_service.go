package services

import "machine-feeds/models"

var mockSessions = []models.Session{
	{ID: "1", MachineID: "12345", Details: "Session #1", Timestamp: "2024-01-02T19:02:00Z"},
	{ID: "2", MachineID: "12345", Details: "Session #2", Timestamp: "2024-01-02T20:00:00Z"},
	{ID: "3", MachineID: "12345", Details: "Session #3", Timestamp: "2024-01-02T22:00:00Z"},
}

func FetchSessions(machineID string) ([]models.Session, error) {
	var sessions []models.Session
	for _, session := range mockSessions {
		if session.MachineID == machineID {
			sessions = append(sessions, session)
		}
	}
	return sessions, nil
}
