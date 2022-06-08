package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
)

// GetCover 生成视频封面，位于static/cover下
func GetCover(videoPath string, baseFinalName string) {
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-vframes", "10", "-f", "singlejpeg", "-")
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	if cmd.Run() != nil {
		fmt.Println("could not generate frame")
	}
	jpg := fmt.Sprintf("./static/cover/%s.jpg", baseFinalName)
	err := ioutil.WriteFile(jpg, buf.Bytes(), 0666)
	if err != nil {
		fmt.Println(err)
	}
}
