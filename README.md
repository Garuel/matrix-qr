# 🧮 Matrix Processor & Statistics API

![Go](https://img.shields.io/badge/Go-1.26-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Node.js](https://img.shields.io/badge/Node.js-20+-339933?style=for-the-badge&logo=nodedotjs&logoColor=white)
![TypeScript](https://img.shields.io/badge/TypeScript-007ACC?style=for-the-badge&logo=typescript&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)

Un reto técnico full-stack compuesto por dos microservicios interconectados y una interfaz web sencilla.

El sistema toma una matriz matemática, realiza una **Factorización QR**, y luego calcula **estadísticas** sobre las matrices resultantes (suma, promedio, comprobación de matriz diagonal, etc.).

---

## 🏗️ Arquitectura del Sistema

1. **Frontend (Cliente Web):** Una interfaz sencilla en HTML/JS que permite al usuario ingresar una matriz y enviarla a procesar.
2. **Go API (Core Matemático - Puerto `3000`):** Recibe la matriz del frontend, realiza la **factorización QR** utilizando la librería matemática `gonum`, y envía los resultados a la API de Node para su análisis.
3. **Node.js API (Calculadora de Estadísticas - Puerto `7787`):** Recibe las matrices (Q y R), valida los datos enviados con `class-validator`, y calcula las estadísticas, estas son retornadas al servicio de Go para la respuesta final.

---

## 🚀 Características Principales

- **Factorización QR (Go):** Procesamiento matemático eficiente gracias a `gonum`.
- **Validación (Node.js):** Uso de DTOs y decoradores de `class-validator` para garantizar la integridad de los datos.
- **CORS Habilitado:** Listo para ser consumido directamente desde la aplicación de frontend.
- **Dockerizado:** Entornos preconfigurados con `docker-compose` y `Dockerfile`.

---

## 🛠️ Tecnologías y Herramientas

| Componente     | Tecnología Principal | Librerías Destacadas                       |
| :------------- | :------------------- | :----------------------------------------- |
| **Api 1**      | Go (Golang)          | Fiber (Web Framework), Gonum (Matemáticas) |
| **Api 2**      | Node.js + TypeScript | Express, Class-Validator, Jest (Testing)   |
| **Frontend**   | Vanilla JS / HTML    | Fetch API                                  |
| **Despliegue** | Docker               | Docker Compose, Alpine Images              |

---

## 📦 Instalación y Despliegue (Docker)

La forma más sencilla de levantar todo el proyecto es utilizando Docker.

### Requisitos previos

- [Docker](https://www.docker.com/get-started) y Docker Compose instalados.

### Pasos

1. Clona el repositorio y ve a la carpeta raíz:
   ```bash
   git clone <url-del-repositorio>
   cd reto-tecnico-interseguro
   ```
2. Ejecuta Docker Compose:

   ```bash
   docker-compose up --build
   ```

   > 💡 _Esto levantará de forma orquestada ambas apis: `go-api-container` en el puerto 3000 y `stats-node-container` en el puerto 7787._

3. Abre el archivo `frontend/index.html` en tu navegador para probar la interacción.

---

## 💻 Ejecución Local (Sin Docker)

Para correr los servicios localmente y sin contenedores:

### 1. Node.js API (Estadísticas)

```bash
cd calculador-estadistica-node-api
npm install
npm run dev
```

_(El servidor correrá en `http://localhost:7787`)_

### 2. Go API (Procesador)

```bash
cd matriz-procesador-go-api
go mod tidy
go run cmd/api/main.go
```

_(El servidor correrá en `http://localhost:3000`)_

---

## 📡 Endpoint Principal

### 🟢 `POST /api/matrix/process` (Go API)

Recibe una matriz nxn, la factoriza en Q y R, y devuelve las matrices junto a las estadísticas calculadas en Node.

**Request:**

```json
{
  "matrix": [
    [1, 2],
    [3, 4]
  ]
}
```

**Response (Ejemplo resumido):**

```json
{
  "matrixQ": [
    [-0.3162277660168379, -0.9486832980505138],
    [-0.9486832980505138, 0.3162277660168379]
  ],
  "matrixR": [
    [-3.162277660168379, -4.427188724235731],
    [0, -0.6324555320336758]
  ],
  "stats": {
    "sum": -9.168598418047976,
    "average": -1.146074802255997,
    "isDiagonal": false
  }
}
```

---

## 🧪 Pruebas Unitarias (Testing)

Para ejecutar las pruebas unitarias de cada API.

- **Para Node.js (Jest):**
  ```bash
  cd calculador-estadistica-node-api
  npm run test
  ```
- **Para Go (Testing nativo):**
  ```bash
  cd matriz-procesador-go-api
  go test ./...
  ```
