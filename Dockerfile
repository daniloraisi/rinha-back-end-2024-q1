FROM golang:alpine AS build

WORKDIR /app

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./api ./cmd/...

FROM scratch AS prod

ARG HTTP_PORT=8080

WORKDIR /opt/app

COPY --from=build /app/api ./

EXPOSE ${HTTP_PORT}

CMD ["./api"]
