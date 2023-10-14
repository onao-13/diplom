# STAGE 1: BUILD
FROM golang:1.21.1-alpine AS build

WORKDIR app

# INSTALL DEPENDENCIES
COPY backend/go.mod ./
COPY backend/go.sum ./

RUN go mod download

COPY backend /app/backend

RUN CGO_ENABLED=0 GOOS=linux go build -o /app 

# STAGE 2: CREATE CONTAINER
FROM alpine:latest

COPY --from=build /app app
#COPY frontend ./

EXPOSE 8085

CMD ["./app"]

ENV port=8085
ENV db-host=db
ENV db-port=5432
ENV db-user=admin
ENV db-pass=adminpass
ENV db-name=myhome