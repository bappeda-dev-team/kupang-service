ARG GO_VERSION=1.24.2


FROM registry.docker.com/library/golang:$GO_VERSION-alpine AS base

# app lives here
WORKDIR /app


# Throw-away build stage to reduce size of final image
FROM base AS build

# Install packages needed to build
RUN apk update -qq && \
    apk add --no-cache git curl tar openjdk17-jre

ARG FLYWAY_VERSION=10.15.0
ENV JAVA_HOME=/usr/lib/jvm/java-17-openjdk
ENV PATH="$JAVA_HOME/bin:/opt/flyway:$PATH"

RUN mkdir -p /opt && \
    curl -fsSL "https://repo1.maven.org/maven2/org/flywaydb/flyway-commandline/${FLYWAY_VERSION}/flyway-commandline-${FLYWAY_VERSION}-linux-x64.tar.gz" \
    | tar -xz -C /opt && \
    ln -s "/opt/flyway-${FLYWAY_VERSION}" /opt/flyway && \
    ln -s /opt/flyway/flyway /usr/local/bin/flyway

COPY . .

RUN go build -o api main.go wire_gen.go

ENTRYPOINT ["/app/api"]

CMD ["app/api"]
