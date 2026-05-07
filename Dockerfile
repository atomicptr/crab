# Do not build this yourself, this is for goreleaser
FROM alpine:3.23
COPY crab /
ENTRYPOINT ["/crab"]
