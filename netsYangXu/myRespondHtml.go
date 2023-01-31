package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8888", nil)
}

/*
上面的handlerFunc实际上是一个适配器，将func转换为handler
handler是一个接口，实现了handler的ServeHTTP方法，就能够为特定路径提供服务
上面传入的nil实际上是指定了默认的ServeMux，可以自己实现但是没有太大必要*/

func hello(w http.ResponseWriter, r *http.Request) {
	str := `<html>
	<head><title>go web</title></head>
	<body><h1>hello!!! world</h1></body>
	</html>`
	w.Write([]byte(str))
}
