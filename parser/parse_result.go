package parser

type RequestInfo struct {
	Url string
	Parser func(src []byte) (ParserResult, error)
}

type ParserResult struct {
	Payload []interface{}
	Requests []RequestInfo
}
