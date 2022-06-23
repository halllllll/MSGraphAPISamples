package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type TokenResp struct {
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	ExtExpiresIn int    `json:"ext_expires_in"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func main() {
	tenant_id := os.Getenv("AZURE_TENANT_ID")
	client_id := os.Getenv("AZURE_CLIENT_ID")
	client_secret := os.Getenv("AZURE_CLIENT_SECRET")
	refresh_token := os.Getenv("AZURE_REFRESH_TOKEN")
	fmt.Println(tenant_id, client_id, client_secret, refresh_token)
	if tenant_id == "" || client_id == "" || client_secret == "" || refresh_token == "" {
		log.Fatalln("required value not found")
	}
	endpoint := fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/token", tenant_id)

	values := url.Values{}
	values.Add("grant_type", "refresh_token")
	values.Add("client_id", client_id)
	values.Add("client_secret", client_secret)
	values.Add("refresh_token", refresh_token)
	client := &http.Client{}
	req, err := http.NewRequest(
		"POST",
		endpoint,
		strings.NewReader(values.Encode()),
	)
	if err != nil {
		log.Fatalln(err)
	}
	// やり方が間違っているのか、以下ふたつともダメだった。SetじゃなくてAddにしてもダメだった
	// なぜかむしろヘッダー設定しなかったらいけた。デフォルトがなんなのかは知らない
	// req.Header.Set("Content-Type", "multipart/form-data")
	// req.Header.Set("Content-Type", "application/json")

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
	// json
	var resp TokenResp
	if err := json.Unmarshal([]byte(body), &resp); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", resp)
	fmt.Println("access token: ", resp.AccessToken)
}
