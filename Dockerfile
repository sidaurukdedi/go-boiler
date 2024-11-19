# Image Builder
FROM playcourt/golang:1.23 AS go-builder

LABEL maintainer="sidauruk.dedi@gmail.com"

# Set Working Directory
WORKDIR /usr/src/app

# Copy Source Code
COPY . ./

# Dependencies installation and binary file builder
RUN make install \
  && make build


# Final Image
# ---------------------------------------------------
FROM dimaskiddo/alpine:base

# Set Working Directory
WORKDIR /usr/src/app

# Copy Anything The Application Needs
COPY --from=go-builder /tmp/app ./

USER user

# Remove the command below if the service doesn't need JWT Signer / Parser
# Check the Makefile
# COPY --from=go-builder /tmp/secret secret

# Expose Application Port
EXPOSE 9000
HEALTHCHECK --interval=30s --timeout=10s --retries=3 CMD curl -f http://localhost:9000/go-boiler || exit 1

# Run The Application
CMD ["./app"]
