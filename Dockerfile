FROM golang:1.15 as builder
WORKDIR $GOPATH/src/pet_store_rest_api
COPY ./ .
RUN go get
RUN GOOS=linux GOARCH=amd64 go build -v
RUN cp pet_store_rest_api /
RUN cp launch.sh /

FROM golang:1.15 as pet_store
COPY --from=builder /pet_store_rest_api /
COPY --from=builder /launch.sh /
ENTRYPOINT ["/launch.sh"]