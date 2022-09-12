package main

import (
	"context"
	"fmt"
	"github.com/broswen/vex-go"
	"log"
)

func main() {
	accountID := "account id here"
	apiToken := "api token here"
	//set WithDebug(true) to see raw http calls in logs
	client, err := vex_go.New(apiToken, vex_go.WithDebug(false))
	if err != nil {
		log.Fatal(err)
	}
	projects, err := client.GetProjects(context.Background(), accountID)
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range projects {
		fmt.Println("==============================")
		fmt.Println("ID:", p.ID)
		fmt.Println("Name:", p.Name)
		fmt.Println("Description:", p.Description)
		flags, err := client.GetFlags(context.Background(), accountID, p.ID)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println("------------------------------")
		for _, f := range flags {
			fmt.Println("	ID:", f.ID)
			fmt.Println("	Key:", f.Key)
			fmt.Println("	Type:", f.Type)
			fmt.Println("	Value:", f.Value)
		}
	}

}
