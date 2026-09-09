package main

import (
	"fmt"
	"lets_go/models"
	"slices"
	"time"
)

func main() {
	Menu()
}

func Menu() {
	fmt.Println("Load Json...")
	models.LoadCourse("db/course.json")
	models.LoadTeacher("db/teacher.json")
	saved := true
	for {
		if !saved {
			fmt.Println("Save To Json...")
			models.SaveCourse("db/course.json")
			models.SaveTeacher("db/teacher.json")
		}
		fmt.Println("\n1-Add Ostad")
		fmt.Println("2-Edit")
		fmt.Println("3-ListOstad")
		fmt.Println("4-ListCourses")
		fmt.Println("5-Create Courses")
		fmt.Println("6-Add Courses")
		fmt.Println("7-Remove Courses")
		fmt.Println("8-Print ListCourses")
		fmt.Println("97-Exit\n\n")

		var choice int
		fmt.Print("Enter choice: ")
		fmt.Scanln(&choice)

		switch choice {

		case 1:
			AddOstad()

		case 2:
			Edit()

		case 3:
			ListOstad()

		case 4:
			ListCourses()

		case 5:
			CreateCourses()

		case 6:
			AddCourses()

		case 7:
			RemoveCourses()

		case 8:
			fmt.Println(models.CourseDB)
		case 97:
			fmt.Println("Exit...")
			return

		default:
			fmt.Println("Invalid option")
		}
		saved = false
	}
}

func AddOstad() {

	var t models.Teacher

	fmt.Print("Name: ")
	fmt.Scanln(&t.TeacherName)

	fmt.Print("LastName: ")
	fmt.Scanln(&t.TeacherLastName)

	fmt.Print("NationalCode: ")
	fmt.Scanln(&t.NationalCode)

	_, err := models.GetByNationalCode(t.NationalCode)
	if err == nil {
		fmt.Println("Duplicated NationalCode!")
		return
	}

	fmt.Print("PhoneNumber: ")
	fmt.Scanln(&t.PhoneNumber)

	_, err = models.GetByPhoneNumber(t.PhoneNumber)
	if err == nil {
		fmt.Println("Duplicated Phone!")
		return
	}

	fmt.Print("Password: ")
	fmt.Scanln(&t.Password)
	salt, err := models.GenerateSalt(97)
	if err != nil {
		return
	}
	t.Password = models.HashPassword(t.Password, salt)
	t.Salt = salt

	models.AddTeacherToDB(&t)

	fmt.Println("Teacher added successfully")
}

func Edit() {

	var nationalCode string
	var password string

	fmt.Print("Enter national code: ")
	fmt.Scanln(&nationalCode)

	fmt.Print("Enter password: ")
	fmt.Scanln(&password)

	t, err := models.GetByNationalCode(nationalCode)
	hashString := models.HashPassword(password, t.Salt)

	if err != nil || t.Password != hashString {
		fmt.Println("National code or password is incorrect")
		return
	}

	var field string
	fmt.Println("Which field do you want to edit?")
	fmt.Println("1: Name")
	fmt.Println("2: NationalCode")
	fmt.Println("3: Password")
	fmt.Println("4: PhoneNumber")
	fmt.Println("5: Expertise")

	fmt.Scanln(&field)

	switch field {

	case "1":
		var name string
		var lastname string

		fmt.Print("Enter new name (or - to skip): ")
		fmt.Scanln(&name)

		fmt.Print("Enter new lastname (or - to skip): ")
		fmt.Scanln(&lastname)

		if name != "-" {
			t.TeacherName = name
		}
		if lastname != "-" {
			t.TeacherLastName = lastname
		}

		fmt.Println("Name updated successfully")

	case "2":
		var newCode string
		fmt.Print("Enter new national code: ")
		fmt.Scanln(&newCode)

		t, err := models.GetByNationalCode(newCode)
		if err == nil {
			fmt.Println("Duplicated NationalCode!")
			return
		}

		t.NationalCode = newCode
		fmt.Println("National code updated")

	case "3":
		var newPass string
		fmt.Print("Enter new password: ")
		fmt.Scanln(&newPass)
		salt, err := models.GenerateSalt(97)
		if err != nil {
			return
		}
		t.Password = models.HashPassword(newPass, salt)
		t.Salt = salt

		fmt.Println("Password updated")

	case "4":
		var phone string
		fmt.Print("Enter new phone number: ")
		fmt.Scanln(&phone)

		t, err := models.GetByPhoneNumber(phone)
		if err == nil {
			fmt.Println("Duplicated Phone!")
			return
		}

		t.PhoneNumber = phone
		fmt.Println("Phone number updated")

	case "5":
		var action string
		var field string

		fmt.Print("Add or Remove: ")
		fmt.Scanln(&action)

		fmt.Print("Enter expertise: ")
		fmt.Scanln(&field)

		if action == "Add" {
			ind := slices.Index(t.ExpertedFields, field)
			if ind < 0 {
				t.ExpertedFields = append(t.ExpertedFields, field)
			}
			fmt.Println("ExpertedFields added")

		} else if action == "Remove" {

			for i, v := range t.ExpertedFields {
				if v == field {
					t.ExpertedFields = append(t.ExpertedFields[:i], t.ExpertedFields[i+1:]...)
					fmt.Println("Expertise removed")
					return
				}
			}

			fmt.Println("Expertise not found")
		}

	default:
		fmt.Println("Invalid option")
	}
}

