package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// 定義POST廣告結構
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

// 定義GET Ad結構
type Ad struct {
	ID       int            `json:"id"`
	Title    string         `json:"title"`
	StartAt  string         `json:"startAt"`
	EndAt    string         `json:"endAt"`
	AgeStart *int           `json:"ageStart"`
	AgeEnd   *int           `json:"ageEnd"`
	Gender   sql.NullString `json:"gender"`
	Country  sql.NullString `json:"country"`
	Platform sql.NullString `json:"platform"`
}
