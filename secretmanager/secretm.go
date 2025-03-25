package secretmanager

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"e-commerce-go/awsgo"
	"e-commerce-go/models"

)

//Funcion que obtiene el secreto de AWS Secret Manager

func GetSecret(nameSecret string) (models.SecretRDSJson, error) { 
	var datosSecret models.SecretRDSJson
	fmt.Println(" > GetSecretv" + nameSecret)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nameSecret),
	})
	if err != nil {
		fmt.Println("Error al obtener el secreto")
		return datosSecret, err
	}
	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println(" > Secreto obtenido" + nameSecret)
	return datosSecret, nil


}