func ListOstad() {

	if len(models.TeacherDB) == 0 {
		fmt.Println("No teachers")
		return
	}

	for i, t := range models.TeacherDB {
		fmt.Printf("%d - %s %s | NationalCode: %s\n",
			i+1,
			t.TeacherName,
			t.TeacherLastName,
			t.NationalCode,
		)
	}
}

func CreateCourses() {

	var c models.Course

	fmt.Print("Course Name: ")
	fmt.Scanln(&c.CourseName)

	fmt.Print("Lesson Code: ")
	fmt.Scanln(&c.LessonCode)

	fmt.Print("Course Code: ")
	fmt.Scanln(&c.CourseCode)
	_, err := models.GetByCourseCode(c.CourseCode)
	if err == nil {
		fmt.Println("Duplicated CourseCode!")
		return
	}

	var day int
	fmt.Print("Presentation day (0=Sunday ... 6=Saturday): ")
	fmt.Scanln(&day)
	c.PresentationDay = time.Weekday(day)

	var timeStr string
	fmt.Print("Presentation time (HH:MM): ")
	fmt.Scanln(&timeStr)

	t, err := time.Parse("15:04", timeStr)
	if err != nil {
		fmt.Println("invalid time")
		return
	}
	c.PresentationTime = t

	fmt.Print("Capacity: ")
	fmt.Scanln(&c.CourseCapacity)
	if c.CourseCapacity <= 0 {
		fmt.Println("Invalid CourseCap!")
		return
	}

	var examStr string
	fmt.Print("Exam datetime (YYYY-MM-DD HH:MM): ")
	fmt.Scanln(&examStr)

	examTime, err := time.Parse("2006-01-02 15:04", examStr)
	if err != nil {
		fmt.Println("invalid exam time")
		return
	}
	c.ExamDateTime = examTime

	models.AddCourseToDB(&c)
	fmt.Println("Course created")
}

func AddCourses() {

	var nationalCode string
	fmt.Print("Enter teacher national code: ")
	fmt.Scanln(&nationalCode)

	var teacher *models.Teacher

	teacher, err := models.GetByNationalCode(nationalCode)

	if err != nil {
		fmt.Println("Teacher not found")
		return
	}

	var courseCode string
	fmt.Print("Enter course code: ")
	fmt.Scanln(&courseCode)

	c, err_ := models.GetByCourseCode(courseCode)
	if err_ != nil {
		fmt.Println("Course not found")
		return

	}
	if slices.Index(teacher.TeachingCourseList, c) < 0 && slices.Index(teacher.ExpertedFields, c.CourseName) >= 0 {
		teacher.TeachingCourseList = append(teacher.TeachingCourseList, c)
	}
	fmt.Println("Course added to teacher")
	return

}

func RemoveCourses() {

	var nationalCode string
	fmt.Print("Enter teacher national code: ")
	fmt.Scanln(&nationalCode)

	var teacher *models.Teacher

	teacher, err := models.GetByNationalCode(nationalCode)

	if err != nil {
		fmt.Println("Teacher not found")
		return
	}

	var courseCode string
	fmt.Print("Enter course code: ")
	fmt.Scanln(&courseCode)

	c, err_ := models.GetByCourseCode(courseCode)
	if err_ != nil {
		fmt.Println("Course not found at all")
		return

	}
	index := slices.Index(teacher.TeachingCourseList, c)
	if index < 0 {
		fmt.Println("Course not found for this teacher")
		return
	}

	teacher.TeachingCourseList = append(teacher.TeachingCourseList[:index], teacher.TeachingCourseList[index+1:]...)
	fmt.Println("Course removed from teacher")
	return

}

func ListCourses() {

	var nationalCode string
	fmt.Print("Enter teacher national code: ")
	fmt.Scanln(&nationalCode)

	var teacher *models.Teacher

	teacher, err := models.GetByNationalCode(nationalCode)

	if err != nil {
		fmt.Println("Teacher not found")
		return
	}

	if len(teacher.TeachingCourseList) == 0 {
		fmt.Println("This teacher has no courses")
		return
	}

	fmt.Println(teacher.TeachingCourseList)
}
