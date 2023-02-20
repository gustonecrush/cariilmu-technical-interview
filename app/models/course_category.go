package models 

type CourseCategory struct {
	ID      int64  `gorm:"not null;unique;primary_key"`
	Name    string `gorm:"size:255;unique;"`
	Courses []Course
}