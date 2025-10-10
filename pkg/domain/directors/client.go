package directors

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	lambda "github.com/tecmise/connector-lib/pkg/adapters/outbound/client_lambda"
	rest "github.com/tecmise/connector-lib/pkg/adapters/outbound/client_rest"
	"github.com/tecmise/connector-lib/pkg/ports/output/connector"
)

type (
	Client interface {
		PaginateDirectors(ctx context.Context, search string, page int, limit int, sort string) (connector.ListResponse[Response], error)
		CreateDirector(ctx context.Context, request any) (Response, error)
		UpdateDirector(ctx context.Context, request any) (Response, error)
		InativeDirector(ctx context.Context, directorID string) (Response, error)
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

func (c client) PaginateDirectors(ctx context.Context, search string, page int, limit int, sort string) (connector.ListResponse[Response], error) {
	var list connector.ListResponse[Response]
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithCredentials(ctx).
		WithResource(fmt.Sprintf("api/directors?search=%s&page=%d&page=%d&page=%s", search, page, limit, sort)).
		WithMethod("GET").
		Build()
	return list, c.mapper.Page(parameter, &list)
}

func (c client) CreateDirector(ctx context.Context, request any) (Response, error) {
	logrus.Debug("CreateDirector")
	var directors Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource("api/directors").
		WithCredentials(ctx).
		WithBody(request).
		WithMethod("POST").
		WithHeader("Content-Type", "multipart/form-data").
		Build()
	return directors, c.mapper.Create(parameter, &directors)
}

func (c client) UpdateDirector(ctx context.Context, request any) (Response, error) {
	logrus.Debug("CreateDirector")
	var directors Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithCredentials(ctx).
		WithResource("api/directors").
		WithBody(request).
		WithMethod("PUT").
		WithHeader("Content-Type", "multipart/form-data").
		Build()
	return directors, c.mapper.Update(parameter, &directors)
}

func (c client) InativeDirector(ctx context.Context, directorID string) (Response, error) {
	var directors Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithCredentials(ctx).
		WithResource(fmt.Sprintf("api/directors/%s", directorID)).
		WithMethod("DELETE").
		Build()
	return directors, c.mapper.Inative(parameter, &directors)
}
