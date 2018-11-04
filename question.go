package main

// Question with a set of options
type Question struct {
	Content string    `json:"content"`
	Options [4]string `json:"options"`
}
