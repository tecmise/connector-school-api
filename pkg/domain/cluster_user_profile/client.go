package cluster_user_profile

import (
	"context"
	"fmt"
	"github.com/tecmise/school-api-connector/pkg/adapters/outbound/lambda"
	"github.com/tecmise/school-api-connector/pkg/adapters/outbound/rest"
	"github.com/tecmise/school-api-connector/pkg/ports/output/connector"
	"github.com/tecmise/school-api-connector/pkg/ports/output/constant"
)

type (
	Client interface {
		FindByUserId(ctx context.Context, cognitoUserId string) ([]Response, error)
		FindSchoolsIdsByClusterId(ctx context.Context, clusterId int64) ([]int64, error)
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
func Lambda(functionName string) Client {
	return &client{
		host:   functionName,
		mapper: lambda.NewConnector[Response](),
	}
}

func (c client) FindByUserId(_ context.Context, cognitoUserId string) ([]Response, error) {
	var list []Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource(fmt.Sprintf("api/permissions/user/%s/clusters", cognitoUserId)).
		WithMethod("GET").
		WithRegion(constant.USEast1).
		Build()
	return list, c.mapper.List(parameter, &list)
}

func (c client) FindSchoolsIdsByClusterId(_ context.Context, clusterId int64) ([]int64, error) {
	var list []int64
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource(fmt.Sprintf("api/schools/cluster/%d/ids", clusterId)).
		WithMethod("GET").
		WithRegion(constant.USEast1).
		Build()
	return list, c.mapper.Ids(parameter, &list)
}
