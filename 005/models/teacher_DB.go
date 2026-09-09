package models

import (
	"encoding/json"
	"fmt"
	"os"
)

var TeacherDB []*Teacher

func AddTeacherToDB(t *Teacher) {
	TeacherDB = append(TeacherDB, t)
}

func indexOfTeacher[T comparable](key T, selector func(*Teacher) T) (*Teacher, error) {
	for i := range TeacherDB {
		t := TeacherDB[i]
		if selector(t) == key {
			return t, nil
		}
	}
	return nil, fmt.Errorf("not found")
}

func GetByNationalCode(key string) (*Teacher, error) {
	return indexOfTeacher(key, func(t *Teacher) string {
		return t.NationalCode
	})
}

func GetByPhoneNumber(key string) (*Teacher, error) {
	return indexOfTeacher(key, func(t *Teacher) string {
		return t.PhoneNumber
	})
}

func LoadTeacher(Path string) error {
	reader, err := os.Open(Path)
	if err != nil {
		temp, err_ := os.Create(Path)
		if err_ != nil {
			return err_
		}
		temp.Close()
		return nil
	} else {
		defer reader.Close()
		dec := json.NewDecoder(reader)

		err = dec.Decode(&TeacherDB)
		if err != nil {
			return err
		}

		return nil
	}

}

func SaveTeacher(Path string) error {

	writer, err := os.Create(Path)
	if err != nil {
		return err
	}
	defer writer.Close()

	enc := json.NewEncoder(writer)
	enc.SetIndent("", "  ")
	enc.Encode(TeacherDB)

	return nil
}
