package prometheus

import (
	"encoding/json"
	"fmt"
	"github.com/xiangjie818/mygo/logging"
	"github.com/xiangjie818/mygo/types"
	"gopkg.in/ini.v1"
	"io"
	"net/http"
	"net/url"
)

type Service interface {
	Query(queryName string) (types.QueryResponse, error)
}

type service struct {
	conf *ini.File
}

func New(conf *ini.File) Service {
	return &service{
		conf: conf,
	}
}

func (s *service) Query(queryName string) (resp types.QueryResponse, err error) {
	promServer := s.conf.Section("prometheus").Key("server").String()
	query := s.conf.Section("query").Key(queryName).String()
	v := url.Values{}
	v.Set("query", query)
	getUrl := fmt.Sprintf("%s/api/v1/query?%s", promServer, v.Encode())
	response, err := http.Get(getUrl)
	if err != nil {
		logging.Error("Error while querying prometheus: ", err)
		return
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		logging.Error("Error while reading response body: ", err)
		return
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		logging.Error("Error while unmarshaling response body: ", err)
		return
	}
	return
}
