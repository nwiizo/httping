package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
	"flag"
)



func main() {
  flag.Usage = func() {
    fmt.Fprintf(os.Stderr, `
Usage of %s:
   %s [OPTIONS] ARGS...
Options\n`, os.Args[0],os.Args[0])
    flag.PrintDefaults()
  }
  var (
    d = flag.String("d", "default-value", "target_url")
    s = flag.String("s", "default-value", "dst_url")
  )
  flag.Parse()
	//コマンドライン引数として渡します
	fmt.Println(get_url(*d, *s))
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
