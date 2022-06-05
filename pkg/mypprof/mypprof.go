package mypprof

import (
	"BAT-douyin/setting"
	"fmt"
	"net/http"
)

func Init(conf *setting.PprofConfig) (err error) {
	err = http.ListenAndServe(fmt.Sprintf("%v:%v", conf.Host, conf.Port), nil)
	if err != nil {
		return err
	}
	return nil
}
