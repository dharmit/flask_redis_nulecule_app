package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	"gopkg.in/redis.v3"
)

var rds, err = url.Parse(os.Getenv("REDIS_PORT"))

var client = redis.NewClient(&redis.Options{
	Addr:     rds.Host,
	Password: "",
	DB:       0,
})

var html_resp = "<div align=center>" + "<img src='https://pbs.twimg.com/profile_images/" +
	"458352291767013376/K9nN_rhH_400x400.png'>" +
	"<h1>This page has been visited %d times!</h1>" +
	"<br>" +
	"</div>"

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	val, err := client.Get("key2").Result()
	ival := 0
	if err == redis.Nil {
		client.Set("key2", "1", 0)
		ival = 1
	} else if err != nil {
		panic(err)
		fmt.Fprintf(w, "panic")
	} else {
		ival, err = strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		ival += 1
		client.Set("key2", ival, 0)
	}
	fmt.Fprintf(w, html_resp, ival)
}

func main() {
	fmt.Println(rds.Host)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", incrementCounter)
	http.ListenAndServe(":5000", router)
}
