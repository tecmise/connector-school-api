package grades

import (
	"context"
	lambda "github.com/tecmise/connector-lib/pkg/adapters/outbound/client_lambda_proxy"
	rest "github.com/tecmise/connector-lib/pkg/adapters/outbound/client_rest"
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
		WithResource("api/grades").
		WithMethod("GET").
		Build()
	return response, c.mapper.List(parameter, &response)
}
