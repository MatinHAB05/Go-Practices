package models

import (
	"encoding/json"
	"fmt"
	"os"
)

var CourseDB []*Course

func AddCourseToDB(c *Course) {
	CourseDB = append(CourseDB, c)
}

func indexOfCourse[T comparable](key T, selector func(*Course) T) (*Course, error) {
	for i := range CourseDB {
		t := CourseDB[i]
		if selector(t) == key {
			return t, nil
		}
	}
	return nil, fmt.Errorf("not found")
}

func GetByCourseCode(key string) (*Course, error) {
	return indexOfCourse(key, func(c *Course) string {
		return c.CourseCode
	})
}

func LoadCourse(Path string) error {
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

		err = dec.Decode(&CourseDB)
		if err != nil {
			return err
		}

		return nil
	}

}

func SaveCourse(Path string) error {

	writer, err := os.Create(Path)
	if err != nil {
		return err
	}
	defer writer.Close()

	enc := json.NewEncoder(writer)
	enc.SetIndent("", "  ")

	enc.Encode(CourseDB)

	return nil
}
