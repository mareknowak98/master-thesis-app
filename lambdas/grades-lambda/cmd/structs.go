package cmd

type InputGrades struct {
	UserId    string `json:"UserId"`
	Date      string `json:"Date"`
	ClassYear string `json:"ClassYear"`
	Subject   string `json:"Subject"`
	Grade     string `json:"Grade"`
}
