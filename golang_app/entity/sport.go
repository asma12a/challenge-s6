package entity

import (
	"github.com/asma12a/challenge-s6/ent"
)

type Sport struct {
	ent.Sport
}

func NewSport(name string, imageUrl string) *Sport {
	return &Sport{
		Sport: ent.Sport{
			Name:     name,
			ImageURL: imageUrl,
		},
	}
}
