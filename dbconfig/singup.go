package dbconfig

import (
	"e-commerce-go/models"
	"e-commerce-go/tools"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("SignUp")

	err := DbConnect()  
	if err != nil {
		return fmt.Errorf("error conectando a DB: %v", err)
	}
	defer Db.Close()

	
	sentencia := "INSERT INTO users (user_Email, user_UUID, user_DateAdd) VALUES ('" + sig.UserEmail + "', '" + sig.UserUUID + "', '" + tools.FechaMySQL() + "')"
	fmt.Println(sentencia)
	_, err = Db.Exec(sentencia)
	if err != nil {
		return fmt.Errorf("error insertando usuario: %v", err)
	}

	fmt.Println("Registro insertado correctamente")
	return nil
}