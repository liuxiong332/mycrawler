package parser

import (
	"regexp"
)

const RegionReStr = `<a [^>]+ href="(http://www.zhenai.com/zhenghun/\w+)">([^<]+)</a>`

var RegionRe = regexp.MustCompile(RegionReStr)

type RegionItem struct {
	Url  string
	Name string
}

func ParseRegion(src []byte) []RegionItem {
	matchRes := RegionRe.FindAllSubmatch(src, -1)
	var regionItems []RegionItem
	for _, m := range matchRes {
		item := RegionItem{Url: string(m[1]), Name: string(m[2])}
		regionItems = append(regionItems, item)
	}
	return regionItems
}

func ParseRegionRes(src []byte) (ParserResult, error) {
	items := ParseRegion(src)

	var payload = []interface{}{items[0]}
	//var payload []interface{} = make([]interface{}, len(items))
	//for i, item := range items {
	//	payload[i] = item
	//}
	var requests []RequestInfo

	for _, item := range items {
		requests = append(requests, RequestInfo{
			Url:    item.Url,
			Parser: ParsePersonBriefRes,
		})
	}
	return ParserResult{Payload: payload, Requests: requests}, nil
}
