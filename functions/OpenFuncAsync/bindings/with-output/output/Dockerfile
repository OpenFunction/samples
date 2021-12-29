FROM golang:1.15 as builder

WORKDIR /app

COPY go.mod go.mod

RUN go mod download

COPY goapp.go main.go

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o main main.go

# Use distroless as minimal base image to package the openfunction binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM openfunction/distroless-static:nonroot
WORKDIR /
COPY --from=builder /app/main .
USER nonroot:nonroot

ENTRYPOINT ["/main"]
