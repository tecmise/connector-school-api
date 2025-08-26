package connector

import (
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/tecmise/connector-school-api/pkg/ports/output/constant"
)

type Parameter struct {
	Host       string
	Resource   string
	Method     string
	Body       any
	Region     constant.AWSRegion
	UserID     uuid.UUID
	UserPoolID string
	Headers    map[string]string
}

type ParameterBuilder struct {
	param Parameter
}

func (b *Parameter) GetHttpURL() string {
	return fmt.Sprintf("%s/%s", b.Host, b.Resource)
}

func NewParameterBuilder() *ParameterBuilder {
	return &ParameterBuilder{param: Parameter{}}
}

func (b *ParameterBuilder) WithHost(host string) *ParameterBuilder {
	b.param.Host = host
	return b
}

func (b *ParameterBuilder) WithResource(resource string) *ParameterBuilder {
	b.param.Resource = resource
	return b
}

func (b *ParameterBuilder) WithMethod(method string) *ParameterBuilder {
	b.param.Method = method
	return b
}

func (b *ParameterBuilder) WithBody(body any) *ParameterBuilder {
	b.param.Body = body
	return b
}

func (b *ParameterBuilder) WithRegion(region constant.AWSRegion) *ParameterBuilder {
	b.param.Region = region
	return b
}

func (b *ParameterBuilder) WithHeader(key, value string) *ParameterBuilder {
	if b.param.Headers == nil {
		b.param.Headers = make(map[string]string)
	}
	b.param.Headers[key] = value
	return b
}

func (b *ParameterBuilder) Build() Parameter {
	if b.param.Region == "" {
		b.param.Region = constant.USEast1
	}
	return b.param
}
