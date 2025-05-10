package opensea

import (
	"context"
	"net/http"
)

func ApiKeyRequestEditorFn(apiKey string) ClientOption {
	return WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Set("x-api-key", apiKey)
		return nil
	})
}
