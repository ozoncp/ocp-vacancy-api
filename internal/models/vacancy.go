package models

import "fmt"

type Vacancy struct {
	ID     uint64
	Link   string
	Status uint64
}

func (v *Vacancy) String() string {
	if v == nil {
		return "Non-initialized Vacancy pointer"
	}
	return fmt.Sprintf("Vacancy ID: %d, Status: %d, Link: %s", v.ID, v.Status, v.Link)
}
