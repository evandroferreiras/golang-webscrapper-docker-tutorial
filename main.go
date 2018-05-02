package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//Result exported
type Result struct {
	url      string
	userName string
	title    string
	likes    string
}

func (r Result) String() string {
	return fmt.Sprint(r.userName, " - ", r.title, " - ", r.likes, " claps")
}

func handlerLastProcess(w http.ResponseWriter, r *http.Request) {
	total, err := GetTotalURLs()
	if err != nil {
		fmt.Fprintln(w, "There is no data")
	}
	for index := 0; index < total; index++ {
		url, err := GetURL(index)
		if err != nil {
			fmt.Fprintln(w, "Error when acessing redis:", err)
			return
		}

		fmt.Fprintln(w, "Url:", url, "Claps:", GetURLLikes(url))
	}
}

func handlerProcessUrls(w http.ResponseWriter, r *http.Request) {

	urlToProcess := []string{
		"https://medium.freecodecamp.org/how-to-columnize-your-code-to-improve-readability-f1364e2e77ba",
		"https://medium.freecodecamp.org/how-to-think-like-a-programmer-lessons-in-problem-solving-d1d8bf1de7d2",
		"https://medium.freecodecamp.org/code-comments-the-good-the-bad-and-the-ugly-be9cc65fbf83",
		"https://uxdesign.cc/learning-to-code-or-sort-of-will-make-you-a-better-product-designer-e76165bdfc2d",
	}
	ini := time.Now()
	cRes := make(chan Result)
	go scrapListURL(urlToProcess, cRes)
	fmt.Fprintln(w, "With goroutines:")

	index := 0
	for res := range cRes {
		err := SetURL(index, res.url)
		if err != nil {
			fmt.Fprintln(w, "Error when acessing redis:", err)
			return
		}
		fmt.Fprintln(w, res)
		index = index + 1
		SetURLLikes(res.url, res.likes)
		SetTotalURLs(index)
	}
	fmt.Fprintln(w, "(Took ", time.Since(ini).Seconds(), "secs)")
}

func handlerMainProcess(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Nothing here. Use /process-urls/ and /last-process/ ")
}

func main() {
	http.HandleFunc("/", handlerMainProcess)
	http.HandleFunc("/process-urls/", handlerProcessUrls)
	http.HandleFunc("/last-process/", handlerLastProcess)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
