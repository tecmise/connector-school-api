package schools

import (
	"context"
	"fmt"

	lambda "github.com/tecmise/connector-lib/pkg/adapters/outbound/client_lambda_proxy"
	rest "github.com/tecmise/connector-lib/pkg/adapters/outbound/client_rest"
	"github.com/tecmise/connector-lib/pkg/ports/output/connector"
)

type (
	Client interface {
		FindByClusterId(ctx context.Context, clusterId int64) ([]int64, error)
		Select(ctx context.Context) ([]Response, error)
		PaginateSchools(ctx context.Context, search string, page int, limit int, sort string) (connector.ListResponse[Response], error)
		CreateSchool(ctx context.Context, request any) (Response, error)
		UpdateSchool(ctx context.Context, request any) (Response, error)
		InativeSchool(ctx context.Context, schoolID string) (Response, error)
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

func (c client) FindByClusterId(ctx context.Context, clusterId int64) ([]int64, error) {
	var list []int64
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithCredentials(ctx).
		WithResource(fmt.Sprintf("api/schools/cluster/%d/ids", clusterId)).
		WithMethod("GET").
		Build()
	return list, c.mapper.Ids(parameter, &list)
}

func (c client) Select(ctx context.Context) ([]Response, error) {
	var schools []Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource("api/schools/select").
		WithMethod("GET").
		WithCredentials(ctx).
		Build()
	return schools, c.mapper.List(parameter, &schools)
}

func (c client) PaginateSchools(ctx context.Context, search string, page int, limit int, sort string) (connector.ListResponse[Response], error) {
	var list connector.ListResponse[Response]
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource(fmt.Sprintf("api/schools?search=%s&page=%d&page=%d&page=%s", search, page, limit, sort)).
		WithMethod("GET").
		WithCredentials(ctx).
		Build()
	return list, c.mapper.Page(parameter, &list)
}

func (c client) CreateSchool(ctx context.Context, request any) (Response, error) {
	var school Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource("api/schools").
		WithBody(request).
		WithCredentials(ctx).
		WithMethod("POST").
		Build()
	return school, c.mapper.Create(parameter, &school)
}

func (c client) UpdateSchool(ctx context.Context, request any) (Response, error) {
	var school Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource("api/schools").
		WithBody(request).
		WithCredentials(ctx).
		WithMethod("PUT").
		Build()
	return school, c.mapper.Update(parameter, &school)
}

func (c client) InativeSchool(ctx context.Context, schoolID string) (Response, error) {
	var school Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithCredentials(ctx).
		WithResource(fmt.Sprintf("api/schools/%s", schoolID)).
		WithMethod("DELETE").
		Build()
	return school, c.mapper.Inative(parameter, &school)
}
