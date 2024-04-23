#!/usr/bin/env bash
# Only supports Linux

set -eux

apt-get update
apt-get install -y redis-tools
