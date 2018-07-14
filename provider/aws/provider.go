package aws

import (
	"encoding/json"

	"github.com/Sirupsen/logrus"
	"github.com/asoorm/serverless/provider"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/pkg/errors"
)

func init() {
	provider.RegisterProvider("aws-lambda", NewProvider)
}

func NewProvider() (provider.Provider, error) {
	return &Provider{}, nil
}

type Provider struct {
	Region string
	aws.Config
}

func (p *Provider) Init(conf provider.Conf) error {
	p.Region = conf["region"]

	awsCfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return errors.Wrap(err, provider.ErrorLoadingDriverConfig)
	}

	p.Config = awsCfg

	return nil
}

func (p Provider) List() ([]provider.Function, error) {

	service := p.getService()

	listFunctionsRequest := service.ListFunctionsRequest(nil)
	listFunctionsOutput, err := listFunctionsRequest.Send()

	if err != nil {
		return nil, errors.Wrap(err, provider.ErrorListingFunctions)
	}

	lfoJs, _ := json.MarshalIndent(listFunctionsOutput.Functions, "aws:", "  ")
	logrus.Info(string(lfoJs))

	functions := make([]provider.Function, 0)
	for _, f := range listFunctionsOutput.Functions {

		myFunc := provider.Function{
			Name:    aws.StringValue(f.FunctionName),
			Version: aws.StringValue(f.Version),
		}

		functions = append(functions, myFunc)
	}

	fJs, _ := json.Marshal(functions)
	logrus.Info(string(fJs))

	return functions, nil
}

func (p Provider) Invoke(function provider.Function, requestBody []byte) (*provider.Response, error) {

	service := p.getService()

	if function.GetVersion() == "" {
		function.SetVersion("$LATEST")
	}

	input := lambda.InvokeInput{
		//ClientContext:  aws.String("index"), // need to investigate context
		FunctionName: aws.String(function.GetName()),
		//InvocationType: lambda.InvocationTypeEvent, // To make async
		LogType:   lambda.LogTypeTail,
		Payload:   requestBody,
		Qualifier: aws.String(function.GetName()),
	}

	request := service.InvokeRequest(&input)

	lambdaRes, err := request.Send()
	if err != nil {
		return nil, errors.Wrap(err, provider.ErrorInvokingFunction)
	}

	res := provider.Response{
		StatusCode: int(aws.Int64Value(lambdaRes.StatusCode)),
		Body:       lambdaRes.Payload,
	}

	if lambdaRes.FunctionError != nil {
		res.Error = errors.New(aws.StringValue(lambdaRes.FunctionError))
	}

	return &res, nil
}

func (p Provider) getService() *lambda.Lambda {
	cfg := p.Config
	cfg.Region = p.Region

	return lambda.New(cfg)
}
