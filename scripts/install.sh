#!/usr/bin/env bash
set -e
cd "$(dirname "$0")/../internal"
go build ./...
