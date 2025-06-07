# To build this with secrets: (sudo) docker build --secret id=sshkey,src=/path/to/id_rsa -t credstack-api:latest .
# To build normally: (sudo) docker build . -t credstack-api:latest

FROM golang:1.24.2-alpine AS builder

RUN apk --no-cache add ca-certificates git openssh-client

# Since credstack-lib is private, we need a way of providing the build context secrets
# In the event this repository ever becomes public, we can remove this
RUN mkdir -p /root/.ssh && chmod 700 /root/.ssh
RUN --mount=type=secret,id=sshkey \
    cp /run/secrets/sshkey /root/.ssh/id_rsa && \
    chmod 600 /root/.ssh/id_rsa && \
    ssh-keyscan github.com >> /root/.ssh/known_hosts

# We also need to re-write URLs here as when dependencies get added with 'go get', they use HTTPS scheme and we need
# private key scheme to clone private repo's
RUN git config --global url."git@github.com:".insteadOf "https://github.com/"

# Describes the OS/Architecture we want to build for and instructs the conmpiler to build static binaries
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPRIVATE=github.com/stevezaluk/credstack-lib

WORKDIR /app

# Copy source files
COPY . .

RUN go mod download

# Strip symbols and debugging information
RUN go build -ldflags="-s -w" -o app

# Once the build process is completed. We remove any key files that may be lingering. Anything provided in /run/secrets
# automatically gets removed by docker after building
RUN rm -rf /root/.ssh/id_rsa

FROM gcr.io/distroless/static-debian12

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/app /app/app

USER nonroot:nonroot

WORKDIR /app

ENTRYPOINT ["/app/app"]
