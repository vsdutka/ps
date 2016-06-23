// sms
package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"gopkg.in/errgo.v1"
)

const (
	g_username = "hazard"
	g_password = "301286"
)

func SmsSend(phone, message string) error {

	form := url.Values{}
	form.Add("Http_username", g_username)
	form.Add("Http_password", g_password)
	form.Add("Phone_list", phone)
	form.Add("Message", message)

	hc := http.Client{}
	req, err := http.NewRequest("GET", "http://www.websms.ru/http_in5.asp?"+form.Encode(), nil)

	req.Header.Add("Content-Type", "text/plain; charset=utf-8")
	resp, err := hc.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errgo.Newf("Ошибка отправки sms: %s", resp.Status)
	}

	var b []byte
	if b, err = ioutil.ReadAll(resp.Body); err != nil {
		return err
	}
	//FIXME Добавить парсинг ответа и сигнализацию об ошибке
	fmt.Println(string(b))
	return nil
}
