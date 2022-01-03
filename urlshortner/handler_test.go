package urlshortner

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/bmizerany/assert"
)

type Input struct {
	Url string `json:"url"`
}

type JsonResp struct {
	// Reserved field to add some meta information to the API response
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func TestEncode(t *testing.T) {
	// go startServer()
	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	inputReq := Input{Url: "https://play.golang.com/"}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(inputReq)
	if err != nil {
		log.Fatal(err)
	}

	r, _ := http.NewRequest("POST", "http://localhost:8082/urlEncoder", &buf)

	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	var jsonResp JsonResp
	json.NewDecoder(resp.Body).Decode(&jsonResp)
	t.Log("jsonResp:", jsonResp)
}

func TestDecode(t *testing.T) {
	client := &http.Client{
		Timeout: 1 * time.Second,
	}
	code := "OTv0FdGU8Ng"
	r, _ := http.NewRequest("GET", "http://localhost:8082/"+code, nil)

	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	t.Log("body:", string(body))
}
