#!/usr/bin/env bash

set -eux

# Copies over welcome message
mkdir -p /usr/local/etc/vscode-dev-containers/
cp .devcontainer/welcome-message.txt /usr/local/etc/vscode-dev-containers/first-run-notice.txt
