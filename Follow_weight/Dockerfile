# Imagen base oficial de Go
FROM golang:1.21

# Crear directorio de trabajo
WORKDIR /app

# Copiar go.mod y descargar dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto del código fuente
COPY . .

# Compilar la aplicación
RUN go build -o seguimiento ./main.go

# Exponer el puerto de escucha del microservicio
EXPOSE 3004

# Comando para ejecutar el binario compilado
CMD ["./seguimiento"]
