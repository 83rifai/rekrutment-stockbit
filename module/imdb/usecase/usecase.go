package usecase

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Get() {

	url := "http://www.omdbapi.com/?apikey=faf7e5bb&s=Batman&page=2"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
