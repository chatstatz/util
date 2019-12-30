ARG GO_VERSION

FROM golang:${GO_VERSION}
WORKDIR /util
COPY . ./
RUN make install
RUN make test
