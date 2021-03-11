package usecase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetList() {

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

func GetDetail(id string) (resp interface{}, err error) {

	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=faf7e5bb&i=%s", id)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	json.Unmarshal(body, &resp)
	return
}
