package models

type Timesheet struct {
	Email     string `json:"email"`
	Token     string `json:"token"`
	Clock_in  string `json:"clock_in"`
	Clock_out string `json:"clock_out"`
}
