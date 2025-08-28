package school_user_profile

import (
	"context"
	"fmt"
	"github.com/tecmise/connector-school-api/pkg/adapters/outbound/lambda"
	"github.com/tecmise/connector-school-api/pkg/adapters/outbound/rest"
	"github.com/tecmise/connector-school-api/pkg/ports/output/connector"
	"github.com/tecmise/connector-school-api/pkg/ports/output/constant"
)

type (
	Client interface {
		FindByUserId(ctx context.Context, cognitoUserId string) ([]SchoolTenantResponse, error)
	}

	client struct {
		mapper connector.Call[SchoolTenantResponse]
		host   string
	}
)

func Rest(host string) Client {
	return &client{
		host:   host,
		mapper: rest.NewConnector[SchoolTenantResponse](),
	}
}
func Lambda(functionName string) Client {
	return &client{
		host:   functionName,
		mapper: lambda.NewConnector[SchoolTenantResponse](),
	}
}

func (c client) FindByUserId(ctx context.Context, cognitoUserId string) ([]SchoolTenantResponse, error) {
	var list []SchoolTenantResponse
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithCredentials(ctx).
		WithResource(fmt.Sprintf("api/permissions/user/%s", cognitoUserId)).
		WithMethod("GET").
		WithRegion(constant.USEast1).
		Build()
	return list, c.mapper.List(parameter, &list)
}
