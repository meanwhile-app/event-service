FROM golang:1.20.2-alpine AS builder

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

# Static build by parameters ( Can run on scratch )
# Ref: https://medium.com/@diogok/on-golang-static-binaries-cross-compiling-and-plugins-1aed33499671
# RUN go build -a -tags netgo -ldflags '-w -extldflags "-static"' -v -o /usr/local/bin/app main.go

# Static build by CGO_ENABLED=0 ( Can run on scratch )
RUN CGO_ENABLED=0 go build -v -o /usr/local/bin/app main.go

# Default build ( CANNOT run on scratch )
# RUN go build -v -o /usr/local/bin/app main.go

FROM scratch
# FROM alpine:3.16

COPY --from=builder /usr/local/bin/app /usr/local/bin/app

CMD ["app"]
