// Use this code snippet in your app.
// If you need more information about configurations or implementing the sample code, visit the AWS docs:
// https://aws.github.io/aws-sdk-go-v2/docs/getting-started/
package helper

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/sirupsen/logrus"
)

type DatabaseSecret struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetDatabaseSecrets() (string, string) {
	db := retrieveSecretFromSecretManager("rds!db-4b3599fa-95a6-4572-8bef-3dc62362c541")
	return db.Username, db.Password
}

func GetSecrets(secretName string) string {
	region := "ap-southeast-1"

	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatal(err)
	}

	// Create Secrets Manager client
	svc := secretsmanager.NewFromConfig(config)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		// For a list of exceptions thrown, see
		// https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetSecretValue.html
		logrus.WithField("err", err.Error()).Error("error gettintg secret")
		log.Fatal(err.Error())
	}

	// Decrypts secret using the associated KMS key.
	var secretString string = *result.SecretString
	return secretString
}

func retrieveSecretFromSecretManager(key string) DatabaseSecret {
	var database DatabaseSecret
	secretString := GetSecrets(key)
	if secretString == "" {
		secretString = os.Getenv(key)
	}
	if err := json.Unmarshal([]byte(secretString), &database); err != nil {
		panic("failed to get secret")
	}
	return database
}
