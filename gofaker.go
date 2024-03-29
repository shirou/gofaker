package gofaker

import (
	"os"
	"fmt"
	"path/filepath"
	"errors"
	"bufio"
	"strings"
	"strconv"
	"unicode/utf8"
	"math/rand"
	"time"
)

type Faker struct {
	Name   Name
	Address Address
	Company Company
	Internet Internet
}

// dictionary path
var dict_path = "dict"


func NewFaker(locale string, dict string) (Faker, error){
	if dict == ""{
		dict = "dict/"
	}
	if !filepath.IsAbs(dict) {
		dict, _ = filepath.Abs(filepath.Join(".", dict))
	}
	dict_path = dict

	f := Faker{}
	f.Name = NewName(locale)
	f.Address = NewAddress(locale)
	f.Company = NewCompany(locale)
	f.Internet = NewInternet(locale)

	return f, nil

}

// readLines reads contents from file and split by new line.
func readLines(filename string) ([]string, error) {
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

// replace can replace # to random number or 0-9
func replace(original string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rep_func := func(s rune) rune {
		if s == '#'{
			str := strconv.Itoa(r.Intn(9))
			rr, _ := utf8.DecodeLastRuneInString(str)
			return rr
		}
		return s
	}
	return strings.Map(rep_func, original)
}

// pickOne picks one value from a array which is read from a file.
func pickOne(filepath string) (string, error){
	s, _ := readLines(filepath)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(len(s))
	orig := s[index]

	return replace(orig), nil
}

// searchDictFile searchs a dictionary file. At first, search locale directory under the dict_path.
// if not find, search "base" directory.
// fieldName should be same as filename.
func searchDictFile(locale string, fieldName string) (string, error){
	filename := filepath.Join(dict_path, locale, fieldName)
	if _, err := os.Stat(filename); err == nil {
		return filename, nil
	}

	// a dict file is not in that locale. try to read base directory.
	filename = filepath.Join(dict_path, "base", fieldName)
	if _, err := os.Stat(filename); err == nil {
		return filename, nil
	}else{
		return "", errors.New("Could not find dictionary:" + filename)
	}
}


// GetValue returnes a value from specified locale and fieldName
// locale example: en_US, ja_JP
func GetValue(locale string, fieldName string) string{
	filename, err := searchDictFile(locale, fieldName)
	if err != nil {
		fmt.Println("Could not find dictionary file: " + filepath.Join(dict_path, fieldName) + " of " + locale)
		os.Exit(-1)
	}

	ret, err := pickOne(filename)
	if err != nil {
		fmt.Println("Could not get value")
		os.Exit(-1)
	}

	return ret
}
