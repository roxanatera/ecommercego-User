package main

import (
	"context"
	"e-commerce-go/awsgo"
	"e-commerce-go/models"
	"e-commerce-go/dbconfig"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(CodigoLambda)
}

func CodigoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InitAWS()

	if !ValidarParametros() {
		fmt.Println("Faltan parametros, debe enviar SecretName")
		err := errors.New("Error: Faltan parametros debe enviar SecretName")
		return event, err
	}
	var datos models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Println("Email: " + datos.UserEmail)
		case "sub":
			datos.UserUUID = att
			fmt.Println("UUID: " + datos.UserUUID)
		}
	}

	err := dbconfig.ReadSecret()
	if err != nil {
		fmt.Println("Error al leer el secreto" + err.Error())
		return event, err
	}

	return event, nil 
}

func ValidarParametros() bool {
    secretName := os.Getenv("SecretName")
    fmt.Printf("SecretName value: '%s'\n", secretName) 
    return secretName != ""
}