package dingTalk

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type dingTalk struct {
	uri       string
	secretKey string
	token     string
	isSign    bool
	body      dingTalkBody
}

type dingTalkBody struct {
	MsgType    string                 `json:"msgtype"`
	Text       map[string]interface{} `json:"text"`
	Link       map[string]string      `json:"link"`
	MarkDown   map[string]string      `json:"markdown"`
	ActionCard map[string]string      `json:"actionCard"`
	At         map[string]interface{} `json:"at"`
}

type dingTallResp struct {
	ErrorCode int    `json:"errcode"`
	ErrorMsg  string `json:"errmsg"`
}

func mathSign(secretKey string) (string, string) {
	timeMillis := time.Now().UnixNano() / 1e6
	stringToSign := fmt.Sprintf("%d\n%s", timeMillis, secretKey)
	sign := computeHmacSha256(stringToSign, secretKey)

	return strconv.Itoa(int(timeMillis)), sign
}

func computeHmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func NewClient(uri, secretKey, token string, isSign bool) *dingTalk {
	client := &dingTalk{
		uri:       uri,
		secretKey: secretKey,
		token:     token,
		isSign:    isSign,
	}

	return client
}

func (d *dingTalk) send() error {
	if res, err := json.Marshal(d.body); err == nil {
		timeMillis, sign := mathSign(d.secretKey)
		v := url.Values{}
		v.Add("timestamp", timeMillis)
		v.Add("sign", sign)
		v.Add("access_token", d.token)
		reader := bytes.NewReader(res)
		request, err := http.NewRequest(http.MethodPost, d.uri, reader)
		if err != nil {
			return err
		}

		request.URL.RawQuery = v.Encode()
		request.Header.Set("Content-Type", "application/json;charset=UTF-8")

		client := http.Client{}
		resp, err := client.Do(request)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		respBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		var ret dingTallResp

		if err := json.Unmarshal(respBytes, &ret); err != nil {
			return fmt.Errorf("send dingTalk message failed, %v", err)
		}

		if ret.ErrorCode != 0 {
			return fmt.Errorf("send dingTalk message failed with error code is not 0, %v", ret)
		}

		return nil
	} else {
		return err
	}

}

func (d *dingTalk) SendTextMsg(data map[string]interface{}, at map[string]interface{}) error {
	d.body.MsgType = "text"
	d.body.Text = data
	d.body.At = at

	err := d.send()

	return err
}

func (d *dingTalk) SendLinkMsg(data map[string]string) error {
	d.body.MsgType = "link"
	d.body.Link = data

	err := d.send()

	return err
}

func (d *dingTalk) SendMarkDownMsg(data map[string]string, at map[string]interface{}) error {
	d.body.MsgType = "markdown"
	d.body.MarkDown = data
	d.body.At = at

	err := d.send()

	return err
}

func (d *dingTalk) SendActionCardMsg(data map[string]string, at map[string]interface{}) error {
	d.body.MsgType = "actionCard"
	d.body.ActionCard = data

	err := d.send()

	return err
}
