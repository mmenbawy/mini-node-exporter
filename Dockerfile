FROM golang:1.16-alpine

WORKDIR /app

COPY . .
RUN go build . && \
    go mod tidy

RUN adduser -S non-root 
USER    non-root

EXPOSE  23333
CMD [ "go", "run", "." ]
