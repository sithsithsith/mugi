package config

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type Config struct {
	AWSRegion       string
	CognitoUserPool string
	CognitoAppID    string
}

func LoadConfig() *Config {
	return &Config{
		AWSRegion:       getEnv("AWS_REGION", "us-east-1"),
		CognitoUserPool: getEnv("COGNITO_USER_POOL_ID", ""),
		CognitoAppID:    getEnv("COGNITO_APP_CLIENT_ID", ""),
	}
}

func GetCognitoClient() *cognitoidentityprovider.CognitoIdentityProvider {
	cfg := LoadConfig()
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(cfg.AWSRegion),
	})
	if err != nil {
		log.Fatalf("Failed to create AWS session: %v", err)
	}

	return cognitoidentityprovider.New(sess)
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
