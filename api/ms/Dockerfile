FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build -o rover-api

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build-env /src/rover-api /app/
ENTRYPOINT ./rover-api
