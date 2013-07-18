package gofaker

import (
	"fmt"
	"testing"
	"strings"
	"strconv"
)

func TestNewFaker(t *testing.T){
	f1, _ := NewFaker("ja_JP")
	fmt.Println(f1.name.Phone)
	fmt.Println(f1.name.Last_name, f1.name.Last_name_alt)
	fmt.Println(f1.name.First_name, f1.name.First_name_alt)
	fmt.Println(f1.name.Title)

	f2, _ := NewFaker("en_US")
	fmt.Println(f2.name.Phone)
	fmt.Println(f2.name.Last_name)
	fmt.Println(f2.name.First_name)
}

func TestCompany(t *testing.T){
	f1, _ := NewFaker("ja_JP")
	fmt.Println(f1.company.Name)
}

func TestInternet(t *testing.T){
	f1, _ := NewFaker("ja_JP")
	fmt.Println(f1.internet.Url)
	fmt.Println(f1.internet.Domain)
	fmt.Println(f1.internet.Email)
}

func TestAddress(t *testing.T){
	f1, _ := NewFaker("ja_JP")
	fmt.Println(f1.address.City)
}


func TestReplace(t *testing.T){
	a, _ := strconv.Atoi(Replace("#"))
	if a > 10 {
		t.Errorf("replace failed")
	}

	// if Upper and Lower is same, this is number. any other good solution?
	b := Replace("###")
	if b != strings.ToLower(b){
		t.Errorf("replace failed")
	}
	if b != strings.ToUpper(b){
		t.Errorf("replace failed")
	}

}
