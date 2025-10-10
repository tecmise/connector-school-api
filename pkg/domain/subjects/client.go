package subjects

import (
	"context"
	"fmt"

	lambda "github.com/tecmise/connector-lib/pkg/adapters/outbound/client_lambda_proxy"
	rest "github.com/tecmise/connector-lib/pkg/adapters/outbound/client_rest"
	"github.com/tecmise/connector-lib/pkg/ports/output/connector"
)

type (
	Client interface {
		Select(ctx context.Context) ([]Response, error)
		PaginateSubjects(ctx context.Context, search string, page int, limit int, sort string) (connector.ListResponse[Response], error)
		CreateSubject(ctx context.Context, request any) (Response, error)
		UpdateSubject(ctx context.Context, request any) (Response, error)
		InativeSubject(ctx context.Context, subjectID string) (Response, error)
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
	var subjects []Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource("api/subjects/select").
		WithMethod("GET").
		WithCredentials(ctx).
		Build()
	return subjects, c.mapper.List(parameter, &subjects)
}

func (c client) PaginateSubjects(ctx context.Context, search string, page int, limit int, sort string) (connector.ListResponse[Response], error) {
	var list connector.ListResponse[Response]
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource(fmt.Sprintf("api/subjects?search=%s&page=%d&page=%d&page=%s", search, page, limit, sort)).
		WithMethod("GET").
		WithCredentials(ctx).
		Build()
	return list, c.mapper.Page(parameter, &list)
}

func (c client) CreateSubject(ctx context.Context, request any) (Response, error) {
	var subject Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource("api/subjects").
		WithBody(request).
		WithCredentials(ctx).
		WithMethod("POST").
		Build()
	return subject, c.mapper.Create(parameter, &subject)
}

func (c client) UpdateSubject(ctx context.Context, request any) (Response, error) {
	var subject Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource("api/subjects").
		WithBody(request).
		WithCredentials(ctx).
		WithMethod("PUT").
		Build()
	return subject, c.mapper.Update(parameter, &subject)
}

func (c client) InativeSubject(ctx context.Context, subjectID string) (Response, error) {
	var subject Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithCredentials(ctx).
		WithResource(fmt.Sprintf("api/subjects/%s", subjectID)).
		WithMethod("DELETE").
		Build()
	return subject, c.mapper.Inative(parameter, &subject)
}
