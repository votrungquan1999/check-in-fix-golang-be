package utils

import (
	"bytes"
	"checkinfix.com/setup"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type sendSMSBodyWithNexmo struct {
	From      string `json:"from"`
	Text      string `json:"text"`
	To        string `json:"to"`
	ApiKey    string `json:"api_key"`
	ApiSecret string `json:"api_secret"`
	Type      string `json:"type"`
}

func SendSMSWithNexmo(from string, to string, text string) (interface{}, error) {
	var url = "https://rest.nexmo.com/sms/json"
	body := sendSMSBodyWithNexmo{
		From:      from,
		Text:      text,
		To:        to,
		ApiKey:    setup.EnvConfig.NexmoApiKey,
		ApiSecret: setup.EnvConfig.NexmoApiSecret,
		Type:      "unicode",
	}

	payloadBuf := new(bytes.Buffer)

	err := json.NewEncoder(payloadBuf).Encode(body)

	fmt.Println(payloadBuf.String())
	if err != nil {
		return nil, ErrorInternal.New(err.Error())
	}
	//
	//res, err := http.Post(url, "application/x-www-form-urlencoded", payloadBuf)
	res, err := http.Post(url, "application/json", payloadBuf)
	if err != nil {
		return nil, ErrorInternal.New(err.Error())
	}

	//client := &http.Client{}
	//r, err := http.NewRequest("POST", url, payloadBuf) // URL-encoded payload
	//if err != nil {
	//	log.Fatal(err)
	//}
	//r.Header.Add("Content-Type", "application/json")
	//r.Header.Add("Content-Length", strconv.Itoa(payloadBuf.Len()))

	//res, err := client.Do(r)
	//if err != nil {
	//	log.Fatal(err)
	//}

	defer func() {
		_ = res.Body.Close()
	}()

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, ErrorInternal.New(err.Error())
	}

	//abc := bytes.Fields(respBody)

	//var respBody []byte

	//_, _ = res.Body.Read(respBody)
	return string(respBody), nil
}

func SendSMSWithTwilio(from string, to string, text string) (interface{}, error) {
	accountSid := "AC8a42c07bb553d72182d57dd5733707fa"
	authToken := "3ac051dd261cbdd4feb1f33f27a3a9b5"
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	v := url.Values{}
	v.Set("To", to)
	v.Set("From", "+12062740788")
	v.Set("Body", text)
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}

	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(req)
	fmt.Println(resp.Status)

	defer func() {
		_ = resp.Body.Close()
	}()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, ErrorInternal.New(err.Error())
	}

	return string(respBody), nil
}
