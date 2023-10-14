# STAGE 1: BUILD
FROM golang:1.21.1-alpine AS build

WORKDIR /app

# INSTALL DEPENDENCIES
COPY go.mod ./
COPY go.sum ./

RUN go mod download

# COPY APP
COPY backend .

RUN CGO_ENABLED=0 GOOS=linux go build -C cmd -o /app 

# STAGE 2: CREATE CONTAINER
FROM alpine:latest

COPY --from=build /app app
COPY frontend frontend

ENV port=8080
ENV db-user=admin
ENV db-pass=adminpass
ENV db-namt=myhome
ENV db-host=db
ENV db-port=5432

EXPOSE 8080

CMD ["./app"] 