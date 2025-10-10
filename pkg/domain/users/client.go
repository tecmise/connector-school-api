package users

import (
	"context"
	"fmt"

	"github.com/tecmise/connector-lib/pkg/adapters/outbound/lambda"
	"github.com/tecmise/connector-lib/pkg/adapters/outbound/rest"
	"github.com/tecmise/connector-lib/pkg/ports/output/connector"
)

type (
	Client interface {
		FindUserInfo(ctx context.Context, userID string) (Response, error)
		PaginateUsers(ctx context.Context, search string, page int, limit int, sort string) (connector.ListResponse[Response], error)
		CreateUser(ctx context.Context, request any) (Response, error)
		UpdateUser(ctx context.Context, request any) (Response, error)
		InativeUser(ctx context.Context, userID string) (Response, error)
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

func (c client) PaginateUsers(ctx context.Context, search string, page int, limit int, sort string) (connector.ListResponse[Response], error) {
	var users connector.ListResponse[Response]
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource(fmt.Sprintf("api/users?search=%s&page=%d&page=%d&page=%s", search, page, limit, sort)).
		WithCredentials(ctx).
		WithMethod("GET").
		Build()
	return users, c.mapper.Page(parameter, &users)
}

func (c client) CreateUser(ctx context.Context, request any) (Response, error) {
	var user Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource("api/users").
		WithCredentials(ctx).
		WithMethod("POST").
		WithBody(request).
		Build()
	return user, c.mapper.Create(parameter, &user)
}

func (c client) UpdateUser(ctx context.Context, request any) (Response, error) {
	var user Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource("api/users").
		WithCredentials(ctx).
		WithMethod("PUT").
		WithBody(request).
		Build()
	return user, c.mapper.Update(parameter, &user)
}

func (c client) InativeUser(ctx context.Context, userID string) (Response, error) {
	var user Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource(fmt.Sprintf("api/users/%s", userID)).
		WithCredentials(ctx).
		WithMethod("DELETE").
		Build()
	return user, c.mapper.Inative(parameter, &user)
}
