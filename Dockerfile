# Do not build this yourself, this is for goreleaser
FROM alpine:3.16
COPY crab /
ENTRYPOINT ["/crab"]
