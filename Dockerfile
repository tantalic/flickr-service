ARG GO_VERSION=1.13.6
ARG ALPINE_VERSION=3.11.3

FROM golang:${GO_VERSION}-alpine as builder

# Download dependencies (this is done in a seperate layer to take advantage
# of Docker layer caching to decrease build times)
WORKDIR /src
COPY ./go.mod ./go.sum ./
RUN go mod download

# Copy source files
COPY ./ ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -installsuffix 'static' -o flickr-service .

# Use alpine linux base, with latest CA roots installed
FROM alpine:${ALPINE_VERSION}
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

COPY --from=builder /src/flickr-service /usr/local/bin/flickr-service

# The following environment variables are used to configure teh application
# 
# ENV HOST
# ENV PORT
# ENV FLICKR_ALBUM
# ENV FLICKR_KEY
# ENV REFRESH_INTERVAL

EXPOSE 3000
CMD /usr/local/bin/flickr-service
