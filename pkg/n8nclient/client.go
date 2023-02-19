// This file holds the go generate directive to build the OpenAPI (Swagger) driven client
//
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen -config ../../codegen/n8nclient/config.yaml ../../codegen/n8nclient/openapi-bundle.yml
package n8nclient

import (
	"context"
	"net/http"
	"net/url"
)

// The oapi-codegen tool is good, but not incredible. We need to add in a couple things

// WithToken includes the API Key in all requests
func WithToken(token string) func(context.Context, *http.Request) error {
	return func(ctx context.Context, req *http.Request) error {

		// TODO: I can't fully remember if this freaks out on nil req.Header
		req.Header.Add("accept", "application/json")
		req.Header.Set("X-N8N-API-KEY", token)

		return nil
	}
}

// WithVersion sets the targetted API version on all requests
// this isn't included in the OpenAPI definition
func WithVersion(version string) func(context.Context, *http.Request) error {
	return func(ctx context.Context, req *http.Request) error {

		newBase := "/api/v" + version
		req.URL.Path, _ = url.JoinPath(newBase, req.URL.Path)

		return nil
	}
}
