FROM golang:1.22.1 as builder
RUN rm -rf /app

RUN mkdir /app

WORKDIR /app/altscore

COPY ./main.go .
COPY ./go.mod .
COPY ./go.sum .
WORKDIR /app/altscore
# Compilar la aplicación
RUN CGO_ENABLED=0 go build -o myapp ./main.go

# Etapa 2: Crear una imagen mínima
FROM alpine:latest
RUN apk add --no-cache tzdata
ENV TZ=America/Bogota
WORKDIR /app

# Copiar solo los archivos necesarios desde la etapa de construcción
COPY --from=builder /app/altscore/myapp .

# Comando para ejecutar la aplicación
CMD ["./myapp"]