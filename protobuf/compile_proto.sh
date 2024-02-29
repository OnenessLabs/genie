#!/usr/bin/env bash

# fail automatically as soon as an error is detected
set -e

buf mod update
buf generate
