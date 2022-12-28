package models

type TaskExamples []TaskExample

type TaskExample []interface{}

type TaskResults [][]int

type ServiceResponse struct {
	User    string  `json:"user_name"`
	Task    string  `json:"task"`
	Results Results `json:"results"`
}

type Results struct {
	Payload TaskExamples `json:"payload"`
	Results TaskResults  `json:"results"`
}
