package models

import (
	"fmt"
	"time"
)

type Course struct {
	CourseName       string
	LessonCode       string
	CourseCode       string
	PresentationDay  time.Weekday
	PresentationTime time.Time
	Teacher          *Teacher
	CourseCapacity   int
	ExamDateTime     time.Time
}

// TODO!
func (c Course) String() string {
	return fmt.Sprintf("CourseName : %s		CourseCode : %s\n", c.CourseName, c.CourseCode)
}
