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

type CheckTestInput struct {
	TestId         string           `json:"testId"`
	QuestionAnswer []QuestionAnswer `json:"questionAnswer"`
}

type QuestionAnswer struct {
	CombinedKey           string `json:"combinedKey"`
	CorrectQuestionAnswer string `json:"correctQuestionAnswer"`
}

type TestResult struct {
	TestId string `json:"testId"`
	UserId string `json:"userId"`
	Result string `json:"result"`
	//QuestionAnswers string `json:"questionAnswers"`
}
