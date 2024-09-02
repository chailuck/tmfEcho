# A MULTI-STAGE DOCKERFILE

# STAGE 1
# use alpine image as builder
FROM golang:alpine AS builder

# golang specific variables
ENV GO111MODULE=on \
  CGO_ENABLED=1 \
  GOOS=linux \
  GOARCH=amd64

# current working directory is /build in the container
WORKDIR /build

# copy over go.mod and go.sum (module dependencies and checksum)
# over to working directory
COPY go.mod .
COPY go.sum .

# download the dependencies
RUN go mod download

# copy our application code into the container
COPY . .




# building the binary called "main"
# RUN go build -o main ./cmd/tmf632.go

# install gcc lib in alpine
RUN apk add build-base

# build the binary and output file as main
RUN go build -o main ./cmd/tmf632.go


RUN echo $(ls /build/)


# STAGE 2
# Build a small image
FROM golang:alpine as appserver

# arguments to be passed during build phase

#ARG JWT_TOKEN_SECRET

# environment variables for the application

#ENV JWT_TOKEN_SECRET=${JWT_TOKEN_SECRET}

# copy from stage-1 image
COPY --from=builder /build/ /build/
COPY --from=builder /build/main /

RUN echo $(ls /)

RUN echo $(ls /build)
#COPY /build/main /


# expose the port to run the application on
EXPOSE 8080

# Command to run
ENTRYPOINT ["/main"]
#ENTRYPOINT [ "sh", "/main" ]
#CMD ["echo", "hello yo!"]
