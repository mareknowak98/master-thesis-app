package cmd

type LessonSlide struct {
	LessonId        string `json:"lessonId"`
	SlideId         string `json:"slideId"`
	SlideType       string `json:"slideType"`
	SlideContent    string `json:"slideContent"`
	QuestionAnswers string `json:"questionAnswers"`
}
