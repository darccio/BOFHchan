FROM golang:1.14 as build

WORKDIR /go/src/github.com/imdario/bofhchan
COPY . /go/src/github.com/imdario/bofhchan
RUN CGO_ENABLED=0 go build -ldflags '-w -extldflags "-static"' -o bofhchan -a ./cmd/bofhchan/... 

FROM scratch
COPY --from=build /go/src/github.com/imdario/bofhchan/bofhchan /bofhchan
COPY --from=build /go/src/github.com/imdario/bofhchan/templates/ /templates/
COPY --from=build /go/src/github.com/imdario/bofhchan/static/ /templates/
CMD ["/bofhchan"]

EXPOSE 3000