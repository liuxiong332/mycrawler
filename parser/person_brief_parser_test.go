package parser

import "testing"

const TestHtml = `
<th><a href="http://album.zhenai.com/u/1715513586" target="_blank">白酒拌馍馍</a></th>
`
func TestParsePersonBrief(t *testing.T) {
	ExpectRes := []struct{
		Url string
		Name string
	}{
		{"http://album.zhenai.com/u/1715513586", "白酒拌馍馍"},
	}

	items := ParsePersonBrief([]byte(TestHtml))
	for i, m := range items {
		if m.Url != ExpectRes[i].Url || m.Name != ExpectRes[i].Name {
			t.Errorf("Expect %s, %s, but got %s, %s", ExpectRes[i].Url, ExpectRes[i].Name, m.Url, m.Name)
		}
	}
}