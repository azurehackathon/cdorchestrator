package main

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
)

type response struct {
	BuildNumber string `json:"BUILD_NUMBER"`
}
func main(){
	if len(os.Args) < 5 {
		fmt.Println("Required arguments are missing.")
		fmt.Println("Usage: go run main.go buildNumber url1 url2 url3")
		return
	}
	currentBuildNumber := os.Args[1]
	urls := os.Args[2:]
	result := make([]bool, len(urls))
	
	client := &http.Client{}
	
	for index,url := range(urls) {
		result[index] = false
		req, err := http.NewRequest(http.MethodGet, url , nil)
		if err!= nil{
			fmt.Println("Error in creating a request for ",url)
		}
		resp, err := client.Do(req)
		if err!= nil{
			fmt.Println("Error in making a request for ",url)
		}

		if(resp.StatusCode != http.StatusOK) {
			fmt.Println("Code received ", resp.StatusCode, "for ", url, "expected ", http.StatusOK)
		} else {
			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error reading the response body for ", url)
			}
			temp:=response{}
			if err:=json.Unmarshal(data, &temp);err!=nil{
				fmt.Println("Error parsing the response body for ", url)
			}
			//check the build number
			if currentBuildNumber == temp.BuildNumber{
				result[index] = true
			}
		}
	}
	fmt.Println(result)
}