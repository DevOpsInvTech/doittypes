package doittypes

import "time"

type Domain struct {
	ID        int       `sql:"not null;unique;AUTO_INCREMENT" json:"id"`
	Name      string    `sql:"unique" json:"name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
