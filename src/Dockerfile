FROM golang:1.24-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o tips-microservice

FROM alpine:edge

WORKDIR /src

COPY --from=build /src/tips-microservice .

RUN chmod +x /src/tips-microservice

RUN apk --no-cache add ca-certificates tzdata

EXPOSE 1234

ENTRYPOINT [ "/src/tips-microservice" ]