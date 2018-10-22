FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build -o rover-api

FROM alpine
WORKDIR /app
COPY --from=build-env /src/rover-api /app/
ENTRYPOINT ./rover-api
