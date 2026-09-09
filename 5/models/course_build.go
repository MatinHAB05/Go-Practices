package models

import (
	"fmt"
	"time"
)

type CourseBuilder struct {
	course Course
}

func NewCourseBuilder() *CourseBuilder {
	return &CourseBuilder{}
}

func (b *CourseBuilder) CourseName(name string) *CourseBuilder {
	b.course.CourseName = name
	return b
}

func (b *CourseBuilder) LessonCode(code string) *CourseBuilder {
	b.course.LessonCode = code
	return b
}

func (b *CourseBuilder) CourseCode(code string) *CourseBuilder {
	_, err := GetByCourseCode(code)
	if err == nil {
		fmt.Println("[-][CourseCode] DuplicateCourseCode!")
		return nil
	}
	b.course.CourseCode = code
	return b
}

func (b *CourseBuilder) PresentationDay(day time.Weekday) *CourseBuilder {
	b.course.PresentationDay = day
	return b
}

func (b *CourseBuilder) PresentationTime(t time.Time) *CourseBuilder {
	b.course.PresentationTime = t
	return b
}

func (b *CourseBuilder) Teacher(t *Teacher) *CourseBuilder {
	b.course.Teacher = t
	return b
}

func (b *CourseBuilder) Capacity(c int) *CourseBuilder {
	b.course.CourseCapacity = c
	return b
}

func (b *CourseBuilder) ExamDateTime(t time.Time) *CourseBuilder {
	b.course.ExamDateTime = t
	return b
}

func (b *CourseBuilder) Build() *Course {
	c := b.course
	return &c
}
