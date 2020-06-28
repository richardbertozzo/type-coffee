ARG GO_VERSION=1.14.0

# BUILD STAGE
FROM golang:${GO_VERSION}-alpine as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Install the Certificate-Authority certificates for the app to be able to make calls to HTTPS endpoints.
# Install git
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache ca-certificates git

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /app

COPY go.mod .
COPY go.sum .

# Fetch dependencies first; they are less susceptible to change on every build
# and will therefore be cached for speeding up the next build
RUN go mod download

COPY . .

# Build the executable to `/app`. Mark the build as statically linked.
RUN go build -o server ./cmd/server

# FINAL STAGE: the running container.
FROM scratch AS final

# Import the Certificate-Authority certificates for enabling HTTPS.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# Import the compiled executable from the first stage.
COPY --chown=0:0 --from=builder /app/server /server

EXPOSE 3000

ENTRYPOINT ["/server"]