package parser

import (
	"bytes"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type PersonInfo struct {
	Name          string
	Id            string
	Region        string
	Age           int
	Edu           string
	MaritalStatus string
	Height        int
	Weight        int
	Salary        [2]int
}

func atoIOrDef(text string, def int) int {
	res, err := strconv.Atoi(text)
	if err == nil {
		return res
	}
	return def
}

func extractInt(text string, reg *regexp.Regexp) int {
	findRes := reg.FindSubmatch([]byte(text))
	if findRes != nil {
		return atoIOrDef(string(findRes[1]), 0)
	}
	return 0
}

var AgeRe = regexp.MustCompile(`(\d+)岁`)

func extractAge(text string) int {
	return extractInt(text, AgeRe)
}

var HeightRe = regexp.MustCompile(`(\d+)cm`)

func extractHeight(text string) int {
	return extractInt(text, HeightRe)
}

var WeightRe = regexp.MustCompile(`(\d+)kg`)

func extractWeight(text string) int {
	return extractInt(text, WeightRe)
}

var IdRe = regexp.MustCompile(`ID：(\d+)`)

func extractId(text string) string {
	findRes := IdRe.FindSubmatch([]byte(text))
	if findRes == nil {
		return ""
	}
	return string(findRes[1])
}

var SalaryRe = regexp.MustCompile(`(\d+)-(\d+)元`)

func extractSalary(text string) [2]int {
	match := SalaryRe.FindSubmatch([]byte(text))
	if match != nil {
		return [2]int{atoIOrDef(string(match[1]), 0), atoIOrDef(string(match[2]), 0)}
	}
	return [2]int{0, 0}
}

func ParsePerson(html []byte) (PersonInfo, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))

	if err != nil {
		return PersonInfo{}, err
	}
	personInfo := PersonInfo{}
	doc.Find(".nickName").Each(func(i int, selection *goquery.Selection) {
		nameStr := selection.Contents().Text()
		personInfo.Name = strings.TrimSpace(string(nameStr))
	})

	doc.Find(".id").Each(func(i int, selection *goquery.Selection) {
		idStr := selection.Contents().Text()
		personInfo.Id = extractId(idStr)
	})
	doc.Find(".des.f-cl").Each(func(i int, selection *goquery.Selection) {
		characterStr := selection.Contents().Text()
		texts := strings.Split(characterStr, "|")
		personInfo.Region = strings.TrimSpace(texts[0])
		personInfo.Age = extractAge(texts[1])
		personInfo.Edu = strings.TrimSpace(texts[2])
		personInfo.MaritalStatus = strings.TrimSpace(texts[3])
		personInfo.Height = extractHeight(texts[4])
		personInfo.Salary = extractSalary(texts[5])
	})

	doc.Find(".m-content-box > .purple-btns > div").Each(func(i int, selection *goquery.Selection) {
		curText := selection.Contents().Text()
		if WeightRe.Match([]byte(curText)) {
			personInfo.Weight = extractWeight(curText)
		}
	})
	return personInfo, nil
}

func ParsePersonRes(src []byte) (ParserResult, error) {

	log.Printf("Will parse person result")
	var requests []RequestInfo

	personInfo, err := ParsePerson(src)
	if err != nil {
		return ParserResult{}, err
	}
	var payload = []interface{}{personInfo}
	return ParserResult{payload, requests}, nil
}
