package school_alert_contact

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/tecmise/connector-lib/pkg/adapters/outbound/lambda"
	"github.com/tecmise/connector-lib/pkg/adapters/outbound/rest"
	"github.com/tecmise/connector-lib/pkg/ports/output/connector"
)

type (
	Client interface {
		PaginateAlertContact(ctx context.Context, search string, page int, limit int, sort string) (connector.ListResponse[Response], error)
		CreateAlertContact(ctx context.Context, request any) (Response, error)
		UpdateAlertContact(ctx context.Context, request any) (Response, error)
		InativeAlertContact(ctx context.Context, contactID string) (Response, error)
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

func (c client) PaginateAlertContact(ctx context.Context, search string, page int, limit int, sort string) (connector.ListResponse[Response], error) {
	var list connector.ListResponse[Response]
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithCredentials(ctx).
		WithResource(fmt.Sprintf("api/alertContacts?search=%s&page=%d&page=%d&page=%s", search, page, limit, sort)).
		WithMethod("GET").
		Build()
	return list, c.mapper.Page(parameter, &list)
}

func (c client) CreateAlertContact(ctx context.Context, request any) (Response, error) {
	logrus.Debug("CreateAlertContact")
	var alertContact Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithResource("api/alertContacts").
		WithCredentials(ctx).
		WithBody(request).
		WithMethod("POST").
		Build()
	return alertContact, c.mapper.Create(parameter, &alertContact)
}

func (c client) UpdateAlertContact(ctx context.Context, request any) (Response, error) {
	logrus.Debug("CreateAlertContact")
	var alertContact Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithCredentials(ctx).
		WithResource("api/alertContacts").
		WithBody(request).
		WithMethod("PUT").
		Build()
	return alertContact, c.mapper.Update(parameter, &alertContact)
}

func (c client) InativeAlertContact(ctx context.Context, contactID string) (Response, error) {
	var alertContact Response
	parameter := connector.NewParameterBuilder().
		WithHost(c.host).
		WithCredentials(ctx).
		WithResource(fmt.Sprintf("api/alertContacts/%s", contactID)).
		WithMethod("DELETE").
		Build()
	return alertContact, c.mapper.Inative(parameter, &alertContact)
}
