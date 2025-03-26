package main

import (
    "context"
    "e-commerce-go/dbconfig"
    "e-commerce-go/models"
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
    if !ValidarParametros() {
        return event, errors.New("error: SecretName no configurado")
    }

    var datos models.SignUp

    // Extraer todos los atributos necesarios
    for row, att := range event.Request.UserAttributes {
        switch row {
        case "email":
            datos.UserEmail = att
        case "sub":
            datos.UserUUID = att
        case "given_name":
            datos.UserFirstName = att
        case "family_name":
            datos.UserLastName = att
        }
    }

    fmt.Printf("Datos del usuario: %+v\n", datos)

    // Conectar a la base de datos
    err := dbconfig.ReadSecret(ctx) 
    if err != nil {
        return event, fmt.Errorf("error al leer secreto: %v", err)
    }
    // Insertar en la base de datos
    err = dbconfig.SignUp(datos)
    if err != nil {
        return event, fmt.Errorf("error al registrar usuario: %v", err)
    }

    return event, nil
}

func ValidarParametros() bool {
    return os.Getenv("SecretName") != ""
}