package client

import (
	"context"

	"github.com/go-resty/resty/v2"
)

type StarWarsClient interface {
	GetPerson(
		ctx context.Context,
		id string,
	) (map[string]any, error)
}

type starWarsClient struct {
	restyClient *resty.Client
}

func newStarWarsClient(
	restyClient *resty.Client,
) StarWarsClient {
	return &starWarsClient{
		restyClient: restyClient,
	}
}

func (c *starWarsClient) GetPerson(
	ctx context.Context,
	id string,
) (map[string]any, error) {
	res := map[string]any{}

	resp, err := c.restyClient.R().
		SetContext(ctx).
		SetPathParams(map[string]string{
			"id": id,
		}).
		SetResult(&res).
		Get("https://swapi.dev/api/people/{id}/")
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, err
	}

	return res, nil
}
