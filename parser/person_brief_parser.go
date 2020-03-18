package parser

import (
	"log"
	"regexp"
)

const BriefReStr = `<a href="(http://album.zhenai.com/u/\w+)" [^>]+>([^<]+)</a>`

var BriefRe = regexp.MustCompile(BriefReStr)

type PersonBriefItem struct {
	Url  string
	Name string
}

func ParsePersonBrief(src []byte) []PersonBriefItem {
	findRes := BriefRe.FindAllSubmatch(src, -1)

	var items []PersonBriefItem
	for _, m := range findRes {
		item := PersonBriefItem{Url: string(m[1]), Name: string(m[2])}
		items = append(items, item)
	}
	return items
}

func NilParser(src []byte) (ParserResult, error) {
	log.Printf("Nil parser")
	return ParserResult{nil, nil}, nil
}

func ParsePersonBriefRes(src []byte) (ParserResult, error) {
	items := ParsePersonBrief(src)
	var payload []interface{}
	var requests []RequestInfo
	for _, m := range items {
		payload = append(payload, m)
	}

	briefItems := ParsePersonBrief(src)
	for _, item := range briefItems {
		requests = append(requests, RequestInfo{
			Url:    item.Url,
			Parser: ParsePersonRes,
		})
	}
	return ParserResult{payload, requests}, nil
}
