package main

import (
	"strconv"

	"github.com/go-redis/redis"
)

func getClient() (r *redis.Client) {
	r = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return
}

//SetURL : Set url
func SetURL(index int, url string) {
	client := getClient()
	err := client.Set("url"+string(index), url, 0).Err()
	if err != nil {
		panic(err)
	}
}

//GetURL : Set url
func GetURL(index int) (url string) {
	client := getClient()
	v, err := client.Get("url" + string(index)).Result()
	if err != nil {
		panic(err)
	}
	url = v
	return
}

//SetTotalURLs : Set total urls to redis
func SetTotalURLs(total int) {
	client := getClient()
	err := client.Set("TotalUrls", total, 0).Err()
	if err != nil {
		panic(err)
	}
}

// GetTotalURLs : Get total urls from redis
func GetTotalURLs() (total int) {
	client := getClient()
	v, err := client.Get("TotalUrls").Result()
	if err != nil {
		panic(err)
	}
	total, err = strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return
}

//SetURLLikes : Set URL to redis
func SetURLLikes(url string, claps string) {
	client := getClient()
	err := client.Set(url, claps, 0).Err()
	if err != nil {
		panic(err)
	}

}

//GetURLLikes : Get URL from Redis
func GetURLLikes(url string) (r string) {
	client := getClient()
	v, err := client.Get(url).Result()
	if err != nil {
		panic(err)
	}
	r = v
	return
}
