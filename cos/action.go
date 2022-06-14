package cos

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"go.uber.org/zap"
	"mime/multipart"
)

func PutVideo2Cos(c *cos.Client, baseFileName string, file multipart.File) error {
	defer file.Close()
	_, err := c.Object.Put(context.Background(), fmt.Sprintf("video/%s.mp4", baseFileName), file, nil)
	if err != nil {
		zap.L().Error("put video to cos wrong")
		return err
	}
	return nil
}
