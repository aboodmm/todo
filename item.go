package todo

import "time"

type Item struct {
	dateAdded time.Time
	message   string
}
