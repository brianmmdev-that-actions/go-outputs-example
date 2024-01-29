FROM golang:latest

WORKDIR /app
COPY go.mod go.mod
COPY *.go ./

RUN go get
RUN go build -o build/action-bin .

RUN cp build/action-bin /usr/bin/action-bin
ENTRYPOINT ["/usr/bin/action-bin"]