package third_party

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/sirupsen/logrus"
)

type SecretManager struct {
	SecretName string
	Region     string
}

func (s *SecretManager) GetSecrets(value any) error {
	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(s.Region))
	if err != nil {
		log.Fatal(err)
	}

	// Create Secrets Manager client
	svc := secretsmanager.NewFromConfig(config)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(s.SecretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		// For a list of exceptions thrown, see
		// https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetSecretValue.html
		logrus.WithField("err", err.Error()).Error("error gettintg secret")
		return err
	}

	// Decrypts secret using the associated KMS key.
	var secretString string = *result.SecretString

	// unmarshal string into interface
	if err := json.Unmarshal([]byte(secretString), &value); err != nil {
		logrus.WithField("error", err).Error("failed to unmarshal secret")
		return err
	}
	return nil
}
