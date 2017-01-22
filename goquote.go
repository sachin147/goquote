package main

import (
	"fmt"
        "flag"
        "net/http"
        "io/ioutil"
        "encoding/json"
)

type Provider struct {
	Name string `json:"name"`
	Id int `json:"id"`
}

type WordExample struct {
	Year int `json:"year"`
	Provider Provider 'json:"year"`
	Url string `json:"url"`
	Word string `json:"word"`
        Text string `json:"text"`
	Title string `json:"title"`
	DocumentId int `json:"documentId"`
	ExampleId int `json:"exampleId"`
	Rating long 'json:"rating"`
        
}

type Quote struct {
    Quote string `json:"quote"`
    Author string `json:"author"`
    Category string `json:"category"`
}

func main() {
	quoteOfDay := flag.String("quoteofday", "movies", "Get quote of the day")
	flag.Parse()
	category := *quoteOfDay
	url := "https://andruxnet-random-famous-quotes.p.mashape.com/?cat="
        postUrl := url + category
	client := &http.Client{}
	req, err := http.NewRequest("POST", postUrl, nil)
	req.Header.Set("X-Mashape-Key","cqUWSfKEA5mshQ2qXdS2FQICJDV9p1BqhONjsnGY2uHwajI0iT")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
        defer resp.Body.Close()      
        if err != nil {
           panic(err.Error())
        }
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
           panic(err.Error())
        }	
	quote, err := ParseQuote(body)
	fmt.Println("=======QUOTE OF THE DAY=======")
	fmt.Println("Quote - ", quote.Quote)
        fmt.Println("Author - ", quote.Author)
        fmt.Println("Category - ", quote.Category)
	
}



func ParseQuote(quotebody []byte) (*Quote, error) {
    var quote Quote
    err := json.Unmarshal(quotebody, &quote)
    if(err != nil){
        fmt.Println("Error ", err)
    }
    return &quote, err
}




