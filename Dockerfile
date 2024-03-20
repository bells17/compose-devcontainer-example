# Build the manager binary
FROM golang:1.21 as builder
WORKDIR /work
ARG LDFLAGS

COPY ./ .
RUN CGO_ENABLED=0 go build -ldflags="-w -s ${LDFLAGS}" -o app ./app/main.go

# the controller image
FROM gcr.io/distroless/static:latest
COPY --from=builder /work/app ./
USER 10000:10000

ENTRYPOINT ["./app"]
