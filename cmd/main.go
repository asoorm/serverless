package main

import (
	"log"

	"github.com/asoorm/serverless/provider"
	_ "github.com/asoorm/serverless/provider/aws"
)

func main() {

	serverless, err := provider.GetProvider("aws-lambda")
	if err != nil {
		log.Fatal(err)
	}

	conf := map[string]string{
		"region": "eu-west-2",
	}

	serverless.Init(conf)

	detail, err := serverless.List()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n", detail)
}
