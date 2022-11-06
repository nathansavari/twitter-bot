package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)


func main(){

	getDate()

	// load .env file from given path
	// we keep it empty it will load .env from current directory

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// getting env variables 

	consumerKey, consumerSecret, accessToken, accessTokenSecret := getEnv()

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

		if getDate() == "21:12:00" {
			
		// the tweet

		date := getDate()

		_, _, err = client.Statuses.Update(date, nil)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		} else {
			fmt.Println("Twitted successfully!")
		}

		} 

	}


}

func getEnv() (string, string, string, string) {

	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN_KEY")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")

	return consumerKey, consumerSecret, accessToken, accessTokenSecret
}

func getDate() string {

    dt := time.Now()
	now := dt.Format(("15:04:05"))

	return now
}