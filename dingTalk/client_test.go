package dingTalk

import "testing"

func TestDingTalk_SendTextMsg(t *testing.T) {
	client := NewClient(
		"https://oapi.dingtalk.com/robot/send",
		"xxxx",
		"xxx",
		true)

	actual := client.SendTextMsg(map[string]interface{}{"content": "test"}, map[string]interface{}{"atMobiles": []string{"xxxx"}})

	if actual != nil {
		t.Errorf("SendTextMsg = %v; expected %v", actual, nil)
	}

}
