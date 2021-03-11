package model

type Anagram struct {
	Data []string `json:"data"`
}

type ResponseAnagram struct {
	Request []string    `json:"request"`
	Result  interface{} `json:"result"`
}
