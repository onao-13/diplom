# STAGE 1: BUILD
FROM golang:1.21.1-alpine AS build

WORKDIR /app

# INSTALL DEPENDENCIES
COPY go.mod ./
COPY go.sum ./

RUN go mod download

# COPY APP
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app 

# STAGE 2: CREATE CONTAINER
FROM alpine:latest

COPY --from=build /app app

EXPOSE 8080

CMD ["./app"] 