package models

import (
	"fmt"
)

type Teacher struct {
	TeacherName        string
	TeacherLastName    string
	NationalCode       string
	PhoneNumber        string
	Password           string
	Salt               string
	ExpertedFields     []string
	TeachingCourseList []*Course
}

func (t Teacher) String() string {
	return fmt.Sprintf("TeacherName : %s | TeacherLastName : %s", t.TeacherName, t.TeacherLastName)
}
