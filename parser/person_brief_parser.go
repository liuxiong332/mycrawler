package parser

import "regexp"

const BriefReStr = `<a href="(http://album.zhenai.com/u/\w+)" [^>]+>([^<]+)</a>`
var BriefRe = regexp.MustCompile(BriefReStr)

type PersonBriefItem struct {
	Url string
	Name string
}

func ParsePersonBrief(src []byte) []PersonBriefItem {
	findRes := BriefRe.FindAllSubmatch(src, -1)

	var items []PersonBriefItem
	for _, m := range findRes {
		item := PersonBriefItem{Url: string(m[1]), Name:string(m[2])}
		items = append(items, item)
	}
	return items
}