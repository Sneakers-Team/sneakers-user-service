FROM golang

WORKDIR /usr/src/sn-user-service

EXPOSE 8081

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

CMD [ "go", "run", "cmd/main.go" ]