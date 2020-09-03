package sendjson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func SendScrape(game string, price string, sitename string, siteurl string, seller string) {

	url := os.Getenv("API_URL")

	fmt.Println(url)
	timeToString := time.Now().Format(time.RFC3339)

	values := map[string]string{"game": game, "siteName": sitename, "siteURL": siteurl, "FetchTime": timeToString, "PriceValue": price, "seller": seller}
	jsonValue, _ := json.Marshal(values)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}
	req.Header.Set("API-KEY", os.Getenv("API_KEY"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))

	defer resp.Body.Close()

}
