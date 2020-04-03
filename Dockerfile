# Do not build this yourself, this is for goreleaser
FROM scratch
COPY crab /
ENTRYPOINT ["/crab"]