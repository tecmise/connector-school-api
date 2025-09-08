package profiles

import (
	"context"
	"fmt"
	"github.com/tecmise/connector-lib/pkg/adapters/outbound/lambda"
	"github.com/tecmise/connector-lib/pkg/adapters/outbound/rest"
	"github.com/tecmise/connector-lib/pkg/ports/output/connector"
	"github.com/tecmise/connector-lib/pkg/ports/output/constant"
)

type (
	Client interface {
		FindRolesByProfileId(ctx context.Context, profileID int64) ([]string, error)
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

func (c client) FindRolesByProfileId(ctx context.Context, profileID int64) ([]string, error) {
	var list []string
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithCredentials(ctx).
		WithResource(fmt.Sprintf("api/roles/profile/%d", profileID)).
		WithMethod("GET").
		WithRegion(constant.USEast1).
		Build()
	return list, c.mapper.Strings(parameter, &list)
}
