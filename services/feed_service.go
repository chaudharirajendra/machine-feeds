package services

import (
	"machine-feeds/models"
	"sort"
	"time"
)

func PrepareFeed(machineID string) ([]models.FeedItem, error) {
	feed := []models.FeedItem{}

	repairs, err := FetchRepairs(machineID)
	if err != nil {
		return nil, err
	}
	for _, repair := range repairs {
		feed = append(feed, models.FeedItem{
			Type:      "repair",
			Details:   repair.Details,
			Timestamp: repair.Timestamp,
		})
	}

	sessions, err := FetchSessions(machineID)
	if err != nil {
		return nil, err
	}
	for _, session := range sessions {
		feed = append(feed, models.FeedItem{
			Type:      "session",
			Details:   session.Details,
			Timestamp: session.Timestamp,
		})
	}

	sort.Slice(feed, func(i, j int) bool {
		t1, _ := time.Parse(time.RFC3339, feed[i].Timestamp)
		t2, _ := time.Parse(time.RFC3339, feed[j].Timestamp)
		return t1.After(t2)
	})

	return feed, nil
}

func FormatFeed(feed []models.FeedItem) []map[string]string {
	formattedFeed := []map[string]string{}
	for _, item := range feed {
		t, _ := time.Parse(time.RFC3339, item.Timestamp)
		formattedFeed = append(formattedFeed, map[string]string{
			"time":    t.Format("15:04"),
			"details": item.Details,
		})
	}
	return formattedFeed
}
