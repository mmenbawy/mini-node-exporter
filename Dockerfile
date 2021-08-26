FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY utils ./utils
COPY main.go ./

RUN go build .
RUN go mod tidy

RUN adduser -S non-root 

USER        non-root

EXPOSE      23333

CMD [ "go", "run", "." ]
