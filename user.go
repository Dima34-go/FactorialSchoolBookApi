package todo

type UserAuth struct{
	UserId int `json:"user_id"`
	Role string `json:"role"`
}

type User struct {
	Id int `json:"-" db:"idученика"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Teacher struct {
	Id int `json:"-" db:"idПреподавателя"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//binding:"required"
type Course struct{
	Id int `json:"id" db:"Idкурса"`
	Title string `json:"title" db:"название"`
	Description string `json:"description" db:"описание"`
}

type Lesson struct {
	Id int `json:"id" db:"Idзанятия"`
	Title string `json:"title" db:"название"`
	Description string `json:"description" db:"описание"`
	StatusLesson string `json:"status_lesson" db:"СтатусЗанятия"`
	AvailableForLearner bool `json:"available_for_learner" db:"ДоступностьДляУченика"`
	AvailableForTeacher bool `json:"available_for_teacher" db:"ДоступностьДляПреподавателя"`
}

type HomeTask struct{
	DeliveryClosingDate string `json:"delivery_closing_date" db:"ДатаЗакрытияПриема"`
	PublicationClosingDate string `json:"publication_closing_date" db:"ДатаЗакрытияПубликации"`
	IsPublished bool `json:"is_published" db:"ЯвляетсяВыложенным"`
}
type Homework struct{
	Comment string `json:"comment" db:"Комментарий"`
	Scores1Task int `json:"scores_1_task" db:"Баллы1Задание"`
	Scores2Task int  `json:"scores_2_task" db:"Баллы2Задание"`
	AddScores int `json:"add_scores" db:"ДопБаллы"`
	FullScores int `json:"full_scores" db:"ВсегоБаллов"`
	SendingStatus bool `json:"sending_status" db:"СтатусПроверки"`
	VerificationStatus bool `json:"verification_status" db:"СтатусОтправки"`
}
type LearnerHomework struct{
	LearnerId int `json:"learner_id" db:"idУченика"`
	Homework
}
type LearnerStatusAtLesson struct{
	AttendanceAtLesson bool `json:"attendance_at_lesson" db:"ПрисутствиеНаЗанятии"'`
	RaisedHand bool `json:"raised_hand" db:"ПоднятаяРука"`
}
type LearnerAttendance struct{
	LearnerId int `json:"learner_id" db:"idУченика"`
	LearnerStatusAtLesson
}
type LearnerStatusAtLessonForTeacher struct{
	LearnerId string `json:"learner_id" db:"idУченика"`
	Name string `json:"name" db:"Имя"`
	Surname string  `json:"surname" db:"Фамилия"`
	LearnerStatusAtLesson
	Homework
}
type NowLesson struct{
	Number int `json:"number" db:"НомерЗанятияКурса"`
}
type NextLesson struct{
	Number int `json:"number" db:"СледующееЗанятие"`
}