FROM golang:1.16
COPY . /app
WORKDIR /app
RUN go build -o main .
ENV PORT=80
ENV GIN_MODE=release
ARG git_shasum
ARG version_tag
ENV GIT_SHASUM=$git_shasum
ENV VERSION_TAG=$version_tag
CMD ["./main"]
