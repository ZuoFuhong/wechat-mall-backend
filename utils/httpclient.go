package utils

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

func HttpGet(url string) (string, error) {
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var buf [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	return result.String(), nil
}

func HttpPost() {

}
