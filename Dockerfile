FROM golang:1.21-alpine
LABEL authors="Tasnim Zotder <hello@tasnimzotder.com>"

# Install make
RUN apk add --no-cache make

WORKDIR /app

COPY go.mod ./
# copy go.sum separately so that it doesn't invalidate the cache
#COPY go.sum ./

RUN go mod download

COPY . .