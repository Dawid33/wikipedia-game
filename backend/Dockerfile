# Build the static files and executable backend
FROM alpine:edge

RUN apk update && apk upgrade

# Install dependencies
RUN apk add --no-cache go

# Copy all files to docker image
RUN mkdir -p /backend/build
COPY / /backend/build
WORKDIR /backend/build
RUN go mod tidy
RUN go build -o executable && mv executable ..

WORKDIR /backend
CMD ./executable