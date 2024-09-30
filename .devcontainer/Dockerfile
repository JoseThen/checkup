# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.194.0/containers/alpine/.devcontainer/base.Dockerfile
# [Choice] Alpine version: 3.14, 3.13, 3.12, 3.11
FROM mcr.microsoft.com/vscode/devcontainers/base:0-alpine-3.14

COPY --from=golang:1.23-alpine /usr/local/go/ /usr/local/go/

ENV PATH="/usr/local/go/bin:${PATH}"

# ** [Optional] Uncomment this section to install additional packages. **
RUN apk update \
  && apk add --no-cache curl wget make python3 py3-pip \
  && alias python=python3 \
  && pip install pre-commit \
  && mkdir /home/vscode/.ssh
