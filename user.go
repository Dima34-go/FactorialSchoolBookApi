package todo

type User struct {
	Id int `json:"-" db:"idученика"`
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
}