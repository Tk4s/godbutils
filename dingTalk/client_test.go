package dingTalk

import "testing"

func TestDingTalk_SendTextMsg(t *testing.T) {
	client := NewClient(
		"https://oapi.dingtalk.com/robot/send",
		"xxx",
		"xxxx",
		true)

	a := `
### 哈哈哈哈
`

	actual := client.SendMarkDownMsg(map[string]string{"content": a, "title": "test"}, map[string]interface{}{"atMobiles": []string{"xxxx"}})

	if actual != nil {
		t.Errorf("SendTextMsg = %v; expected %v", actual, nil)
	}

}
