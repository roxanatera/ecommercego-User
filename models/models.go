package models

import (

)

type SecretRDSJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Engine string `json:"engine"`
	Host string `json:"host"`
	Port string `json:"port"`
	Dbname string `json:"dbname"`
}

type SignUp struct {
	UserEmail string `json:"userEmail"`
	UserUUID string `json:"userUUID"`
}