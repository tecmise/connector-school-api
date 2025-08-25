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
		FindByClusterId(ctx context.Context, clusterId int64) ([]int64, error)
		PaginateSchools(ctx context.Context, search string, page int, limit int, sort string) (connector.ListResponse[Response], error)
		CreateSchool(ctx context.Context, request any) (Response, error)
		UpdateSchool(ctx context.Context, request Response) (Response, error)
		InativeSchool(ctx context.Context, schoolID uint) (Response, error)
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

func (c client) FindByClusterId(_ context.Context, clusterId int64) ([]int64, error) {
	var list []int64
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource(fmt.Sprintf("api/schools/cluster/%d/ids", clusterId)).
		WithMethod("GET").
		Build()
	return list, c.mapper.Ids(parameter, &list)
}

func (c client) PaginateSchools(_ context.Context, search string, page int, limit int, sort string) (connector.ListResponse[Response], error) {
	var list connector.ListResponse[Response]
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource(fmt.Sprintf("api/schools?search=%s&page=%d&page=%d&page=%s", search, page, limit, sort)).
		WithMethod("GET").
		Build()
	return list, c.mapper.Page(parameter, &list)
}

func (c client) CreateSchool(_ context.Context, request any) (Response, error) {
	var school Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource("api/schools").
		WithBody(request).
		WithMethod("POST").
		Build()
	return school, c.mapper.Create(parameter, &school)
}

func (c client) UpdateSchool(_ context.Context, request Response) (Response, error) {
	var school Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource("api/schools").
		WithBody(request).
		WithMethod("PUT").
		Build()
	return school, c.mapper.Update(parameter, &school)
}

func (c client) InativeSchool(_ context.Context, schoolID uint) (Response, error) {
	var school Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource(fmt.Sprintf("api/schools/%d", schoolID)).
		WithMethod("PUT").
		Build()
	return school, c.mapper.Inative(parameter, &school)
}
