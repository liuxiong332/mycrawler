package parser

type RequestParser func([]byte) (ParserResult, error)

type RequestInfo struct {
	Url    string
	Parser func(src []byte) (ParserResult, error)
}

type ParserResult struct {
	Payload  []interface{}
	Requests []RequestInfo
}
