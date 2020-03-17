package parser

import "testing"

const TestStr = `
<dd data-v-5e16505f="">
	<a data-v-5e16505f="" href="http://www.zhenai.com/zhenghun/aba">阿坝</a>
	<a data-v-5e16505f="" href="http://www.zhenai.com/zhenghun/akesu">阿克苏</a>
</dd>`


func TestRegionParser(t *testing.T) {
	var expectRes = []struct {
		Url string
		Name string
	}{{"http://www.zhenai.com/zhenghun/aba", "阿坝"}, {"http://www.zhenai.com/zhenghun/akesu", "阿克苏"}}

	items := parseRegion([]byte(TestStr))

	for i, res := range items {
		if res.Name != expectRes[i].Name || res.Url != expectRes[i].Url {
			t.Errorf("Expectd %s, %s, but got %s %s", expectRes[i].Url, expectRes[i].Name, res.Url, res.Name)
		}
	}
}
