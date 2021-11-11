package httputil

import (
	"encoding/json"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func GetJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	//body, err := ioutil.ReadAll(r.Body)

	//fmt.Println("RESPONSE ", r.Body)

	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
