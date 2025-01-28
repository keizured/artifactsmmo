package main

import (
	"fmt"
	"log"
	"os"

	"artifactsmmo/pkg/api"

	"github.com/joho/godotenv"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load .env file. Err: %s", err)
	}

	token, exists := os.LookupEnv("API_TOKEN")
	if !exists {
		log.Fatal("Could not load read API_TOKEN")
	}

	user := api.NewArtifactsUser(token)
	user.SetToken(token)

	status, err := user.Status()

	if err != nil {
		log.Panic(err)
	}
	fmt.Println(status)

	// Loop for cow farming

	for {
		response, err := user.ActionFight("Azazel", true)
		if err != nil {
			log.Panic(err)
		}

		err = user.CooldownFromResponse(response)
		if err != nil {
			log.Panic(err)
		}

		response, err = user.ActionRest("Azazel", true)
		if err != nil {
			log.Panic(err)
		}

		err = user.CooldownFromResponse(response)
		if err != nil {
			log.Panic(err)
		}

		rawBeefCount, err := user.ItemCount(response, "raw_beef")
		if err != nil {
			log.Panic(err)
		}

		milkBucket, err := user.ItemCount(response, "milk_bucket")
		if err != nil {
			log.Panic(err)
		}

		cowHideCount, err := user.ItemCount(response, "cowhide")
		if err != nil {
			log.Panic(err)
		}

		if rawBeefCount >= 10 {
			response, err := user.ActionDeleteItem("Azazel", "raw_beef", 10, true)
			if err != nil {
				log.Panic(err)
			}

			err = user.CooldownFromResponse(response)
			if err != nil {
				log.Panic(err)
			}
		}

		if milkBucket >= 10 {
			response, err := user.ActionDeleteItem("Azazel", "milk_bucket", 10, true)
			if err != nil {
				log.Panic(err)
			}

			err = user.CooldownFromResponse(response)
			if err != nil {
				log.Panic(err)
			}
		}

		if cowHideCount >= 10 {
			response, err := user.ActionDeleteItem("Azazel", "cowhide", 10, true)
			if err != nil {
				log.Panic(err)
			}

			err = user.CooldownFromResponse(response)
			if err != nil {
				log.Panic(err)
			}
		}
	}

}
