#!/usr/bin/env bash
docker build . -t microportal/login-service:${1} --no-cache
