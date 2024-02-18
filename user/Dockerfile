# STAGE 1: BUILD
FROM golang:1.21.1-alpine AS build

WORKDIR /app

# INSTALL DEPENDENCIES
COPY backend/go.mod ./
COPY backend/go.sum ./

RUN go mod download

# COPY APP
COPY backend ./

RUN CGO_ENABLED=0 GOOS=linux go build -C ./cmd -o /backend

# STAGE 2: CREATE CONTAINER
FROM alpine:latest

COPY --from=build /backend app
COPY frontend frontend

ENV port=8120
ENV db-user=admin
ENV db-pass=adminpass
ENV db-name=myhome
ENV db-host=db
ENV db-port=5432

CMD ["./app"]