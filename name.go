package gofaker

import (
	"strings"
)

type Name struct {
	First_name string
	First_name_alt string
	Last_name string
	Last_name_alt string
	Phone  string
	Gender string
	Title  string
}

func NewName(locale string) Name{
	b := Name{}

	b.Gender = GetValue(locale, "genders")
	b.Phone = GetValue(locale, "phone")
	b.Title = GetValue(locale, "name_titles")
	b.Last_name = GetValue(locale, "last_names")
	if strings.Contains(b.Last_name , "|") {
		t := strings.Split(b.Last_name, "|")
		b.Last_name = t[0]
		b.Last_name_alt = t[1]
	}
	if b.Gender == "male" {
		b.First_name = GetValue(locale, "male_first_names")
	}else{
		b.First_name = GetValue(locale, "female_first_names")
	}
	if strings.Contains(b.First_name , "|") {
		t := strings.Split(b.First_name, "|")
		b.First_name = t[0]
		b.First_name_alt = t[1]
	}

	return b
}
