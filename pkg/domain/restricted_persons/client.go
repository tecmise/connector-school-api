package restricted_persons

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	lambda "github.com/tecmise/connector-lib/pkg/adapters/outbound/client_lambda_proxy"
	rest "github.com/tecmise/connector-lib/pkg/adapters/outbound/client_rest"
	"github.com/tecmise/connector-lib/pkg/ports/output/connector"
)

type (
	Client interface {
		PaginateRestricted(ctx context.Context, search string, page int, limit int, sort string) (connector.ListResponse[Response], error)
		CreateRestricted(ctx context.Context, request any) (Response, error)
		UpdateRestricted(ctx context.Context, request any) (Response, error)
		InativeRestricted(ctx context.Context, restrictedID string) (Response, error)
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

func (c client) PaginateRestricted(ctx context.Context, search string, page int, limit int, sort string) (connector.ListResponse[Response], error) {
	var list connector.ListResponse[Response]
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithCredentials(ctx).
		WithResource(fmt.Sprintf("api/restrictedPersons?search=%s&page=%d&page=%d&page=%s", search, page, limit, sort)).
		WithMethod("GET").
		Build()
	return list, c.mapper.Page(parameter, &list)
}

func (c client) CreateRestricted(ctx context.Context, request any) (Response, error) {
	logrus.Debug("CreateRestricted")
	var restricted Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource("api/restrictedPersons").
		WithCredentials(ctx).
		WithBody(request).
		WithMethod("POST").
		WithHeader("Content-Type", "multipart/form-data").
		Build()
	return restricted, c.mapper.Create(parameter, &restricted)
}

func (c client) UpdateRestricted(ctx context.Context, request any) (Response, error) {
	logrus.Debug("CreateRestricted")
	var restricted Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithCredentials(ctx).
		WithResource("api/restrictedPersons").
		WithBody(request).
		WithMethod("PUT").
		WithHeader("Content-Type", "multipart/form-data").
		Build()
	return restricted, c.mapper.Update(parameter, &restricted)
}

func (c client) InativeRestricted(ctx context.Context, restrictedID string) (Response, error) {
	var restricted Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithCredentials(ctx).
		WithResource(fmt.Sprintf("api/restrictedPersons/%s", restrictedID)).
		WithMethod("DELETE").
		Build()
	return restricted, c.mapper.Inative(parameter, &restricted)
}
