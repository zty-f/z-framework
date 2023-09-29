package main

import (
	"fmt"
	"log"
	"net/http"
	"zcache/zcache"
)

var db = map[string]string{
	"a": "1",
	"b": "2",
	"c": "3",
}

/*
curl "http://localhost:8888/_zcache/score/a"
1

curl "http://localhost:8888/_zcache/score/3"
key not found:3
*/

func main() {
	zcache.NewGroup("score", 2<<10, zcache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("get key from db", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("key not found:%s", key)
		}))
	addr := "localhost:8888"
	pool := zcache.NewHTTPPool(addr)
	log.Println("starting server at", addr)
	log.Fatal(http.ListenAndServe(addr, pool))
}
