package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xiangjie818/mygo/logging"
	"github.com/xiangjie818/mygo/prometheus"
	"gopkg.in/ini.v1"
	"log"
	"net/http"
	"strconv"
)

func alertWX(volPath, usage, webhook string) {
	url := "http://sendweixin.pso-monitor.paas.corp/send_weixin"
	valueFloat, _ := strconv.ParseFloat(usage, 64)
	var threshold float64 = 80
	if valueFloat > threshold {
		valueStr := strconv.FormatFloat(valueFloat, 'f', 2, 64)
		alertTitle := "# OSD使用率告警\n"
		alertPath := "> **OSD：**" + volPath + "\n"
		alertValue := "> **使用率：**" + valueStr + "%\n"

		alertData := map[string]interface{}{
			"webhook": webhook,
			"msgtype": "markdown",
			"content": alertTitle + alertPath + alertValue,
		}
		jsonData, _ := json.Marshal(alertData)
		fmt.Println(bytes.NewBuffer(jsonData))
		_, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println(err)
		}
	} else {
		log.Printf("Threshold: %f is greater than usage: %f", threshold, valueFloat)
	}
}

func main() {
	//webhook := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=accc5ea0-922d-4176-9317-83f0e1599fd1"
	logging.InitLoggers("/tmp/mygo.log")
	cfg, err := ini.Load("conf.ini")
	if err != nil {
		logging.Error("load config file error")
		return
	}
	prom := prometheus.New(cfg)
	data, err := prom.Query("osd_disk_usage")
	if err != nil {
		logging.Error("query prometheus error")
		return
	}
	for _, v := range data.Data.Result {
		usage := v.Value[1].(string)
		usageFloat, _ := strconv.ParseFloat(usage, 64)
		if usageFloat > 80 {
			logging.Error("usage is greater than 80. OSD: ", v.Metric["ceph_daemon"].(string))
		}
	}
}
