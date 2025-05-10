package main

import (
	"context"
	"fmt"
	"github.com/stein-f/go-opensea"
	"os"
)

const (
	baseURL = "https://api.opensea.io"
	address = "GordoGordo"
)

func main() {
	apiKey := os.Getenv("OPENSEA_API_KEY")
	ctx := context.Background()

	openseaClient, err := opensea.NewClientWithResponses(baseURL, opensea.ApiKeyRequestEditorFn(apiKey))
	if err != nil {
		panic(err)
	}

	response, err := openseaClient.GetAccountWithResponse(ctx, address)
	if err != nil {
		panic(err)
	}

	if response.StatusCode() != 200 {
		panic(fmt.Sprintf("unexpected status code: %d", response.StatusCode()))
	}

	collectionResponse, err := openseaClient.GetCollectionWithResponse(ctx, "proof-moonbirds")
	if err != nil {
		panic(err)
	}

	fmt.Println(collectionResponse.JSON200)
}
