package levels

import (
	"context"
	"github.com/tecmise/connector-lib/pkg/adapters/outbound/lambda"
	"github.com/tecmise/connector-lib/pkg/adapters/outbound/rest"
	"github.com/tecmise/connector-lib/pkg/ports/output/connector"
)

type (
	Client interface {
		Select(ctx context.Context) ([]Response, error)
	}

	client struct {
		mapper connector.Call[Response]
		host   string
	}
)

func Rest(host string) Client {
	return &client{
		host:   host,
		mapper: rest.NewConnector[Response](),
	}
}

func Lambda(identifier string) Client {
	return &client{
		host:   identifier,
		mapper: lambda.NewConnector[Response](),
	}
}

func (c client) Select(ctx context.Context) ([]Response, error) {
	var response []Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithCredentials(ctx).
		WithResource("api/levels").
		WithMethod("GET").
		Build()
	return response, c.mapper.List(parameter, &response)
}
