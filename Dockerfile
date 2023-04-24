FROM golang:latest AS build

WORKDIR /usr/src/shortly
COPY go.* .

RUN go mod download

COPY cmd cmd/
COPY internal internal/

RUN CGO_ENABLED=0 go build cmd/shortly/shortly.go

FROM scratch

WORKDIR /

COPY --from=build /usr/src/shortly/shortly .

EXPOSE 8080

CMD [ "/shortly" ]
