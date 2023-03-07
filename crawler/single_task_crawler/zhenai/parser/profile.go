package parser

import (
	"crawler/engine"
	"crawler/model"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<td width="180"><span class="grayL">年龄：</span>([\d]+)</td>`)
var genderRe = regexp.MustCompile(`<td width="180"><span class="grayL">性别：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="grayL">学   历：</span>([^<]+)</td>`)
var incomeRe = regexp.MustCompile(`<td><span class="grayL">月   薪：</span>([^<]+)</td>`)
var marriageRe = regexp.MustCompile(`<td width="180"><span class="grayL">婚况：</span>([^<]+)</td>`)
var heightRe = regexp.MustCompile(`<td width="180"><span class="grayL">身   高：</span>([0-9]+)</td>`)
var nameRe = regexp.MustCompile(`<a href="http://album.zhenai.com/u/[0-9]+" target="_blank">([^<]+)</a>`)

func ParseProfile(bytes []byte) engine.ParseResult {
	var profile model.Profile
	name := string(extractString(bytes, nameRe)[1])
	fmt.Println(name)
	profile.Name = name
	age := string(extractString(bytes, ageRe)[1])
	fmt.Println(age)
	profile.Age, _ = strconv.Atoi(age)
	height, _ := strconv.Atoi(string(extractString(bytes, heightRe)[1]))
	fmt.Println(height)
	profile.Height = height
	gender := string(extractString(bytes, genderRe)[1])
	fmt.Println(gender)
	profile.Gender = gender
	education := string(extractString(bytes, educationRe)[1])
	fmt.Println(education)
	profile.Education = education
	income := string(extractString(bytes, incomeRe)[1])
	fmt.Println(income)
	profile.Income = income
	marriage := string(extractString(bytes, marriageRe)[1])
	fmt.Println(marriage)
	profile.Marriage = marriage

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	if err := (&profile).CreateUser(); err != nil {
		result.Items = []interface{}{errors.New("入库失败")}
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) [][]byte {
	submatch := re.FindSubmatch(contents)
	if len(submatch) >= 2 {
		return submatch
	} else {
		return nil
	}
}
