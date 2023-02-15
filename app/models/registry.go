package models

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model {
		{Model: Admin{}},
		{Model: User{}},
		{Model: Course{}},
		{Model: CourseCategory{}},
	}
}