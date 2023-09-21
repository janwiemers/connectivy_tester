FROM 224935473367.dkr.ecr.eu-central-1.amazonaws.com/ratepay-proxy:latest as builder

WORKDIR /build
COPY . .

RUN apk add upx make gcc libc-dev git && \
    go env -w GOPROXY=direct
ENV GO111MODULE=auto
ENV CGO_ENABLED: 0
ENV GOOS: linux
RUN make build
RUN upx -9 connectivity_tester

FROM 224935473367.dkr.ecr.eu-central-1.amazonaws.com/ratepay-proxy:latest

WORKDIR /app

RUN addgroup -g 1000 -S app && \
  adduser -u 1000 -S app -G app && \
  date -u > BUILD_TIME

COPY --from=Builder --chown=app:app /build/connectivity_tester /app/connectivity_tester
RUN chown -R app:app /app
USER app
CMD ["./connectivity_tester", "-port=8080"]