package main

// Result of the quiz for a certain user
type Result struct {
	Choices []string `json:"choices"`
	Rate    float64  `json:"successRate"`
	Correct int      `json:"correct"`
}
