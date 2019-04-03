FROM scratch

LABEL maintainer="Wasith Theerapattrathamrong <wasith.t@gmail.com>"

COPY go-rest go-rest

CMD ["go-rest"]
