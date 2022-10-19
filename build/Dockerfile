FROM golang:1.16 AS builder

ARG SERVICE
ARG REVISION_ID
ARG BUILD_DATE
ARG GIT_HASH

WORKDIR /build/

RUN go get github.com/swaggo/swag/cmd/swag

COPY . ./
RUN swag init -d cmd/$SERVICE --parseInternal --parseDependency
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux \
    go build -o ./service \
    -ldflags "-s -X $SERVICE/internal/pkg/version.serviceName=$SERVICE \
    -X $SERVICE/internal/pkg/version.buildID=$REVISION_ID \
    -X $SERVICE/internal/pkg/version.buildDate=$BUILD_DATE \
    -X $SERVICE/internal/pkg/version.gitHash=$GIT_HASH" \
    ./cmd/$SERVICE/...

FROM alpine:3.16.1 AS release
WORKDIR /svc/
COPY --from=builder /build/service ./
RUN find $HOME -type d -exec 'chmod' '555' '{}' ';' && \
    find $HOME -type f -exec 'chmod' '444' '{}' ';' && \
    find $HOME -type f -exec 'chown' 'root:root' '{}' ';' && \
    chmod 555 ./service
USER nobody
ENTRYPOINT [ "./service" ]