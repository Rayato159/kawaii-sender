<h1>Kawaii Sender</h1>
<p>This is just a stupid API shooter project for Golang.</p>

<h2>📝 Version</h2>
v0.1.0

<h2>🧙 Functions List</h2>

```go
FireHttpRequest(method HttpMethod, url string, headers map[string]string, body any, timeout time.Duration) ([]byte, error)
PrintJson(obj any)
PrintJsonWithTitle(title string, obj any)
WriteToJson(obj *Output)
```

<h2>⌨️ Quickstart</h2>

<h3>Installation</h3>

```bash
go get github.com/Rayato159/kawaii-sender
```
<h3>Usage</h3>

```go
package main

import (
	"fmt"
	"time"

	"github.com/Rayato159/kawaii-sender"
)

type testBody struct {
	Message string
}

func main() {
	// GET, POST, PUT, PATCH, DELETE
	methodGet := kawaiihttp.Get
	methodPost := kawaiihttp.Post

	// The url is expect for http protocol only
	url := "http://localhost:3000"
	timeout := time.Second * 120
	headers := map[string]string{
		"Authorization": "xxxxxxxxxxxxxxxxxxxx",
	}

	// Http Get
	g, err := kawaiihttp.FireHttpRequest(methodGet, url, headers, nil, timeout)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(g))

	// Http Post
	body := &testBody{
		Message: "Hello, Post!",
	}
	p, err := kawaiihttp.FireHttpRequest(methodPost, url, headers, body, timeout)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(p))
}
```