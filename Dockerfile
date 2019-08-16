FROM golang:1.13beta1-alpine3.10 AS build-env

# Allow Go to retrieve the dependencies for the build step
RUN apk add --no-cache github

# Secure against running as root
RUN adduser -D -u 10000 tawfeeq
RUN makdir /app/ && chown tawfeeq /app/
USER tawfeeq


ADD . .
WORKDIR /app/
# Compile the binary, we don't want to run the cgo resolver
RUN CGO_ENABLED=0 go build -o /app/go_ssl_server .
# final stage
#FROM alpine:3.10

# Secure against running as root
#RUN adduser -D -u 10000 tawfeeq
#USER tawfeeq

#WORKDIR /
#COPY --from=build-env C:\Users\77995\go\src\github.com\tawfeeq0\go_ssl_server\cert\*.* /
#COPY --from=build-env C:\Users\77995\go\src\github.com\tawfeeq0\go_ssl_server\ /
#EXPOSE 3030

CMD ["/app"]