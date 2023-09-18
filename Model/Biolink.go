package model

import ()

type BioLink struct {
	Id    string `json:"id,omitempty`
	Title string `json:"title" validate:"required"`
	Link  string `json:"link" validate:"required"`
}
