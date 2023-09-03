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

var conn redis.Conn

const lua string = `
-- KEYS[1] target key
-- KEYS[2] max count key from redis config
-- ARGV[n = 3] max count, duration(milliseconds), current timestamp

-- HASH: KEYS[1]
--   field:count
--   field:limit
--   field:duration
--   field:reset

local res = {}
local limit = redis.call('hmget', KEYS[1], 'count', 'limit', 'duration', 'reset')

if limit[1] then
  res[1] = tonumber(limit[1]) - 1
  res[2] = tonumber(limit[2])
  res[3] = tonumber(limit[3]) or ARGV[3]
  res[4] = tonumber(limit[4])
  if res[1] >= -1 then
    redis.call('hincrby', KEYS[1], 'count', -1)
  else
    res[1] = -1
  end
else
  local total = tonumber(redis.call('get', KEYS[2]) or ARGV[2])
  res[1] = total - 1
  res[2] = total
  res[3] = tonumber(ARGV[3])
  res[4] = tonumber(ARGV[1]) + res[3]
  redis.call('hmset', KEYS[1], 'count', res[1], 'limit', res[2], 'duration', res[3], 'reset', res[4])
  redis.call('pexpire', KEYS[1], res[3])
end

return res
`

func init() {
	var err error
	conn, err = redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}
}


func main() {
	defer conn.Close()
	total := "5"
	s := redis.NewScript(2, lua)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			now := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
			duration := strconv.FormatInt(10*time.Second.Milliseconds(), 10)
			log.Println(total, now, duration)
			v, err := redis.Ints(s.Do(conn, "user-1001", "total-user", now, total, duration))
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
