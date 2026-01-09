package zip_file

import (
	"context"
	"fmt"

	lambda "github.com/tecmise/connector-lib/pkg/adapters/outbound/client_lambda_proxy"
	rest "github.com/tecmise/connector-lib/pkg/adapters/outbound/client_rest"
	"github.com/tecmise/connector-lib/pkg/ports/output/connector"
)

type (
	Client interface {
		FindAll(ctx context.Context) (connector.ListResponse[Response], error)
		FindByID(ctx context.Context, zipFileID string) (Response, error)
		Create(ctx context.Context, request any) (Response, error)
		Update(ctx context.Context, zipFileID string, request any) (Response, error)
		Delete(ctx context.Context, zipFileID string) (Response, error)
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
		WithResource("api/zipFile").
		WithMethod("GET").
		WithCredentials(ctx).
		Build()
	return list, c.mapper.Page(parameter, &list)
}

func (c client) FindByID(ctx context.Context, zipFileID string) (Response, error) {
	var zipFile Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource(fmt.Sprintf("api/zipFile/%s", zipFileID)).
		WithCredentials(ctx).
		WithMethod("GET").
		Build()
	return zipFile, c.mapper.Find(parameter, &zipFile)
}

func (c client) Create(ctx context.Context, request any) (Response, error) {
	var zipFile Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource("api/zipFile").
		WithBody(request).
		WithCredentials(ctx).
		WithMethod("POST").
		WithHeader("Content-Type", "multipart/form-data").
		Build()
	return zipFile, c.mapper.Create(parameter, &zipFile)
}

func (c client) Update(ctx context.Context, zipFileID string, request any) (Response, error) {
	var zipFile Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource(fmt.Sprintf("api/zipFile/%s", zipFileID)).
		WithBody(request).
		WithCredentials(ctx).
		WithMethod("PUT").
		Build()
	return zipFile, c.mapper.Update(parameter, &zipFile)
}

func (c client) Delete(ctx context.Context, zipFileID string) (Response, error) {
	var zipFile Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithCredentials(ctx).
		WithResource(fmt.Sprintf("api/zipFile/%s", zipFileID)).
		WithMethod("DELETE").
		Build()
	return zipFile, c.mapper.Inative(parameter, &zipFile)
}
