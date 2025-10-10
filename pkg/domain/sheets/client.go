package sheets

import (
	"context"
	"fmt"

	"github.com/tecmise/connector-lib/pkg/adapters/outbound/lambda"
	"github.com/tecmise/connector-lib/pkg/adapters/outbound/rest"
	"github.com/tecmise/connector-lib/pkg/ports/output/connector"
)

type (
	Client interface {
		FindAll(ctx context.Context) (connector.ListResponse[Response], error)
		FindByID(ctx context.Context, sheetID string) (Response, error)
		Create(ctx context.Context, request any) (Response, error)
		Update(ctx context.Context, sheetID string, request any) (Response, error)
		Delete(ctx context.Context, sheetID string) (Response, error)
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

func (c client) FindAll(ctx context.Context) (connector.ListResponse[Response], error) {
	var list connector.ListResponse[Response]
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource("api/sheets").
		WithMethod("GET").
		WithCredentials(ctx).
		Build()
	return list, c.mapper.Page(parameter, &list)
}

func (c client) FindByID(ctx context.Context, sheetID string) (Response, error) {
	var sheet Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource(fmt.Sprintf("api/sheets/%s", sheetID)).
		WithCredentials(ctx).
		WithMethod("GET").
		Build()
	return sheet, c.mapper.Find(parameter, &sheet)
}

func (c client) Create(ctx context.Context, request any) (Response, error) {
	var sheet Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource("api/sheets").
		WithBody(request).
		WithCredentials(ctx).
		WithMethod("POST").
		WithHeader("Content-Type", "multipart/form-data").
		Build()
	return sheet, c.mapper.Create(parameter, &sheet)
}

func (c client) Update(ctx context.Context, sheetID string, request any) (Response, error) {
	var sheet Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource(fmt.Sprintf("api/sheets/%s", sheetID)).
		WithBody(request).
		WithCredentials(ctx).
		WithMethod("PUT").
		Build()
	return sheet, c.mapper.Update(parameter, &sheet)
}

func (c client) Delete(ctx context.Context, sheetID string) (Response, error) {
	var sheet Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithCredentials(ctx).
		WithResource(fmt.Sprintf("api/sheets/%s", sheetID)).
		WithMethod("DELETE").
		Build()
	return sheet, c.mapper.Inative(parameter, &sheet)
}
