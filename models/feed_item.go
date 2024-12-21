package models

type FeedItem struct {
	Type      string `json:"type"` // "repair" or "session"
	Details   string `json:"details"`
	Timestamp string `json:"timestamp"`
}
