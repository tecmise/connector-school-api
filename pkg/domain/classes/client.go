package classes

import (
	"context"
	"fmt"

	"github.com/tecmise/connector-school-api/pkg/adapters/outbound/lambda"
	"github.com/tecmise/connector-school-api/pkg/adapters/outbound/rest"
	"github.com/tecmise/connector-school-api/pkg/ports/output/connector"
)

type (
	Client interface {
		PaginateClasses(ctx context.Context, search string, page int, limit int, sort string) (connector.ListResponse[Response], error)
		FindBySchoolID(ctx context.Context, schoolID string) ([]Response, error)
		CreateClass(ctx context.Context, request any) (Response, error)
		UpdateClass(ctx context.Context, request any) (Response, error)
		InativeClass(ctx context.Context, classID string) (Response, error)
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

func (c client) PaginateClasses(_ context.Context, search string, page int, limit int, sort string) (connector.ListResponse[Response], error) {
	var classes connector.ListResponse[Response]
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource(fmt.Sprintf("api/classes?search=%s&page=%d&page=%d&page=%s", search, page, limit, sort)).
		WithMethod("GET").
		Build()
	return classes, c.mapper.Page(parameter, &classes)
}

func (c client) FindBySchoolID(_ context.Context, schoolID string) ([]Response, error) {
	var classes []Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource(fmt.Sprintf("api/classes/select/school/%s", schoolID)).
		WithMethod("GET").
		Build()
	return classes, c.mapper.List(parameter, &classes)
}

func (c client) CreateClass(_ context.Context, request any) (Response, error) {
	var classes Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource("api/classes").
		WithBody(request).
		WithMethod("POST").
		Build()
	return classes, c.mapper.Create(parameter, &classes)
}

func (c client) UpdateClass(_ context.Context, request any) (Response, error) {
	var classes Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource("api/classes").
		WithBody(request).
		WithMethod("PUT").
		Build()
	return classes, c.mapper.Update(parameter, &classes)
}

func (c client) InativeClass(_ context.Context, classID string) (Response, error) {
	var classes Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource(fmt.Sprintf("api/classes/%s", classID)).
		WithMethod("DELETE").
		Build()
	return classes, c.mapper.Inative(parameter, &classes)
}
