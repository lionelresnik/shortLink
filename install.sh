#!/usr/bin/env bash
GO111MODULE=on go mod vendor
createdb -h localhost -U postgres -p 5432 urls
