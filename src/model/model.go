package model

import "time"

type Book struct {
	Id string `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Year int `json:"year"`
	Author string `json:"author"`
	Summary string `json:"summary"`
	Publisher string `json:"publisher"`
	PageCount int `json:"pageCount"`
	ReadPage int `json:"readPage"`
	Reading bool `json:"reading"`
	Finished bool `json:"finished"`
	InsertedAt time.Time `json:"insertedAt"`
	UpdateAt time.Time `json:"updatedAt"`
}

type Input struct {
	Name string `json:"name"`
	Year int `json:"year"`
	Author string `json:"author"`
	Summary string `json:"summary"`
	Publisher string `json:"publisher"`
	PageCount int `json:"pageCount"`
	ReadPage int `json:"readPage"`
	Reading bool `json:"reading"`
}