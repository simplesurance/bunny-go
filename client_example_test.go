package bunny_test

import (
	"context"
	"fmt"
	"log"
	"os"

	bunny "github.com/simplesurance/bunny-go"
)

func Example() {
	apiKey := os.Getenv("BUNNY_API_KEY")
	clt := bunny.NewClient(apiKey)

	pz, err := clt.PullZone.Get(context.Background(), 1234)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("pull zone name: %s\n", *pz.Name)
}
