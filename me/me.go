package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type MeResp struct {
	OdataContext      string        `json:"@odata.context"`
	BusinessPhones    []interface{} `json:"businessPhones"`
	DisplayName       string        `json:"displayName"`
	GivenName         string        `json:"givenName"`
	JobTitle          interface{}   `json:"jobTitle"`
	Mail              string        `json:"mail"`
	MobilePhone       interface{}   `json:"mobilePhone"`
	OfficeLocation    interface{}   `json:"officeLocation"`
	PreferredLanguage interface{}   `json:"preferredLanguage"`
	Surname           string        `json:"surname"`
	UserPrincipalName string        `json:"userPrincipalName"`
	ID                string        `json:"id"`
}

func main() {
	url := "https://graph.microsoft.com/v1.0/me"
	token := os.Getenv("AZURE_ACCESS_TOKEN")
	req, err := http.NewRequest(
		"GET",
		url,
		nil,
	)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res.Status)
	if res.StatusCode != 200 {
		log.Fatalf("network error?\n%s\n", err)
	}

	// json
	var resp MeResp
	if err := json.Unmarshal([]byte(body), &resp); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", resp)
	fmt.Println("context: ", resp.OdataContext)
	fmt.Println("display name: ", resp.DisplayName)
	fmt.Println("principal name: ", resp.UserPrincipalName)
}
