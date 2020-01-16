package httphelper

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"taking_you/app/logging"
)

func Get(path string, para map[string]string) (data []byte, err error) {

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil{
		return
	}
	params := url.Values{}
	for k, v := range para{
		params.Add(k, v)
	}
	req.URL.RawQuery = params.Encode()
	logging.Info("Info:", req.URL.RawQuery)
	resp, err := client.Do(req)
	if err != nil{
		return
	}
	defer resp.Body.Close()
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil{
		return
	}
	return
}

func PostWithJSONForm(path string, header map[string]string, para map[string]interface{}) (data []byte, err error) {
	var bodyData []byte
	client := &http.Client{}
	bodyData, err = json.Marshal(para)
	if err != nil{
		return
	}
	req, err := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(bodyData))
	if err != nil{
		return
	}
	if header != nil{
		for k, v := range header{
			req.Header.Set(k, v)
		}
	}
	resp, err := client.Do(req)
	if err != nil{
		return
	}
	defer resp.Body.Close()
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil{
		return
	}
	return
}

func PostWithStringForm(path string, header map[string]string, paraBody string) (data []byte, err error) {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, path, bytes.NewBuffer([]byte(paraBody)))
	if err != nil{
		return
	}
	if header != nil{
		for k, v := range header{
			req.Header.Set(k, v)
		}
	}
	resp, err := client.Do(req)
	if err != nil{
		return
	}
	if resp.StatusCode != http.StatusOK{
		err = errors.New(fmt.Sprintf("Invalid status code:%d", resp.StatusCode))
		return
	}
	defer resp.Body.Close()
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil{
		return
	}
	return
}

func PostWithRawData(path string, header map[string]string, sendData interface{})(retData []byte, err error)  {
	var(
		data []byte
	)
	data, err = json.Marshal(sendData)
	if err != nil{
		return
	}
	fmt.Println(string(data))
	req, err := http.NewRequest(http.MethodPost, path, bytes.NewBuffer([]byte(data)))
	if err != nil{
		return
	}
	if header != nil{
		for k, v := range header{
			req.Header.Set(k, v)
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		return
	}
	if resp.StatusCode != http.StatusOK{
		err = errors.New(fmt.Sprintf("Invalid status code:%d", resp.StatusCode))
		return
	}
	defer resp.Body.Close()
	retData, err = ioutil.ReadAll(resp.Body)

	return
}


