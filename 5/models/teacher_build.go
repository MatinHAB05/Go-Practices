package models

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type TeacherBuilder struct {
	Teacher
}

func NewTeacherBuilder() *TeacherBuilder {
	return &TeacherBuilder{}
}

func (teach_b *TeacherBuilder) SetFullName(Name, LastName string) *TeacherBuilder {
	teach_b.TeacherName = Name
	teach_b.TeacherLastName = LastName
	return teach_b

}

func (teach_b *TeacherBuilder) SetNationalCode(Code string) *TeacherBuilder {
	_, err := GetByNationalCode(Code)
	if err == nil {
		fmt.Println("[-][SetNatGetByNationalCode] DuplicateNatGetByNationalCode!")
		return nil
	}

	teach_b.NationalCode = Code
	return teach_b

}

func (teach_b *TeacherBuilder) SetPhoneNumber(Phone string) *TeacherBuilder {
	_, err := GetByPhoneNumber(Phone)
	if err == nil {
		fmt.Println("[-][SetPhoneNumber] DuplicatePhoneNumber!")
		return nil
	}

	teach_b.PhoneNumber = Phone
	return teach_b

}

func (teach_b *TeacherBuilder) SetPassword(Password string) *TeacherBuilder {
	// hash := sha256.Sum256([]byte(Password))
	// teach_b.Password = hex.EncodeToString(hash[:])
	salt, err := GenerateSalt(97)
	if err != nil {
		return nil
	}
	teach_b.Password = HashPassword(Password, salt)
	teach_b.Salt = salt
	return teach_b

}

func (teach_b *TeacherBuilder) SetExpertedFields(ExpertedFields ...string) *TeacherBuilder {
	copy(teach_b.ExpertedFields, ExpertedFields)
	return teach_b
}

func (teach_b *TeacherBuilder) SetTeachingCourseList(TeachingCourseList ...*Course) *TeacherBuilder {
	copy(teach_b.TeachingCourseList, TeachingCourseList)
	return teach_b

}

func (teach_b *TeacherBuilder) AddExpertedFields(ExpertedFields ...string) *TeacherBuilder {
	index := make(map[string]bool)
	for _, e := range teach_b.ExpertedFields {
		index[e] = true
	}
	for _, e := range ExpertedFields {
		if !index[e] {
			teach_b.ExpertedFields = append(teach_b.ExpertedFields, e)
		}
	}
	return teach_b
}

func (teach_b *TeacherBuilder) AddTeachingCourseList(TeachingCourseList ...Course) *TeacherBuilder {
	index := make(map[Course]bool)
	for _, e := range teach_b.TeachingCourseList {
		index[*e] = true
	}
	for _, e := range TeachingCourseList {
		if !index[e] {
			cp := e
			teach_b.TeachingCourseList = append(teach_b.TeachingCourseList, &cp)
		}
	}
	return teach_b

}

func (b *TeacherBuilder) Build() *Teacher {
	t := b.Teacher
	return &t
}

func GenerateSalt(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func HashPassword(password string, salt string) string {
	data := password + salt

	hash := sha256.Sum256([]byte(data))

	return hex.EncodeToString(hash[:])
}

func (t *Teacher) HashPassword() string {
	data := t.Password + t.Salt

	hash := sha256.Sum256([]byte(data))

	return hex.EncodeToString(hash[:])
}
