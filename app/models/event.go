package models

import "encoding/json"

//go:generate easyjson

type Message interface {
	json.Marshaler
}

//easyjson:json
type Event struct {
	Name string `json:"name"`
	Data string `json:"data"`
}
