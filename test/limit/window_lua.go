package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
)

const limitWindow  =`
local res = {}
local key = KEYS[1]
local score = tonumber(ARGV[1])
local limit = tonumber(redis.call('get', KEYS[2]) or ARGV[2])
local duration = tonumber(ARGV[3])
local max = tonumber(score - duration*1000000)
redis.call("zremrangebyscore", key, 0, max)
local count = redis.call("zcard", key)
if count >= limit then
	res[1] = -1
else
	res[1] = limit - count - 1
	redis.call("zadd", key, score, score)
	redis.call("pexpire", key, duration)
end
res[2] = limit
res[3] = duration
res[4] = duration*100000 + score
return res
`

var connWindow redis.Conn

func init() {
	var err error
	connWindow, err = redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}
}


func main() {
	defer connWindow.Close()
	total := "5"
	s := redis.NewScript(2, limitWindow)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			now := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
			duration := strconv.FormatInt(10*time.Second.Milliseconds(), 10)
			log.Println(total, now, duration)
			v, err := redis.Ints(s.Do(connWindow, "user-1001", "total-user", now, total, duration))
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			if v[0] >= 0 {
				w.WriteHeader(200)
				fmt.Fprintf(w, "Path: %q\n", html.EscapeString(r.URL.Path))
				fmt.Fprintf(w, "Remaining: %d\n", v[0])
				fmt.Fprintf(w, "Total: %d\n", v[1])
				fmt.Fprintf(w, "Duration: %v\n", v[2])
				fmt.Fprintf(w, "Reset: %v\n", v[3])
			} else {
				w.WriteHeader(429)
				fmt.Fprintf(w, "Rate limit exceeded, retry seconds.\n")
			}
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}