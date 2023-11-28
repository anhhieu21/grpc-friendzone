package req

import "time"

type MovieRequest struct {
	ID        string
	Title     string
	Genre     string
	UpdatedAt time.Time
}
