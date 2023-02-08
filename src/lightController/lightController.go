package lightController

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetTriggerKeyMatchingTimeOfADay() string {
	URL := fmt.Sprintf("%s", os.Getenv("TRIGGER_API_SERVER_ADDRESS"))
	resp, err := http.Get(URL)

	if err != nil {
		fmt.Printf("Request Failed: %s", err)
		panic("Request Failed")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Reading body failed: %s", err)
		panic("Reading body failed")
	}
	var s = new(Response)
	err2 := json.Unmarshal(body, &s)
	if err2 != nil {
		fmt.Println("Error reading response from Trigger API")
		fmt.Println(err2)
	}

	return s.Current
}

type Response struct {
	Current string `json:"current"`
}
