FROM golang:alpine
WORKDIR /var/lib/gopher
RUN apk update && \
    apk upgrade
ADD go.mod .
ADD go.sum .
ADD . .
RUN source .env
CMD go run src/
