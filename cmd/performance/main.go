package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Worker struct {
	method      string
	url         string
	token       string
	contentType string
}

func rangeIn(low, hi int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return low + rand.Intn(hi-low)
}

func main() {
	now := time.Now()
	ch := make(chan int, 1)
	var wg sync.WaitGroup
	obj := Worker{
		method:      "POST",
		url:         "http://127.0.0.1:7173/api/restapi/v1/companies/1001/companies",
		token:       "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InN1cGVyIiwiaWQiOjExLCJsYW5ndWFnZSI6Imt1IiwiY29tcGFueV9pZCI6MTAwMSwibm9kZV9pZCI6MTAxLCJleHAiOjE2MTQ4MzE2Nzl9.4MBfh_JHH4ys4UXiRJyN_9Xi4GIWnlcNrNW7ve8cgrM",
		contentType: "application/JSON",
	}

	for i := 1; i <= 100; i++ {
		go obj.createMatCompany(i, ch, &wg)
	}

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		ch <- i
	}

	// time.Sleep(2 * time.Second)

	wg.Wait()
	fmt.Println(time.Since(now))
}

func (p *Worker) createMatCompany(n int, ch chan int, wg *sync.WaitGroup) {
	client := &http.Client{}
	var response *http.Response

	for v := range ch {
		time.Sleep(1 * time.Millisecond)
		// fmt.Println(n, v)

		randomvalue := strconv.Itoa(rangeIn(10000000, 99999999))

		payload := strings.NewReader(fmt.Sprintf(`{"name":"%v","address":"%v","notes":"%v"}`,
			randomvalue, n, v))

		req, err := http.NewRequest(p.method, p.url, payload)
		if err != nil {
			log.Fatalln(err)
			return
		}

		req.Header.Add("Authorization", p.token)
		req.Header.Add("Content-Type", p.contentType)

		response, err = client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer response.Body.Close()

		_, err = ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		// _ = v
		wg.Done()
	}

	fmt.Println("bye")
}
