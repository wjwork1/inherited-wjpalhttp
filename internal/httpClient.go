package ihttp

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"strconv"
)

type CHttpClient struct {
	reqUrl         string
	reqBody        []byte
	mapHeader      map[string]string
	lastStatusCode int
}

func NewHttpClient() *CHttpClient {
	return &CHttpClient{mapHeader: make(map[string]string), reqBody: nil}
}

func (pInst *CHttpClient) Initialize() error {
	return nil
}
func (pInst *CHttpClient) PutUrl(url string) {
	pInst.reqUrl = url
}
func (pInst *CHttpClient) AppendHeader(headKey, headValue string) {
	pInst.mapHeader[headKey] = headValue
}
func (pInst *CHttpClient) PutBody(body []byte) {
	pInst.reqBody = body
}
func (pInst *CHttpClient) Clean() {
	pInst.reqUrl = ""
	pInst.reqBody = nil
	pInst.mapHeader = make(map[string]string)
	pInst.lastStatusCode = 0
}
func (pInst *CHttpClient) GetLastStatusCode() int {
	return pInst.lastStatusCode
}

func (pInst *CHttpClient) Get() ([]byte, error) {
	return pInst.Request("GET")
}
func (pInst *CHttpClient) Post() ([]byte, error) {
	return pInst.Request("POST")
}
func (pInst *CHttpClient) Put() ([]byte, error) {
	return pInst.Request("PUT")
}
func (pInst *CHttpClient) Request(requestMethod string) ([]byte, error) {
	pInst.lastStatusCode = -100
	var bodyBuf *bytes.Buffer = nil
	if pInst.reqBody != nil {
		bodyBuf = bytes.NewBuffer(pInst.reqBody)
	}

	req1, err := http.NewRequest(requestMethod, pInst.reqUrl, bodyBuf)
	if err != nil {
		pInst.lastStatusCode = -101
		return nil, err
	}
	for k, v := range pInst.mapHeader {
		req1.Header.Add(k, v)
	}

	client := &http.Client{}
	resp1, err := client.Do(req1)
	if err != nil {
		pInst.lastStatusCode = -102
		return nil, err
	}

	pInst.lastStatusCode = resp1.StatusCode
	if resp1.StatusCode != http.StatusOK {
		return nil, errors.New(resp1.Status + " " + strconv.Itoa(resp1.StatusCode))
	}

	body1, err := io.ReadAll(resp1.Body)
	if err != nil {
		return nil, err
	}

	return body1, nil
} /*
func (pInst *CHttpClient) Get() ([]byte, error) {
	req1, err := http.NewRequest("GET", pInst.reqUrl, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range pInst.mapHeader {
		req1.Header.Add(k, v)
	}

	client := &http.Client{}
	resp1, err := client.Do(req1)
	if err != nil {
		return nil, err
	}

	if resp1.StatusCode != http.StatusOK {
		return nil, errors.New(resp1.Status + " " + strconv.Itoa(resp1.StatusCode))
	}

	body1, err := io.ReadAll(resp1.Body)
	if err != nil {
		return nil, err
	}

	return body1, nil
}

func (pInst *CHttpClient) Post() ([]byte, error) {
	var bodyBuf *bytes.Buffer = nil
	if pInst.reqBody != nil {
		bodyBuf = bytes.NewBuffer(pInst.reqBody)
	}

	req1, err := http.NewRequest("POST", pInst.reqUrl, bodyBuf)
	if err != nil {
		return nil, err
	}
	for k, v := range pInst.mapHeader {
		req1.Header.Add(k, v)
	}

	client := &http.Client{}
	resp1, err := client.Do(req1)
	if err != nil {
		return nil, err
	}

	if resp1.StatusCode != http.StatusOK {
		return nil, errors.New(resp1.Status + " " + strconv.Itoa(resp1.StatusCode))
	}

	body1, err := io.ReadAll(resp1.Body)
	if err != nil {
		return nil, err
	}

	return body1, nil
}*/
