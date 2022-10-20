FROM golang:1.18-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o announcerd

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /app/announcerd /announcerd

EXPOSE 6000 

USER nonroot:nonroot

ENTRYPOINT ["/announcerd"]
