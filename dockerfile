# stage 1
FROM golang:1.18.10 AS builder
ENV GOPROXY="https://goproxy.cn,direct"

WORKDIR /bin
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-s -w --extldflags "-static -fpic"' -a -installsuffix nocgo -o /server .

# stage 2
# FROM scratch
FROM public.ecr.aws/amazonlinux/amazonlinux:2
COPY --from=builder /server bin/server

# slime
ENTRYPOINT ["/bin/server"]

# debug
# ENTRYPOINT ["bash","-c"]
# CMD ["/bin/server"]