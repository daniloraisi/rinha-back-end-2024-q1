FROM golang:alpine

WORKDIR /opt/app

COPY . .

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

EXPOSE 8080

CMD [ "air --build.cmd \"go build -o ./bin/api ./cmd/...\" --build-bin \"./bin/api\"" ]

