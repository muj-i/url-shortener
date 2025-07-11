package model

import "time"

type URL struct {
	Code       string
	LongURL    string
	ClickCount int
	ExpiresAt  *time.Time
	CreatedAt  time.Time
}
