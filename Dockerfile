FROM golang:1.14

WORKDIR /go/src/github.com/imdario/bofhchan
COPY . .
COPY internal/ .
COPY cmd/ .
RUN go build -o bofhchan ./cmd/bofhchan/...

FROM scratch
COPY --from=0 /go/src/github.com/imdario/bofhchan/bofhchan /bofhchan
CMD ["/bofhchan"]