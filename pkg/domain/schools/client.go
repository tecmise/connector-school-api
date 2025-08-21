package schools

import (
	"context"
	"fmt"
	"github.com/tecmise/school-api-connector/pkg/adapters/outbound/lambda"
	"github.com/tecmise/school-api-connector/pkg/ports/output/connector"
)

type (
	Client interface {
		FindByClusterId(ctx context.Context, clusterId int64) ([]int64, error)
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

type client struct {
	identifier string
	responses  connector.Call[[]Response]
}

func NewClient(identifier string) Client {
	return &client{
		host:   functionName,
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
