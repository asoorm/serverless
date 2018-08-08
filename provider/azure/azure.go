package azure

import (
	"github.com/asoorm/serverless/provider"
	"github.com/kataras/go-errors"
)

const (
	//azureUrlFormat = "https://tyk-hello.azurewebsites.net/api/HttpTriggerJS1?code=mZT2aajwLKmWb64URhNgJr5LuCyVah66loou8nDBJj4qF8B30FlMrg=="
	azureUrlFormat = "https://%s.azurewebsites.net/api/%s?code=%s"
)

type Conf struct {
	AppName      string
	FunctionName string
	AuthCode     string
}

func init() {
	provider.RegisterProvider("azure-functions", NewProvider)
}

func NewProvider() (provider.Provider, error) {
	return &Provider{}, nil
}

type Provider struct {
	Conf
}

func (p *Provider) Init(conf provider.Conf) error {

	c, ok := conf.(Conf)
	if !ok {
		return errors.New("unable to resolve conf type")
	}
	p.AppName = c.AppName
	p.AuthCode = c.AuthCode
	p.FunctionName = c.FunctionName

	return nil
}

func (Provider) List() ([]provider.Function, error) {

	return nil, errors.New("azure doesnt allow listing functions")
}

// TODO
func (p Provider) Invoke(function provider.Function, requestBody []byte) (*provider.Response, error) {
	return nil, nil
}
