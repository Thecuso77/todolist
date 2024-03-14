FROM golang:alpine as builder

LABEL maintainer="Aglodin D <aglodin@techart.ru>"

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN apk add redis

RUN go mod download -x

#RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

RUN go install -mod=mod github.com/githubnemo/CompileDaemon

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o vuego .

ENTRYPOINT CompileDaemon --directory="/app" --build="go build -a -installsuffix cgo -o /app/vuego" --command="./vuego"

EXPOSE 8080