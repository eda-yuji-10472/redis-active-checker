// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START memorystore_main_go]

// Command redis is a basic app that connects to a managed Redis instance.
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

func main() {
	ex := os.Getenv("EX")
	hostName := os.Getenv("HOSTNAME")
	redisHost := os.Getenv("REDISHOST")
	redisPort := os.Getenv("REDISPORT")
	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)

	var redisPool *redis.Pool

	const maxConnections = 2
	redisPool = &redis.Pool{
		MaxIdle: maxConnections,
		Dial:    func() (redis.Conn, error) { return redis.Dial("tcp", redisAddr) },
	}

	conn := redisPool.Get()
	defer conn.Close()

	i, err := strconv.Atoi(ex)
	if err != nil {
		// ... handle error
		panic(err)
	}
	// read hostname
	s, err := redis.String(conn.Do("GET", "hostName"))
	if err != nil || s == hostName {
		// write hostname
		r, err := conn.Do("SET", "hostName", hostName, "EX", ex)
		if err != nil {
			fmt.Print(err)
			return
		}
		fmt.Println(s)
		fmt.Println(r) // OK
		time.Sleep(time.Second * 1 * time.Duration(i))
		return
	} else {
		fmt.Println("Active Node: ", s)
		time.Sleep(time.Second * 1 * time.Duration(i))
		//os.Exit(1)
		return
	}

}

// [END memorystore_main_go]
