package main

import (
	"log"

	"encoding/json"

	"github.com/asoorm/serverless/provider"
	"github.com/asoorm/serverless/provider/aws"
	_ "github.com/asoorm/serverless/provider/azure"
)

type Event struct {
	Username string
}

func main() {

	lambda, err := provider.GetProvider("aws-lambda")
	if err != nil {
		log.Fatal(err)
	}

	conf := aws.Conf{
		Region: "eu-west-2",
	}

	if err := lambda.Init(conf); err != nil {
		log.Fatal(err)
	}

	detail, err := lambda.List()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", detail)

	e, err := json.Marshal(Event{
		Username: "michael",
	})
	if err != nil {
		log.Fatal(err)
	}

	res, err := lambda.Invoke(provider.Function{
		Name: "HelloWorldGoFunctions",
	}, e)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("body: %s\n", res.Body)
	log.Printf("status: %d\n", res.StatusCode)
	log.Printf("error: %s\n", res.Error)
}
