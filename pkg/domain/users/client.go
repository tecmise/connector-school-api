package users

import (
	"context"
	"fmt"

	"github.com/tecmise/connector-school-api/pkg/adapters/outbound/lambda"
	"github.com/tecmise/connector-school-api/pkg/adapters/outbound/rest"
	"github.com/tecmise/connector-school-api/pkg/ports/output/connector"
)

type (
	Client interface {
		FindUserInfo(ctx context.Context, userID string) (Response, error)
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

func (c client) FindUserInfo(ctx context.Context, userID string) (Response, error) {
	var user Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource(fmt.Sprintf("api/users/%s", userID)).
		WithCredentials(ctx).
		WithMethod("GET").
		Build()
	return user, c.mapper.Find(parameter, &user)
}
