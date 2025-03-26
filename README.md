# 🛍️ E-Commerce Backend con AWS Cognito y Lambda (Go)

## 📌 Descripción del Proyecto
Sistema de autenticación para e-commerce que registra usuarios en **AWS Cognito** y almacena sus datos en **MySQL** mediante una función **AWS Lambda** escrita en Go.

---

## 🔧 **Tecnologías Utilizadas**
### **Cloud Services (AWS)**
- **AWS Lambda**: Ejecución del código Go en entorno serverless
- **Amazon RDS**: Instancia MySQL para almacenamiento persistente
- **Secrets Manager**: Gestión segura de credenciales
- **IAM**: Permisos y políticas de acceso
- **CloudWatch**: Monitoreo de logs

### **Lenguajes y Librerías**
- **Golang** (v1.21+)
  - `aws-sdk-go-v2`: SDK oficial de AWS
  - `database/sql`: Conexiones SQL estándar
  - `github.com/go-sql-driver/mysql`: Driver MySQL para Go

---

## 🛠️ **Implementación Técnica**
### **Estructura del Código**
```bash
.
├── main.go          # Handler Lambda
├── dbconfig/
│   ├── db.go        # Lógica de conexión a RDS
├── secretmanager/
│   ├── secretm.go   # Gestión de Secrets Manager
└── models/
    └── models.go    # Estructuras de datos