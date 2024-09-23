
# Sistema de Administración de Préstamos

Este proyecto es un sistema de administración de préstamos que permite a los clientes solicitar préstamos y a los administradores gestionar solicitudes, aprobar o rechazar préstamos y ajustar el puntaje de crédito de los clientes. Tambi+en proporciona gestión para marcar morosos préstamos caducados del lado de un administrador e integraciones externas (APIs).

---

## Tabla de Contenidos

- [Arquitectura](#arquitectura)
- [Requerimientos](#requerimientos)
- [Cómo ejecutar la aplicación](#cómo-ejecutar-la-aplicación)
- [Docker](#docker)
  - [Construir y ejecutar con Docker](#construir-y-ejecutar-con-docker)
  - [Usar Docker Compose](#usar-docker-compose)
- [Endpoints](#endpoints)
 - [POST /clients](#post-clients)
  - [POST /clients/login](#post-clientslogin)
  - [PATCH /clients/{clientID}/credit-score](#patch-clientsclientidcredit-score)
  - [POST /loans](#post-loans)
  - [GET /loans/history](#get-loanshistory)
  - [PATCH /loans/{loanID}/approve](#patch-loansloanidapprove)
  - [PATCH /loans/{loanID}/reject](#patch-loansloanidreject)
  - [POST /loans/payment](#post-loanspayment)
  - [GET /loans/active](#get-loansactive)
  - [PATCH /loans/{loanID}/delinquent](#patch-loansloaniddelincuent)
  - [PATCH /loans/delinquent/all](#patch-loansdelinquentall)
  - [POST /admins](#post-admins)
  - [POST /admins/login](#post-adminslogin)
  - [GET /admins/{adminID}](#get-adminsadminid)
  - [GET /clients](#get-clients)
  - [GET /clients/{clientID}](#get-clientsclientid)
  - [GET /api/clients/{clientID}/validate-documents](#get-apiclientsclientidvalidate-documents)
  - [GET /api/clients/{clientID}/credit-score](#get-apiclientsclientidcredit-score)
- [Pruebas](#pruebas)

---

## Arquitectura

El proyecto sigue el patrón **Clean Architecture** con principios **SOLID** para una mayor mantenibilidad y escalabilidad. La estructura del proyecto es la siguiente:
```
.
├── cmd                     # Punto de entrada de la aplicación
│   └── http
│       └── main.go         # Inicia el servidor HTTP
├── internal
│   ├── adapter             # Adaptadores (Handlers y Routers)
│   │   ├── handler         # Definición de Handlers para los endpoints
│   │   └── router          # Configuración de rutas, usa Fiber
│   ├── core                # Lógica de negocio (Dominios, servicios y puertos)
│   │   ├── domain          # Definición de estructuras de Go (Loan, Client, Admin, Apis Externas)
│   │   ├── port            # Definición de interfaces
│   │   └── service         # Implementaciones de la lógica de negocio
│   └── util                # Utilidades y helpers
├── Dockerfile              # Dockerfile
├── docker-compose.yml      # Configuración de Docker Compose
├── go.mod                  # Dependencias del proyecto
└── go.sum                  # Hash de dependencias
```

---

## Requerimientos

- **Go** versión 1.22.4
- **Docker** para la creación de contenedores
- **Docker Compose** para la orquestación de contenedores

---

## Cómo ejecutar la aplicación (local sin Docker)

1. **Compilar el proyecto**:

   ```bash
   go build -o loan-management-system ./cmd/http
   ```

2. **Ejecutar la aplicación**:

   ```bash
   ./loan-management-system
   ```

   o directamente:

   ```bash
   go run ./cmd/http
   ```
  
La aplicación estará disponible en http://localhost:9002 (Se puede modificar el puerto en el archivo `main.go`).

---

## Docker

1. **Construir la imagen de Docker:**
   ```bash
   docker build -t loan-management-system .
   ```
2. **Ejecutar la imagen de Docker:**
   ```bash
   docker run -p 9002:9002 loan-management-system
   ```

## Usar Docker Compose

1. **Levantar el contenedor con Docker Compose:**
   ```bash
   docker-compose up -d
   ```
2. **Detener el contenedor:**
   ```bash
   docker-compose stop
   ```

---

## Endpoints

### POST /clients/register

Registra un cliente en el sistema para que pueda solicitar préstamos.

- **URL**: `/clients/register`
- **Método HTTP**: `POST`
- **Cuerpo de la solicitud**:
  ```json
  {
    "fullName": "Marlon Muete",
    "email": "marlonmuete@gmail.com",
    "password": "12345"
  }
  ```
- **Respuesta exitosa:**
  `Código: 201 Created`
- **Cuerpo de respuesta**:
  ```json
  {
    "ID": "c9d32f32-1178-4dde-90d8-48077bf6c83d",
    "FullName": "Marlon Muete",
    "Email": "marlonmuete@gmail.com"
  }
  ```

### POST /clients/login

Inicia sesión un cliente en el sistema y genera un token JWT para autenticación.

- **URL**: `/clients/login`
- **Método HTTP**: `POST`
- **Cuerpo de la solicitud**:
  ```json
  {
    "id": "c9d32f32-1178-4dde-90d8-48077bf6c83d",
    "password": "12345"
  }
  ```
- **Respuesta exitosa:**
  `Código: 200 OK`
- **Cuerpo de respuesta**:
  ```json
  {
    "id": "c9d32f32-1178-4dde-90d8-48077bf6c83d",
    "token": "jwt-token"
  }
  ```

### PATCH /clients/{clientID}/credit-score

Actualiza el puntaje de crédito de un cliente.

- **URL**: `/clients/{clientID}/credit-score`
- **Método HTTP**: `PATCH`
- **Cuerpo de la solicitud**:
  ```json
  {
    "creditScore": 700
  }
  ```
- **Respuesta exitosa:**
  `Código: 200 OK`
- **Cuerpo de respuesta**:
  ```json
  {
    "ID": "c9d32f32-1178-4dde-90d8-48077bf6c83d",
    "FullName": "Marlon Muete",
    "CreditScore": 700
  }
  ```

### POST /loans

Permite a un cliente solicitar un préstamo.

- **URL**: `/loans`
- **Método HTTP**: `POST`
- **Cuerpo de la solicitud**:
  ```json
  {
    "clientId": "c9d32f32-1178-4dde-90d8-48077bf6c83d",
    "amount": 5000,
    "termInMonths": 12
  }
  ```
- **Respuesta exitosa:**
  `Código: 201 Created`
- **Cuerpo de respuesta**:
  ```json
  {
    "ID": "870352df-e651-4fe6-8cc8-1f97510e652d",
    "ClientID": "c9d32f32-1178-4dde-90d8-48077bf6c83d",
    "Amount": 5000,
    "TermInMonths": 12,
    "Status": "Pending",
    "RemainingAmount": 5000
  }
  ```

  #### PATCH `/loans/:id/approve`
**Description**: Se aprueba un prestamo en tanto su id.

**Response**:
```json
{
    "message": "Loan approved successfully."
}
```

#### PATCH `/loans/:id/reject`
**Description**: Se rechaza un prestamo en tanto su id.

**Response**:
```json
{
    "message": "Loan rejected successfully."
}
```

#### POST `/loans/payment`
**Description**: Se paga el prestamo completo o por partes.

**Request Body**:
```json
{
    "loan_id": "870352df-e651-4fe6-8cc8-1f97510e652d",
    "amount": 1000
}
```

**Response**:
```json
{
    "ID": "9ce5945b-08a9-44f3-9401-e832e11ac906",
    "LoanID": "870352df-e651-4fe6-8cc8-1f97510e652d",
    "Amount": 1000,
    "Date": "2024-09-22T14:38:21.612246-05:00",
    "Status": "completed"
}
```

#### PATCH `/loans/:id/delinquent`
**Description**: Se marca como moroso un prestamo.

**Response**:
```json
{
    "message": "Loan marked as delinquent."
}
```

#### PATCH `/loans/delinquent/all`
**Description**: Se marcan todos los prestamos como morosos.

**Response**:
```json
{
    "message": "All delinquent loans updated."
}
```

#### GET `/loans/active`
**Description**: Devuelve los prestamos activos de un cliente.

**Response**:
```json
{
    "ID": "870352df-e651-4fe6-8cc8-1f97510e652d",
    "ClientID": "c9d32f32-1178-4dde-90d8-48077bf6c83d",
    "Amount": 5000,
    "InterestRate": 5,
    "TermInMonths": 12,
    "Status": "Pending",
    "CreatedAt": "2024-09-22T14:36:21.55704-05:00",
    "ApprovedAt": null,
    "RejectedAt": null,
    "DueDate": "2025-09-22T14:36:21.55704-05:00",
    "RemainingAmount": 5000,
    "IsPaid": false
}
```

### Admin Endpoints

#### POST `/admins`
**Description**: Se registra un nuevo admin.

**Request Body**:
```json
{
    "fullName": "Jerson Monterroso",
    "role": "General",
    "password": "adminpass123"
}
```

**Response**:
```json
{
    "ID": "cde6eb7e-688d-4364-97b0-91c43f2efe91",
    "FullName": "Jerson Monterroso",
    "Role": "General"
}
```

#### POST `/admins/login`
**Description**: Login para un admin.

**Request Body**:
```json
{
    "id": "cde6eb7e-688d-4364-97b0-91c43f2efe91",
    "password": "adminpass123"
}
```

**Response**:
```json
{
    "token": "jwt-token"
}
```

#### GET `/admins/:id`
**Description**: Devuelve la info de un admin a razón de su id.

**Response**:
```json
{
    "ID": "cde6eb7e-688d-4364-97b0-91c43f2efe91",
    "FullName": "Jerson Monterroso",
    "Role": "General"
}
```

### GET /loans/history

Obtiene el historial de préstamos en el sistema. Solo los administradores pueden acceder a este endpoint.

- **URL**: `/loans/history`
- **Método HTTP**: `GET`
- **Respuesta exitosa**: `Código: 200 OK`

### PATCH /loans/{loanID}/approve

Aprueba un préstamo para un cliente.

- **URL**: `/loans/{loanID}/approve?adminId={adminID}`
- **Método HTTP**: `PATCH`

### PATCH /loans/{loanID}/reject

Rechaza un préstamo para un cliente.

- **URL**: `/loans/{loanID}/reject?adminId={adminID}`
- **Método HTTP**: `PATCH`

---
