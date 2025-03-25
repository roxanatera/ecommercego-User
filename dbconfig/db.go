package dbconfig

import (
	"e-commerce-go/models"
	"e-commerce-go/secretmanager"
	"os"
	"fmt"
)

var SecretModel models.SecretRDSJson

func ReadSecret() error {
	secretName := os.Getenv("SecretName")
	if secretName == "" {
		return fmt.Errorf("SecretName no está definido en las variables de entorno")
	}

	var err error
	SecretModel, err = secretmanager.GetSecret(secretName)
	return err
}