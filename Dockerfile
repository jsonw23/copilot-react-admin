FROM golang:1.18

WORKDIR /usr/src/app

# pre-fetch the dependencies in a separate layer
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# copy in the full codebase and build
COPY . .
RUN go build -v -o /usr/local/bin/app ./

# open the default port and run the executable
EXPOSE 8080
CMD ["app"]