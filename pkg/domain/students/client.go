package students

import (
	"context"
	"fmt"

	"github.com/tecmise/connector-school-api/pkg/adapters/outbound/lambda"
	"github.com/tecmise/connector-school-api/pkg/adapters/outbound/rest"
	"github.com/tecmise/connector-school-api/pkg/ports/output/connector"
)

type (
	Client interface {
		PaginateStudents(ctx context.Context, search string, page int, limit int, sort string) (connector.ListResponse[Response], error)
		CreateStudent(ctx context.Context, request any) (Response, error)
		UpdateStudent(ctx context.Context, request any) (Response, error)
		InativeStudent(ctx context.Context, studentID string) (Response, error)
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

func (c client) PaginateStudents(ctx context.Context, search string, page int, limit int, sort string) (connector.ListResponse[Response], error) {
	var list connector.ListResponse[Response]
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithHeader("Authorization", fmt.Sprintf("Bearer %s", ctx.Value("bearer-token").(string))).
		WithHeader("x-api-key", ctx.Value("x-api-key").(string)).
		WithResource(fmt.Sprintf("api/students?search=%s&page=%d&page=%d&page=%s", search, page, limit, sort)).
		WithMethod("GET").
		Build()
	return list, c.mapper.Page(parameter, &list)
}

func (c client) CreateStudent(ctx context.Context, request any) (Response, error) {
	var students Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource("api/students").
		WithHeader("Authorization", fmt.Sprintf("Bearer %s", ctx.Value("bearer-token").(string))).
		WithHeader("x-api-key", ctx.Value("x-api-key").(string)).
		WithBody(request).
		WithMethod("POST").
		WithHeader("Content-Type", "multipart/form-data").
		Build()
	return students, c.mapper.Create(parameter, &students)
}

func (c client) UpdateStudent(ctx context.Context, request any) (Response, error) {
	var students Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithHeader("Authorization", fmt.Sprintf("Bearer %s", ctx.Value("bearer-token").(string))).
		WithHeader("x-api-key", ctx.Value("x-api-key").(string)).
		WithResource("api/students").
		WithBody(request).
		WithMethod("PUT").
		WithHeader("Content-Type", "multipart/form-data").
		Build()
	return students, c.mapper.Update(parameter, &students)
}

func (c client) InativeStudent(ctx context.Context, studentID string) (Response, error) {
	var students Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithHeader("Authorization", fmt.Sprintf("Bearer %s", ctx.Value("bearer-token").(string))).
		WithHeader("x-api-key", ctx.Value("x-api-key").(string)).
		WithResource(fmt.Sprintf("api/students/%s", studentID)).
		WithMethod("DELETE").
		Build()
	return students, c.mapper.Inative(parameter, &students)
}
