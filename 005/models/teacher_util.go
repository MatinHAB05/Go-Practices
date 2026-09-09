package models

import (
	"fmt"
	"slices"
)

func (tech *Teacher) ListCourses() {
	fmt.Println(tech.TeachingCourseList)
}

func (tech *Teacher) AddCourse(c *Course) {
	if slices.Contains(tech.ExpertedFields, c.CourseName) {
		if !slices.Contains(tech.TeachingCourseList, c) {
			tech.TeachingCourseList = append(tech.TeachingCourseList, c)
		}
		fmt.Printf("[+] {Course:%s} -add-> {Teacher:%s}\n", c.CourseName, tech.TeacherName+" "+tech.TeacherLastName)
	} else {
		fmt.Printf("[-] {Course:%s} <-Cant Tech-> {Teacher:%s}\n", c.CourseName, tech.TeacherName+" "+tech.TeacherLastName)
	}

}

func (tech *Teacher) RemoveCourse(CourseCode string) {
	temp := make([]string, len(tech.TeachingCourseList))
	for i := 0; i < len(tech.TeachingCourseList); i++ {
		temp = append(temp, tech.TeachingCourseList[i].CourseCode)
	}
	index := slices.Index(temp, CourseCode)
	if index < 0 {
		fmt.Printf("[-][RemoveCourse] {CourseCode:%s} <-CantFind-> {TeacherName:%s}\n", CourseCode, tech.TeacherName+" "+tech.TeacherLastName)
		return
	}
	tech.TeachingCourseList = append(tech.TeachingCourseList[0:index], tech.TeachingCourseList[index+1:]...)
	fmt.Printf("[+][RemoveCourse] {CourseCode:%s} <-> {TeacherName:%s}\n", CourseCode, tech.TeacherName+" "+tech.TeacherLastName)
}
