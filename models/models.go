package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Ads struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	StartAt   string    `json:"startAt"`
	EndAt     string    `json:"endAt"`
	Condition Condition `json:"condition"`
}

type Condition struct {
	AgeStart *int     `json:"ageStart"`
	AgeEnd   *int     `json:"ageEnd"`
	Gender   string   `json:"gender"`
	Country  []string `json:"country"`
	Platform []string `json:"platform"`
}

// 定義Ad結構
type Ad struct {
	ID       int            `json:"id"`
	Title    string         `json:"title"`
	StartAt  string         `json:"startAt"`
	EndAt    string         `json:"endAt"`
	AgeStart sql.NullInt64  `json:"ageStart"`
	AgeEnd   sql.NullInt64  `json:"ageEnd"`
	Gender   sql.NullString `json:"gender"`
	Country  sql.NullString `json:"country"`
	Platform sql.NullString `json:"platform"`
}
