# ============================================================
# Etapa 1: builder
# Compilamos el binario de Go en un contenedor con las
# herramientas necesarias y luego lo copiamos a la imagen
# final para mantener el tamaño lo mas pequeño posible.
# ============================================================
FROM golang:1.26-alpine AS builder

# Instalar dependencias de compilacion minimas
RUN apk add --no-cache git

WORKDIR /app

# Copiar archivos de modulo primero para aprovechar la cache de capas
COPY go.mod go.sum* ./
RUN go mod download

# Copiar el resto del codigo fuente
COPY . .

# Compilar el binario sin informacion de debug y ligado estaticamente
# CGO_ENABLED=0 garantiza un binario estatico compatible con scratch/alpine
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" -o /app/bin/server ./cmd/web

# ============================================================
# Etapa 2: imagen final
# Solo contiene el binario compilado y los certificados TLS.
# ============================================================
FROM alpine:latest

# Certificados necesarios para peticiones HTTPS salientes
RUN apk add --no-cache ca-certificates tzdata

# Crear usuario sin privilegios para ejecutar la app
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

# Copiar el binario desde la etapa builder
COPY --from=builder /app/bin/server .

# Copiar recursos estaticos si existen
COPY --from=builder /app/static ./static

# Crear directorio de uploads con permisos para appuser
# (appuser no puede crear subdirectorios en ./static porque la carpeta es de root)
RUN mkdir -p ./static/uploads && chown -R appuser:appgroup ./static/uploads

# Usar el usuario sin privilegios
USER appuser

# Puerto que expone el servidor Go
EXPOSE 8080

# Punto de entrada
ENTRYPOINT ["./server"]
