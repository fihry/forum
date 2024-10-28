FROM golang:1.22.3-alpine

# Install bash and SQLite3 development libraries using apk
RUN apk add --no-cache bash sqlite sqlite-dev gcc musl-dev

WORKDIR /app
# Copy the rest of the application code
COPY . /app

# Copy go.mod and go.sum before running go mod tidy
RUN rm go.sum

# Install dependencies
RUN go get github.com/gofrs/uuid@v4.4.0+incompatible
RUN go get github.com/mattn/go-sqlite3@v1.14.23
RUN go get golang.org/x/crypto

EXPOSE 5000

LABEL version="0.0.1"
LABEL description="This is a Dockerfile for the forum project"

CMD go run . 5000