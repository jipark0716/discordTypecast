package services

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

type configService struct {
	params map[string]interface{}
}

var configInstance *configService

// return config service singleton instance
func GetConfigInstance() *configService {
	if configInstance == nil {
		configInstance = new(configService)
		configInstance.init()
	}

	return configInstance
}

func (cs *configService) Get(key string) interface{} {
	return cs.params[key]
}

func (cs *configService) init() {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String("ap-northeast-2")},
		SharedConfigState: session.SharedConfigEnable,
	})

	if err != nil {
		log.Fatal(err)
	}

	ssmsvc := ssm.New(sess, aws.NewConfig())

	param, err := ssmsvc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String("/env/dev/discord/typecast"),
		WithDecryption: aws.Bool(false),
	})

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal([]byte(*param.Parameter.Value), &cs.params)

	if err != nil {
		log.Fatal(err)
	}
}
