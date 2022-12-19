FROM golang:latest as build

WORKDIR /src
COPY main.go go.mod ./
RUN go mod download && go mod verify

RUN CGO_ENABLED=0 go build -v -o simple-pop3 .

FROM scratch
COPY --from=build /src/simple-pop3 /simple-pop3

ENTRYPOINT ["/simple-pop3"]
