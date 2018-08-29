package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis"
)

const port = "3003"

var redisClient *redis.Client

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	redisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatal("Unable to connect to redis")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		url := params.Get("url")

		if url != "" {
			handleSet(url, w, r)
		} else {
			handleGet(r.URL.Path[1:], w, r)
		}
	})

	log.Println("Listening on " + port)

	panic(http.ListenAndServe(":"+port, nil))
}

func handleGet(key string, w http.ResponseWriter, r *http.Request) {
	if key == "" {
		fmt.Fprintln(w, "Missing key, expected "+r.Host+"/key")
		return
	}

	value, err := redisClient.Get(key).Result()
	if err != nil {
		log.Println("Error getting data", key)
		http.Error(w, "Error getting data", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, value, http.StatusSeeOther)
}

func handleSet(url string, w http.ResponseWriter, r *http.Request) {
	key := randomKey()

	if !validURL(url) {
		fmt.Fprintln(w, "URL too long, must be between 4 and 8000 characters")
		return
	}

	value := url
	if url[:4] != "http" {
		value = "http://" + url
	}

	err := redisClient.Set(key, value, 0).Err()
	if err != nil {
		log.Println("Error setting data", key, value)
		http.Error(w, "Error setting data", http.StatusInternalServerError)
		return
	}

	response := "<a href=\"https://" + r.Host + "/" + key + "\">" + r.Host + "/" + key + "</a>"
	fmt.Fprintln(w, response)
}

func validURL(url string) bool {
	return len(url) <= 8000 && len(url) >= 4
}

func randomKey() string {
	letter := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, 3)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
