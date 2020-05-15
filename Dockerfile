FROM scratch
WORKDIR $GOPATH/src/github.com/adeljck/makespace
COPY . $GOPATH/src/github.com/adeljck/makespace

EXPOSE 3000
ENTRYPOINT ["./makespace"]