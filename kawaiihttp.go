package kawaiihttp

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Rayato159/kawaii-sender/utils"
)

type HttpMethod string

const (
	Get    HttpMethod = "GET"
	Post   HttpMethod = "PUT"
	Put    HttpMethod = "PUT"
	Patch  HttpMethod = "PATCH"
	Delete HttpMethod = "DELETE"
)

func PrintJson(obj any) {
	objStr, _ := json.MarshalIndent(obj, "", "	")
	fmt.Println(string(objStr))
}

func PrintJsonWithTitle(title string, obj any) {
	objStr, _ := json.MarshalIndent(obj, "", "	")
	fmt.Println("", string(objStr))
}

func FireHttpRequest(method HttpMethod, url string, body any, timeout time.Duration) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Url validation
	if err := utils.UrlValidate(url); err != nil {
		return nil, err
	}

	// Body validation if that method is not allowed
	if err := utils.BodyValidate(string(method), body); err != nil {
		return nil, err
	}

	// Config before request
	config := http.Header{
		"Content-Type": {"application/json"},
	}
	req, err := http.NewRequestWithContext(ctx, string(method), url, nil)
	if err != nil {
		return nil, fmt.Errorf("http request config error: %v", err)
	}
	req.Header = config
	client := new(http.Client)

	// Request fire!!!
	res, err := client.Do(req)
	if err != nil {
		defer func() {
			req.Close = true
			res.Body.Close()
		}()
		return nil, fmt.Errorf("http request error: %v", err)
	}
	defer func() {
		req.Close = true
		res.Body.Close()
	}()

	// Body response
	resJson, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("response error: %v", err)
	}
	return resJson, nil
}
