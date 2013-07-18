package gofaker

import (
	"os"
	"fmt"
	"path"
	"errors"
	"bufio"
	"strings"
	"strconv"
	"unicode/utf8"
	"math/rand"
	"time"
)

type Faker struct {
	name   Name
	address Address
	company Company
	internet Internet
}

const DICTDIRNAME = "dict"


func NewFaker(locale string) (Faker, error){

	fmt.Println(locale)

	f := Faker{}
	f.name = NewName(locale)
	f.address = NewAddress(locale)
	f.company = NewCompany(locale)
	f.internet = NewInternet(locale)

	return f, nil

}

// Read contents from file and split by new line.
func ReadLines(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return []string{""}, err
	}
	defer f.Close()

	ret := make([]string, 0)

	r := bufio.NewReader(f)
	line, err := r.ReadString('\n')
	for err == nil {
		ret = append(ret, strings.Trim(line, "\n"))
		line, err = r.ReadString('\n')
	}

	return ret, err
}

func Replace(original string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rep_func := func(s rune) rune {
		if s == '#'{
			str := strconv.Itoa(r.Intn(9))
			rr, _ := utf8.DecodeLastRuneInString(str)
			return rr
		}else{
			return s
		}
	}
	return strings.Map(rep_func, original)
}


func PickOne(filepath string) (string, error){
	s, _ := ReadLines(filepath)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(len(s))
	orig := s[index]

	return Replace(orig), nil
}


func SearchDictFile(locale string, fieldName string) (string, error){
	filename := path.Join(DICTDIRNAME, locale, fieldName)
	if _, err := os.Stat(filename); err == nil {
		return filename, nil
	}

	// a dict file is not in that locale. try to read base values.
	filename = path.Join(DICTDIRNAME, "base", fieldName)
	if _, err := os.Stat(filename); err == nil {
		return filename, nil
	}else{
		return "", errors.New("Could not find dictionaly:" + filename)
	}
}

func GetValue(locale string, fieldName string) string{
	filepath, err := SearchDictFile(locale, fieldName)
	if err != nil {
		fmt.Println("Could not find dictionary: file: " + fieldName + " of " + locale)
		os.Exit(-1)
	}

	ret, err := PickOne(filepath)
	if err != nil {
		fmt.Println("Could not get value")
		os.Exit(-1)
	}

	return ret
}
