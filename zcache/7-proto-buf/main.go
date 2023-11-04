package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"zcache/zcache"
)

var db = map[string]string{
	"Tom":  "你好",
	"Jack": "589",
	"Sam":  "567",
}

func createGroup() *zcache.Group {
	return zcache.NewGroup("scores", 2<<10, zcache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))
}

func startCacheServer(addr string, addrs []string, z *zcache.Group) {
	peers := zcache.NewHTTPPool(addr)
	peers.Set(addrs...)
	z.RegisterPeers(peers)
	log.Println("zcache is running at", addr)
	log.Fatal(http.ListenAndServe(addr[7:], peers))
}

func startAPIServer(apiAddr string, z *zcache.Group) {
	http.Handle("/api", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			key := r.URL.Query().Get("key")
			view, err := z.Get(key)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/octet-stream")
			w.Write(view.ByteSlice())

		}))
	log.Println("fontend server is running at", apiAddr)
	log.Fatal(http.ListenAndServe(apiAddr[7:], nil))

}

/*
2023/11/04 17:29:43 zcache is running at http://localhost:8002
2023/11/04 17:29:43 zcache is running at http://localhost:8001
2023/11/04 17:29:43 zcache is running at http://localhost:8003
2023/11/04 17:29:43 fontend server is running at http://localhost:9999
>>> start test
2023/11/04 17:29:45 [Server http://localhost:8003] pick peer: http://localhost:8001
2023/11/04 17:29:45 [Server http://localhost:8001] GET /_zcache/scores/Tom
2023/11/04 17:29:45 [SlowDB] search key Tom
你好你好你好你好你好你好你好

*/

func main() {
	var port int
	var api bool
	flag.IntVar(&port, "port", 8001, "zcache server port")
	flag.BoolVar(&api, "api", false, "Start a api server?")
	flag.Parse()

	apiAddr := "http://localhost:9999"
	addrMap := map[int]string{
		8001: "http://localhost:8001",
		8002: "http://localhost:8002",
		8003: "http://localhost:8003",
	}

	var addrs []string
	for _, v := range addrMap {
		addrs = append(addrs, v)
	}

	z := createGroup()
	if api {
		go startAPIServer(apiAddr, z)
	}
	startCacheServer(addrMap[port], []string(addrs), z)
}
