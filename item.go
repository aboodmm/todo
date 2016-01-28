package todo

import "time"

type Item struct {
	DateAdded time.Time
	Message   string
}

func NewItem(m string) Item {
	i := new(Item)
	i.DateAdded = time.Now()
	i.Message = m
	return i
}
