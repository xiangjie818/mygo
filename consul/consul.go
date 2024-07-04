package consul

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xiangjie818/mygo/logging"
	"github.com/xiangjie818/mygo/types"
	"io"
	"net/http"
)

type Service interface {
	GetKey(key string) []byte
	Register(svcRegistration types.ServiceRegistration)
	Deregister(serviceID string)
}

type service struct {
	server string
}

func (s *service) GetKey(key string) []byte {
	respUrl := s.server + "/v1/kv/" + key
	resp, err := http.Get(respUrl)
	if err != nil {
		logging.Error("地址：" + respUrl + "请求失败，错误信息为：" + err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logging.Error(err)
	}
	return body
}

func (s *service) Register(svcRegistration types.ServiceRegistration) {
	url := s.server + types.RegisterServicePath
	jsonData, err := json.Marshal(svcRegistration)
	if err != nil {
		logging.Error("Json解析失败: " + err.Error())
		return
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		logging.Error("请求失败: " + err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logging.Error("请求失败: " + err.Error())
		return
	}
	defer resp.Body.Close()
	logging.Info("请求状态码:", resp.Status)
}

func (s *service) Deregister(serviceID string) {
	url := fmt.Sprintf(s.server + types.DeregisterServicePath + "/" + serviceID)
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		logging.Error("请求失败: " + err.Error())
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logging.Error("请求失败: " + err.Error())
		return
	}
	defer resp.Body.Close()
	logging.Info("请求状态码:", resp.Status)
}

func New(server string) Service {
	return &service{
		server: server,
	}
}
