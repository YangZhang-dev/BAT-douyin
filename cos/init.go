package cos

import (
	"BAT-douyin/setting"
	"errors"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
)

var CosClient *cos.Client

func Init(conf *setting.CosConfig) error {
	u, _ := url.Parse(conf.Url)
	b := &cos.BaseURL{BucketURL: u}
	CosClient = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  conf.Sid,
			SecretKey: conf.Skey,
		},
	})
	if CosClient == nil {
		return errors.New("connect cos failed")
	}
	return nil
}
