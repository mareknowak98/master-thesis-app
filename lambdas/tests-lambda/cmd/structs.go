package cmd

type TestQuestionInput struct {
	TestId        string   `json:"testId"`
	CombinedKey   string   `json:"combinedKey"`
	Question      string   `json:"question"`
	Answers       []string `json:"answers"`
	CorrectAnswer string   `json:"correctAnswer"`
}

type TestQuestionOutput struct {
	TestId      string   `json:"testId"`
	CombinedKey string   `json:"combinedKey"`
	Question    string   `json:"question"`
	Answers     []string `json:"answers"`
}

type TestId struct {
	TestId string `json:"testId"`
}
