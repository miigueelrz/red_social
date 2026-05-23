# Despliegue y ejecucion

## 1. Objetivo del documento

Este documento describe la estrategia de ejecucion local, configuracion, integracion continua y despliegue de la red social desarrollada con Go, htmx y templ.

El objetivo es que la aplicacion pueda ejecutarse de forma reproducible, que el entorno este documentado y que el proyecto pueda entregarse o desplegarse con claridad.

## 2. Entornos previstos

Se contemplan tres entornos:

- Desarrollo local.
- Pruebas/validacion.
- Produccion o despliegue final.

### Desarrollo local

Entorno utilizado por el equipo durante la implementacion.

Caracteristicas:

- Go instalado localmente.
- templ instalado para generar componentes.
- Base de datos local o mediante Docker.
- Recarga manual o herramienta de desarrollo si se incorpora.

### Pruebas/validacion

Entorno usado para ejecutar tests y comprobar que la aplicacion compila.

Caracteristicas:

- Ejecucion automatica mediante CI.
- Base de datos de test si se requiere.
- Validacion de build.

### Produccion/despliegue final

Entorno donde la aplicacion estara disponible para demostracion.

Caracteristicas:

- Variables de entorno configuradas.
- Base de datos persistente.
- Binario Go compilado.
- Archivos estaticos disponibles.
- Migraciones aplicadas.

## 3. Variables de entorno

La aplicacion se configurara mediante variables de entorno para separar configuracion y codigo.

Variables previstas:

```text
APP_ENV=development
APP_PORT=8080
DATABASE_URL=postgres://user:password@localhost:5432/social_app?sslmode=disable
SESSION_SECRET=change-me
CSRF_SECRET=change-me
```

### Descripcion

| Variable | Descripcion |
| --- | --- |
| `APP_ENV` | Entorno de ejecucion: development, test o production |
| `APP_PORT` | Puerto en el que se ejecuta el servidor |
| `DATABASE_URL` | Cadena de conexion con la base de datos |
| `SESSION_SECRET` | Secreto usado para sesiones |
| `CSRF_SECRET` | Secreto usado para proteccion CSRF si se implementa |

## 4. Ejecucion local sin Docker

Pasos previstos:

```text
1. Clonar el repositorio.
2. Instalar Go.
3. Instalar templ.
4. Configurar variables de entorno.
5. Levantar la base de datos.
6. Aplicar migraciones.
7. Generar componentes templ.
8. Ejecutar la aplicacion.
```

Comandos orientativos:

```bash
go mod download
templ generate
go run ./cmd/web
```

La aplicacion deberia quedar disponible en:

```text
http://localhost:8080
```

## 5. Ejecucion local con Docker Compose

Docker Compose permitira levantar la aplicacion y la base de datos de forma reproducible.

Servicios previstos:

- `app`: aplicacion Go.
- `db`: base de datos PostgreSQL.

Estructura esperada:

```text
Dockerfile
docker-compose.yml
.env.example
```

Comandos previstos:

```bash
docker compose up --build
```

Para detener el entorno:

```bash
docker compose down
```

Para eliminar volumenes de base de datos durante desarrollo:

```bash
docker compose down -v
```

## 6. Base de datos

La base de datos recomendada es PostgreSQL.

Motivos:

- Soporta relaciones y restricciones.
- Permite claves unicas para evitar likes y follows duplicados.
- Es adecuada para despliegue real.
- Tiene buen soporte en plataformas cloud.

Tablas principales:

- users.
- sessions.
- posts.
- comments.
- likes.
- follows.

## 7. Migraciones

Las migraciones se almacenaran en:

```text
migrations/
```

Objetivos:

- Crear estructura inicial.
- Mantener cambios de base de datos versionados.
- Facilitar ejecucion local y despliegue.

Ejemplo de nombres:

```text
001_create_users.sql
002_create_sessions.sql
003_create_posts.sql
004_create_comments.sql
005_create_likes.sql
006_create_follows.sql
```

Buenas practicas:

- Las migraciones deben ser pequenas y claras.
- Deben poder ejecutarse en orden.
- Las restricciones importantes deben definirse en base de datos.
- Los indices deben crearse cuando sean necesarios para rendimiento.

## 8. Dockerfile previsto

El Dockerfile debera:

- Usar una imagen oficial de Go.
- Descargar dependencias.
- Generar componentes templ si es necesario.
- Compilar la aplicacion.
- Copiar archivos estaticos.
- Ejecutar el binario final.

Flujo general:

```text
1. Imagen de build con Go.
2. Copia del codigo.
3. Descarga de dependencias.
4. Generacion de templ.
5. Compilacion del binario.
6. Imagen final ligera.
7. Ejecucion del binario.
```

## 9. Integracion continua

La integracion continua se utilizara para comprobar automaticamente que el proyecto mantiene calidad tecnica.

Comprobaciones recomendadas:

```bash
go test ./...
go vet ./...
go build ./cmd/web
```

Si el proyecto usa templ, tambien debera ejecutarse:

```bash
templ generate
```

## 10. Workflow CI propuesto

Pipeline recomendado:

```text
1. Checkout del repositorio.
2. Configurar Go.
3. Instalar templ.
4. Descargar dependencias.
5. Generar templates.
6. Ejecutar tests.
7. Ejecutar go vet.
8. Compilar aplicacion.
```

Ejemplo conceptual para GitHub Actions:

