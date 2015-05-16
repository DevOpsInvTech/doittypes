package doittypes

import "time"

//Domain a container for all related objects under the same namespace
type Domain struct {
	ID        int       `sql:"not null;unique;AUTO_INCREMENT" json:"id"`
	Name      string    `sql:"unique" json:"name"`
	Hosts     []*Host   `json:"hosts,omitempty"`
	Vars      []*Var    `json:"vars,omitempty"`
	Groups    []*Group  `json:"groups,omitempty"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
