package parser

import "testing"

const PersonTestStr = `
<body>
	<h1 data-v-5b109fc3="" class="nickName">静曼</h1>
	<div data-v-5b109fc3="" class="id">ID：1394669573</div>
	<div data-v-5b109fc3="" class="des f-cl">
		阿拉善盟 | 33岁 | 大学本科 | 离异 | 162cm | 3001-5000元
	</div>
	<div class="m-content-box">
		<div class="purple-btns">
			<div class="m-btn purple">离异</div>
			<div class="m-btn purple">33岁</div>
			<div class="m-btn purple">魔羯座(12.22-01.19)</div>
			<div class="m-btn purple">162cm</div>
			<div class="m-btn purple">55kg</div>
			<div class="m-btn purple">工作地:阿拉善盟阿拉善左旗</div>
			<div class="m-btn purple">月收入:3-5千</div>
			<div class="m-btn purple">其他职业</div>
			<div class="m-btn purple">大学本科</div>
		</div>
	</div>
</body>
`
func TestParsePerson(t *testing.T) {
	expectPerson := PersonInfo{
		Name:          "静曼",
		Id:            "1394669573",
		Region:        "阿拉善盟",
		Age:           33,
		Edu:           "大学本科",
		MaritalStatus: "离异",
		Height:        162,
		Weight:        55,
		Salary:        [2]int{3001, 5000},
	}
	person, err := ParsePerson([]byte(PersonTestStr))
	if err != nil {
		t.Errorf("Parse error %s", err)
		return
	}
	if person != expectPerson {
		t.Errorf("Expect %v, but got %v", expectPerson, person)
	}
}