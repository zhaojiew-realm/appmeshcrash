# stage 1
FROM public.ecr.aws/amazonlinux/amazonlinux:2 AS builder
RUN yum update -y && \
    yum install -y ca-certificates unzip tar gzip git golang && \
    yum clean all && \
    rm -rf /var/cache/yum

ENV PATH="${PATH}:/usr/local/go/bin"
ENV GOPATH="${HOME}/go"
ENV PATH="${PATH}:${GOPATH}/bin"
ENV GOPROXY="https://goproxy.cn,direct"

WORKDIR /bin

# go.mod and go.sum go into their own layers.
# This ensures `go mod download` happens only when go.mod and go.sum change.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-s -w --extldflags "-static -fpic"' -a -installsuffix nocgo -o /server .

# stage 2
FROM scratch
COPY --from=builder /server bin/server

# slime
ENTRYPOINT ["/bin/server"]

# debug
# ENTRYPOINT ["bash","-c"]
# CMD ["/bin/server"]