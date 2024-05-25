FROM golang:1.22-alpine as builder

WORKDIR /go/src/github.com/cheeseNA/owlback

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY cmd/ cmd/
COPY internal/ internal/

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -o app ./cmd/main.go

FROM gcr.io/distroless/base-nossl-debian11

COPY --from=builder /go/src/github.com/cheeseNA/owlback/app /

ENTRYPOINT ["/app"]