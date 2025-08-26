package levels

import (
	"context"

	"github.com/tecmise/connector-school-api/pkg/adapters/outbound/lambda"
	"github.com/tecmise/connector-school-api/pkg/adapters/outbound/rest"
	"github.com/tecmise/connector-school-api/pkg/ports/output/connector"
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

func (c client) Select(_ context.Context) ([]Response, error) {
	var response []Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource("api/levels").
		WithMethod("GET").
		Build()
	return response, c.mapper.List(parameter, &response)
}
