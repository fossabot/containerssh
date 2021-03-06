FROM golang AS build
RUN mkdir -p /go/src/github.com/janoszen/containerssh
WORKDIR /go/src/github.com/janoszen/containerssh/
COPY . /go/src/github.com/janoszen/containerssh
RUN make build

FROM scratch AS run
COPY --from=build /go/src/github.com/janoszen/containerssh/build/containerssh /containerssh
ENTRYPOINT ["/containerssh"]
CMD ["--config", "/etc/containerssh/config.yaml"]
VOLUME /etc/containerssh
VOLUME /var/secrets
EXPOSE 2222
