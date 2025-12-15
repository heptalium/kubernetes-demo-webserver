FROM docker.io/library/golang:latest
WORKDIR /go/src/app
COPY . .
ARG VARIANT="black"
RUN CGO_ENABLED=0 go install -ldflags="-s -w -X main.variant=${VARIANT}" .

FROM scratch
COPY --from=0 /go/bin/webserver /webserver
ENTRYPOINT ["/webserver"]
EXPOSE 80
