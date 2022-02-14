# Build the static files and executable backend
FROM alpine:edge

RUN apk update && apk upgrade

# Install dependencies
RUN apk add --no-cache go

# Copy all files to docker image
RUN mkdir -p /backend_build/
COPY / /backend_build/

WORKDIR /backend_build
RUN go mod tidy
RUN go build -o backend && mv backend ..
WORKDIR /
CMD /backend