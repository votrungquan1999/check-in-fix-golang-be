package setup

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

type Config struct {
	FirebaseProjectID           string `mapstructure:"FIREBASE_PROJECT_ID"`
	FirebasePrivateKeyID        string `mapstructure:"FIREBASE_PRIVATE_KEY_ID"`
	FirebasePrivateKey          string `mapstructure:"FIREBASE_PRIVATE_KEY"`
	FirebaseClientEmail         string `mapstructure:"FIREBASE_CLIENT_EMAIL"`
	FirebaseClientID            string `mapstructure:"FIREBASE_CLIENT_ID"`
	FirebaseType                string `mapstructure:"FIREBASE_TYPE"`
	FirebaseAuthURI             string `mapstructure:"FIREBASE_AUTH_URI"`
	FirebaseTokenURI            string `mapstructure:"FIREBASE_TOKEN_URI"`
	FirebaseAuthProviderCertURL string `mapstructure:"FIREBASE_AUTH_PROVIDER_CERT_URL"`
	FirebaseClientCertURL       string `mapstructure:"FIREBASE_CLIENT_CERT_URL"`
	NexmoApiKey                 string
	NexmoApiSecret              string
}

var EnvConfig Config

func stringDefault(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)

	if !exists {
		return defaultValue
	}

	return value
}

func mustString(key string) string {
	value, exists := os.LookupEnv(key)

	if !exists {
		log.Fatalf("can not load env %s", key)
	}

	return value
}

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		println("no .env file found")
	}

	firebasePrivateKey := mustString("FIREBASE_PRIVATE_KEY")
	firebasePrivateKey = strings.ReplaceAll(firebasePrivateKey, "\\n", "\n")

	EnvConfig = Config{
		FirebaseType:                mustString("FIREBASE_TYPE"),
		FirebaseProjectID:           mustString("FIREBASE_PROJECT_ID"),
		FirebasePrivateKeyID:        mustString("FIREBASE_PRIVATE_KEY_ID"),
		FirebasePrivateKey:          firebasePrivateKey,
		FirebaseClientEmail:         mustString("FIREBASE_CLIENT_EMAIL"),
		FirebaseClientID:            mustString("FIREBASE_CLIENT_ID"),
		FirebaseAuthURI:             mustString("FIREBASE_AUTH_URI"),
		FirebaseTokenURI:            mustString("FIREBASE_TOKEN_URI"),
		FirebaseAuthProviderCertURL: mustString("FIREBASE_AUTH_PROVIDER_X509_CERT_URL"),
		FirebaseClientCertURL:       mustString("FIREBASE_CLIENT_X509_CERT_URL"),
		NexmoApiKey:                 stringDefault("NEXMO_API_KEY", ""),
		NexmoApiSecret:              stringDefault("NEXMO_API_SECRETE", ""),
	}
}

func GetFirebaseConfig() map[string]string {
	var config = map[string]string{
		"type":                        EnvConfig.FirebaseType,
		"project_id":                  EnvConfig.FirebaseProjectID,
		"private_key_id":              EnvConfig.FirebasePrivateKeyID,
		"private_key":                 EnvConfig.FirebasePrivateKey,
		"client_email":                EnvConfig.FirebaseClientEmail,
		"client_id":                   EnvConfig.FirebaseClientID,
		"auth_uri":                    EnvConfig.FirebaseAuthURI,
		"token_uri":                   EnvConfig.FirebaseTokenURI,
		"auth_provider_x509_cert_url": EnvConfig.FirebaseAuthProviderCertURL,
		"client_x509_cert_url":        EnvConfig.FirebaseClientCertURL,
	}

	return config
}
