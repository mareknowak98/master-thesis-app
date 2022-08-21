package cmd

type LessonSlide struct {
	LessonId        int    `json:"lessonId"`
	SlideId         int    `json:"slideId"`
	SlideType       string `json:"slideType"`
	SlideContent    string `json:"slideContent"`
	QuestionAnswers string `json:"questionAnswers"`
}
