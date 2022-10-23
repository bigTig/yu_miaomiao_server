package utils

import (
	"bytes"
	"go.uber.org/zap"
	"io"
	"net/http"
	"time"
	"yuyu/global"
)

func Get(url string) (int, []byte) {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		global.GvaLog.Error("error sending GET request", zap.Error(err))
		return http.StatusInternalServerError, nil
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			global.GvaLog.Error("error decoding response from GET request", zap.Error(err))
		}
	}
	return resp.StatusCode, result.Bytes()
}
