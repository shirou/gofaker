package gofaker

import "strings"

type Internet struct {
	Domain string
	Email string
	Url string
}

func NewInternet(locale string) Internet{
	b := Internet{}

	b.Domain = GetValue("base", "top_level_domains")
	name := strings.ToLower(GetValue("base", "last_names"))
	companyname := strings.ToLower(GetValue("base", "last_names"))
	b.Email = name + "@" + companyname + "." + b.Domain
	b.Url = "http://" + companyname + "." + b.Domain

	return b
}
