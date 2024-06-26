package client

import (
	"context"
	"net/http"
	"net/http/httptrace"

	"github.com/go-resty/resty/v2"
	"go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type Client interface {
	StarWars() StarWarsClient
}

type client struct {
	restyClient    *resty.Client
	starWarsClient StarWarsClient
}

func NewClient() Client {
	c := resty.NewWithClient(
		&http.Client{
			Transport: otelhttp.NewTransport(http.DefaultTransport,
				otelhttp.WithClientTrace(func(ctx context.Context) *httptrace.ClientTrace {
					return otelhttptrace.NewClientTrace(ctx)
				})),
		},
	)

	return &client{
		restyClient:    c,
		starWarsClient: newStarWarsClient(c),
	}
}

func (c *client) StarWars() StarWarsClient {
	return c.starWarsClient
}
