package models

type Course struct {
	ID               int64 `gorm:"not null;unique;primary_key"`
	Title            int64 `gorm:"size:255"`
	CourseCategoryID int64 `gorm:"index"`
	CourseCategory   CourseCategory
}