```yaml
name: CI

on:
  push:
  pull_request:

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Setup Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.22'

      - name: Install templ
        run: go install github.com/a-h/templ/cmd/templ@latest

      - name: Download dependencies
        run: go mod download

      - name: Generate templ
        run: templ generate

      - name: Test
        run: go test ./...

      - name: Vet
        run: go vet ./...

      - name: Build
        run: go build ./cmd/web
```

## 11. Opciones de despliegue

### Render

Ventajas:

- Facil de configurar.
- Soporta servicios web.
- Permite conectar PostgreSQL.
- Adecuado para proyectos academicos.

### Railway

Ventajas:

- Despliegue rapido.
- Buen soporte para PostgreSQL.
- Configuracion sencilla de variables.

### Fly.io

Ventajas:

- Buen soporte para aplicaciones Go.
- Permite despliegue con Docker.
- Adecuado para aplicaciones web pequenas.

### VPS con Docker

Ventajas:

- Mayor control.
- Permite usar Docker Compose.
- Buena opcion si ya se dispone de servidor.

Inconveniente:

- Requiere mas configuracion manual.

## 12. Estrategia recomendada de despliegue

Para este proyecto se recomienda:

1. Preparar Dockerfile y docker-compose para ejecucion local.
2. Mantener la aplicacion funcionando localmente.
3. Configurar CI para validar tests y build.
4. Desplegar en Render o Railway si el tiempo lo permite.
5. Documentar la URL final en la memoria.

Si no se llega al despliegue externo, se documentara la ejecucion local con Docker como alternativa reproducible.

## 13. Archivos necesarios para despliegue

Archivos previstos:

```text
Dockerfile
docker-compose.yml
.env.example
README.md
migrations/
static/
```

## 14. README

El README debe incluir:

- Descripcion del proyecto.
- Tecnologias usadas.
- Requisitos previos.
- Instalacion.
- Variables de entorno.
- Ejecucion local.
- Ejecucion con Docker.
- Tests.
- Migraciones.
- Despliegue.
- Credenciales de prueba si se crean.

## 15. Seguridad en despliegue

Medidas necesarias:

- No subir secretos reales al repositorio.
- Usar `.env.example` para documentar variables.
- Configurar `SESSION_SECRET` seguro en produccion.
- Usar conexion segura a base de datos si la plataforma lo requiere.
- Activar cookies seguras en entorno de produccion si se usa HTTPS.
- Evitar logs con informacion sensible.

## 16. Gestion de archivos estaticos

Los archivos estaticos se almacenaran en:

```text
static/
```

Contenido previsto:

- CSS.
- Imagenes.
- JavaScript minimo si fuera necesario.

El servidor Go debera exponer estos archivos mediante una ruta como:

```text
/static/
```

## 17. Build de templ

templ requiere generar codigo Go a partir de los ficheros `.templ`.

Durante desarrollo:

```bash
templ generate
```

Antes de tests o build:

```bash
templ generate
go test ./...
go build ./cmd/web
```

En CI y Docker tambien se debera incluir este paso si el codigo generado no se sube al repositorio.

## 18. Plan de despliegue paso a paso

### Paso 1: Preparar configuracion

- Crear `.env.example`.
- Definir variables necesarias.
- Documentar configuracion en README.

### Paso 2: Preparar base de datos

- Crear migraciones.
- Probar migraciones localmente.
- Configurar PostgreSQL en Docker Compose.

### Paso 3: Preparar Docker

- Crear Dockerfile.
- Crear docker-compose.yml.
- Probar `docker compose up --build`.

### Paso 4: Preparar CI

- Crear workflow.
- Ejecutar tests.
- Ejecutar build.
- Corregir fallos detectados.

### Paso 5: Desplegar

- Elegir plataforma.
- Configurar variables de entorno.
- Conectar base de datos.
- Aplicar migraciones.
- Ejecutar aplicacion.
- Comprobar URL final.

### Paso 6: Documentar

- Anadir instrucciones al README.
- Incluir capturas en la memoria.
- Documentar errores y soluciones.

## 19. Comprobaciones finales

Antes de entregar:

```text
[ ] La aplicacion arranca localmente
[ ] Docker Compose funciona
[ ] Las variables de entorno estan documentadas
[ ] Las migraciones estan disponibles
[ ] Los tests pasan
[ ] El build funciona
[ ] CI esta configurado
[ ] README contiene instrucciones claras
[ ] No hay secretos reales en el repositorio
[ ] El despliegue esta documentado
[ ] La URL final esta incluida si existe
```

## 20. Evidencias para la memoria

Para la memoria final se podran incluir:

- Captura de la aplicacion ejecutandose localmente.
- Captura de Docker Compose funcionando.
- Captura del pipeline CI.
- Captura del despliegue en Render, Railway o Fly.io.
- Explicacion de variables de entorno.
- Explicacion del proceso de despliegue.
- Problemas encontrados y soluciones.

## 21. Relacion con el uso de IA

La IA se utilizara para:

- Proponer configuracion inicial de Docker.
- Revisar el Dockerfile.
- Generar workflow de CI.
- Revisar errores de despliegue.
- Documentar pasos de instalacion.
- Mejorar el README.

Cada intervencion importante se documentara en:

```text
docs/ia/
```

## 22. Conclusión

La estrategia de despliegue se basa en reproducibilidad, configuracion mediante variables de entorno, Docker, integracion continua y documentacion clara.

Esto permite demostrar que el proyecto no solo funciona en desarrollo, sino que tambien esta preparado para ejecutarse, probarse y desplegarse de forma ordenada.
