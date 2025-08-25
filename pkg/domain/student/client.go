package schools

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
		UpdateStudent(ctx context.Context, request Response) (Response, error)
		InativeStudent(ctx context.Context, schoolID uint) (Response, error)
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

func (c client) PaginateStudents(_ context.Context, search string, page int, limit int, sort string) (connector.ListResponse[Response], error) {
	var list connector.ListResponse[Response]
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource(fmt.Sprintf("api/students?search=%s&page=%d&page=%d&page=%s", search, page, limit, sort)).
		WithMethod("GET").
		Build()
	return list, c.mapper.Page(parameter, &list)
}

func (c client) CreateStudent(_ context.Context, request any) (Response, error) {
	var students Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource("api/students").
		WithBody(request).
		WithMethod("POST").
		Build()
	return students, c.mapper.Create(parameter, &students)
}

func (c client) UpdateStudent(_ context.Context, request Response) (Response, error) {
	var students Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource("api/students").
		WithBody(request).
		WithMethod("PUT").
		Build()
	return students, c.mapper.Update(parameter, &students)
}

func (c client) InativeStudent(_ context.Context, studentID uint) (Response, error) {
	var students Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource(fmt.Sprintf("api/students/%d", studentID)).
		WithMethod("PUT").
		Build()
	return students, c.mapper.Inative(parameter, &students)
}
