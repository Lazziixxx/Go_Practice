package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollar float32

//重新定义String()方法 定制类型的字符串形式的输出
//在fmt.Printf() 中会自动使用String()方法进行输出
//需要注意的是不要在(d dollar)String中涉及调用(d dollar)String方法的方法
func (d dollar) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollar

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}

func main() {
	db := database{"shoes":50, "socks":5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}
