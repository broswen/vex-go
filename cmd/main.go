package main

import (
	"context"
	"github.com/broswen/vex-go/vex"
	"log"
)

func main() {

	c, err := vex.New("",
		vex.WithHost("https://vex.broswen.com"),
		vex.WithDebug(false))
	if err != nil {
		log.Fatal(err)
	}
	a, err := c.GetAccount(context.Background(), "f07f5da1-bfcc-42ad-9794-f251c6837ad5")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v", a)
	projects, err := c.GetProjects(context.Background(), a.ID)
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range projects {
		p2, err := c.GetProject(context.Background(), a.ID, p.ID)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%#v", p2)
		flags, err := c.GetFlags(context.Background(), a.ID, p2.ID)
		if err != nil {
			log.Fatal(err)
		}
		for _, f := range flags {
			log.Printf("%#v", f)
		}
	}
}
