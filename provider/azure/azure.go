package azure

import (
	"github.com/asoorm/serverless/provider"
	"github.com/kataras/go-errors"
)

func init() {
	provider.RegisterProvider("azure-functions", NewAzureProvider)
}

func NewAzureProvider() (provider.Provider, error) {
	return &Provider{}, nil
}

type Provider struct{}

func (p *Provider) Init(conf provider.Conf) error {

	return nil
}

func (Provider) List() ([]provider.Function, error) {

	return nil, errors.New("azure doesnt allow listing functions")
}

func (p Provider) Invoke(function provider.Function, requestBody []byte) (*provider.Response, error) {

	return nil, nil
}
