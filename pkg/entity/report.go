package entity

import "time"

type Report struct {
	Date      time.Time
	Content   string
	Students  []int
	CreatedAt time.Time
}
