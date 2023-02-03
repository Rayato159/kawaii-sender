<h1>Kawaii Sender</h1>
<p>This is just a stupid API shooter project for Golang.</p>

<h2>📝 Version</h2>
v0.1.0

<h2>🧙 Functions List</h2>

```go
FireHttpRequest(method HttpMethod, url string, body any, timeout time.Duration) ([]byte, error)
PrintJson(obj any)
PrintJsonWithTitle(title string, obj any)
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
	method := kawaiihttp.Get

	// The url is expect for http protocol only
	url := "http://localhost:3000"
	timeout := time.Second * 120

	// Http Get
	g, err := kawaiihttp.FireHttpRequest(method, url, nil, timeout)
	if err != nil {
		panic(err)
	}
	fmt.Println(g)

	// Http Post
	body := &testBody{
		Message: "Hello, Post!",
	}
	p, err := kawaiihttp.FireHttpRequest(method, url, body, timeout)
	if err != nil {
		panic(err)
	}
	fmt.Println(p)
}
```