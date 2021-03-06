package main

import "github.com/mingz2013/login-server-go/server"
import "encoding/json"

func main() {

	confMap := map[string]interface{}{

		"host":    "redis-mq",
		"port":    "6379",
		"db":      1,
		"channel": "connector-server",
	}

	data, _ := json.Marshal(confMap)

	a := server.NewApp(data)
	a.Start()
}
