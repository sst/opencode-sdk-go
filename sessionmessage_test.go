// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package opencode_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/opencode-go"
	"github.com/stainless-sdks/opencode-go/internal/testutil"
	"github.com/stainless-sdks/opencode-go/option"
)

func TestSessionMessageNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Session.Message.New(
		context.TODO(),
		"id",
		opencode.SessionMessageNewParams{
			ModelID: "modelID",
			Parts: []opencode.SessionMessageNewParamsPartUnion{{
				OfText: &opencode.SessionMessageNewParamsPartText{
					Text:      "text",
					ID:        opencode.String("id"),
					Synthetic: opencode.Bool(true),
					Time: opencode.SessionMessageNewParamsPartTextTime{
						Start: 0,
						End:   opencode.Float(0),
					},
				},
			}},
			ProviderID: "providerID",
			Agent:      opencode.String("agent"),
			MessageID:  opencode.String("msg"),
			System:     opencode.String("system"),
			Tools: map[string]bool{
				"foo": true,
			},
		},
	)
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionMessageGet(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Session.Message.Get(
		context.TODO(),
		"messageID",
		opencode.SessionMessageGetParams{
			ID: "id",
		},
	)
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSessionMessageList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := opencode.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Session.Message.List(context.TODO(), "id")
	if err != nil {
		var apierr *opencode.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
