package util

import (
	"crawler/parser"
	"reflect"
)

func ConvertReqType(reqType string) parser.RequestParser {
	var reqParser parser.RequestParser
	switch reqType {
	case "Person":
		reqParser = parser.ParsePersonRes
	case "PersonBrief":
		reqParser = parser.ParsePersonBriefRes
	case "Region":
		reqParser = parser.ParseRegionRes
	}
	return reqParser
}

func ConvertReqParser(reqParser parser.RequestParser) string {
	switch reflect.ValueOf(reqParser).Pointer() {
	case reflect.ValueOf(parser.ParsePersonRes).Pointer():
		return "Person"
	case reflect.ValueOf(parser.ParsePersonBriefRes).Pointer():
		return "PersonBrief"
	case reflect.ValueOf(parser.ParseRegionRes).Pointer():
		return "Region"
	}
	return ""
}
