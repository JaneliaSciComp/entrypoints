# Entrypoints

Simple support for multiple entrypoint scripts in an OCI container. Just install this tiny binary as your ENTRYPOINT and put all of your entrypoint shell scripts in one directory. The binary takes care of providing a listing of valid "apps" and lets the user invoke them without knowing the internals of the container filesystem.

## Installation

The following Dockerfile commands install the `entrypoint` binary into /app/entrypoints and set it as the entrypoint:

```
ARG ENTRYPOINTS_VERSION=latest
WORKDIR app
RUN curl -sL https://github.com/JaneliaSciComp/entrypoints/releases/download/${ENTRYPOINTS_VERSION}/entrypoints_linux_x86_64.tar.gz | tar -xz -C /app
ENTRYPOINT [ "/app/entrypoints", "-s", "/app/scripts/cmd" ]
```

A multi-stage build is recommended to improve build speed:

```
FROM janeliascicomp/common-tools:1.0.2 as entrypoints-builder
ARG ENTRYPOINTS_VERSION=0.1.1
RUN curl -sL https://github.com/JaneliaSciComp/entrypoints/releases/download/${ENTRYPOINTS_VERSION}/entrypoints_linux_x86_64.tar.gz | tar -xz -C /tmp
```

In the final container:
```
COPY --from=entrypoints-builder --chown=root:root /tmp/entrypoints /app/entrypoints
ENTRYPOINT [ "/app/entrypoints", "-s", "/app/scripts/cmd" ]

```

See also: [SCIF](https://sci-f.github.io/)

## Development

### Build

```
go install
```

### Release
```
VERSION=<version number>
git tag -a $VERSION -m "Release $VERSION"
git push origin $VERSION
goreleaser --rm-dist
```

