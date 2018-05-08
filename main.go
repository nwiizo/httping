package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	//コマンドライン引数として渡します
	fmt.Println(get_url(os.Args[1], os.Args[2]))
}

func get_url(target_url string, dst_url string) string {
	var RedirectAttemptedError = errors.New("redirect")
	client := &http.Client{
		Timeout: time.Duration(3) * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return RedirectAttemptedError
		},
	}
	resp, err := client.Head(target_url)
	defer resp.Body.Close()
	if urlError, ok := err.(*url.Error); ok && urlError.Err == RedirectAttemptedError {
		src_url := (strings.Join(resp.Header["Location"], " "))
		if src_url == dst_url {
			return "OK"
		} else {
			defer fmt.Println("Ideal:", dst_url, "Real:", src_url)
			return "NG"
		}
	}
	return "NOT"
}
