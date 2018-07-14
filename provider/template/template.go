package template

import (
	"github.com/asoorm/serverless/provider"
)

func init() {
	provider.RegisterProvider("tempate-function", NewProvider)
}

func NewProvider() (provider.Provider, error) {
	return &Provider{}, nil
}

type Provider struct{}

func (p *Provider) Init(conf provider.Conf) error {

	return nil
}

func (Provider) List() ([]provider.Function, error) {

	return nil, nil
}

func (p Provider) Invoke(function provider.Function, requestBody []byte) (*provider.Response, error) {

	return nil, nil
}
