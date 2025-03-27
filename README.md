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
- **API Gateway.**:

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


### Características clave:
1. **Formato Markdown optimizado** para GitHub
2. **Tablas claras** para tecnologías y versiones
3. **Diagramas Mermaid** para flujos de trabajo
4. **Secciones de solución de problemas**
5. **Guía de contribución** estándar
6. **Estructura de proyecto profesional**

### Mejoras sugeridas:
- Añadir badges de CI/CD y coverage
- Incluir ejemplos de requests/responses
- Añadir capturas de pantalla de la configuración AWS

## 🛠 Requisitos Previos
- Cuenta de AWS con acceso a API Gateway.
- Conocimiento básico de rutas HTTP y métodos (PUT, GET, etc.).
- AWS CLI instalada (opcional, para despliegue desde terminal).

---


### 1. Crear el Recurso `{folder}`
1. En la consola de API Gateway, ve a tu API.
2. En el panel de recursos, haz clic en **Crear recurso**.
3. Configura:
   - **Nombre del recurso**: `folder`.
   - **Ruta del recurso**: `{folder}` (con llaves `{}`).
4. Haz clic en **Crear recurso**.

### 2. Crear el Recurso Hijo `{object}`
1. Selecciona el recurso `/{folder}`.
2. Haz clic en **Crear recurso** nuevamente.
3. Configura:
   - **Nombre del recurso**: `object`.
   - **Ruta del recurso**: `{object}`.
4. Haz clic en **Crear recurso**.

---

## 🔄 Agregar Métodos HTTP (Ej: PUT)
1. Selecciona el recurso `/{folder}/{object}`.
2. En el panel de acciones, elige **Crear método**.
3. Selecciona `PUT` y confirma con el icono de check (✓).
4. Configura la integración (Lambda, HTTP, etc.).
