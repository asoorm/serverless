package main

import (
	"log"

	"github.com/asoorm/serverless/provider"
	"github.com/asoorm/serverless/provider/aws"
	_ "github.com/asoorm/serverless/provider/azure"
)

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
}
