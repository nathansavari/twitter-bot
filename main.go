package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

func main() {

	getDate()

	// load .env file from given path
	// we keep it empty it will load .env from current directory

	if os.Getenv("ENV") != "prod" {
		err := godotenv.Load(".env")

		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}

	// getting env variables

	consumerKey, consumerSecret, accessToken, accessTokenSecret, apiKey := getEnv()

	fact := getData(apiKey)

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	// user verification

	user, _, err := client.Accounts.VerifyCredentials(&twitter.AccountVerifyParams{})
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	fmt.Printf("Account: @%s (%s)\n", user.ScreenName, user.Name)

	for {

		if getDate() == "12:00:00" {

			// the tweet

			_, _, err = client.Statuses.Update(fact, nil)
			if err != nil {
				fmt.Printf("err: %v\n", err)
			} else {
				fmt.Println("Twitted successfully!")
			}

		}

	}

}

func getEnv() (string, string, string, string, string) {

	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN_KEY")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	apiKey := os.Getenv("API_KEY")

	return consumerKey, consumerSecret, accessToken, accessTokenSecret, apiKey
}

func getDate() string {

	dt := time.Now()
	now := dt.Format(("15:04:05"))

	return now
}

func getData(apiKey string) string {

	type Obj struct {
		Fact string `json:"fact"`
	}

	// init the objet

	var data []Obj

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.api-ninjas.com/v1/facts?limit=1", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("X-Api-Key", apiKey)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// fetch the body & transform it into bytes

	body, err := ioutil.ReadAll(resp.Body)

	// put bytes in the object

	json.Unmarshal(body, &data)

	fact := data[0].Fact

	return fact
}
