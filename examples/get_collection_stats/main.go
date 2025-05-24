package main

import (
	"context"
	"fmt"
	"github.com/stein-f/go-opensea"
	"os"
)

const (
	baseURL        = "https://api.opensea.io"
	collectionSlug = "proof-moonbirds"
)

func main() {
	apiKey := os.Getenv("OPENSEA_API_KEY")
	ctx := context.Background()

	openseaClient, err := opensea.NewClientWithResponses(baseURL, opensea.ApiKeyRequestEditorFn(apiKey))
	if err != nil {
		panic(err)
	}

	response, err := openseaClient.GetCollectionStatsWithResponse(ctx, collectionSlug)
	if err != nil {
		panic(err)
	}

	if response.StatusCode() != 200 {
		panic(fmt.Sprintf("unexpected status code: %d", response.StatusCode()))
	}

	fmt.Println(response.JSON200.Total.FloorPrice)
}
