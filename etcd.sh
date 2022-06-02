#!/usr/bin/env bash
docker run -d --rm --name etcd \
    -p 3379:2379 \
    -p 3380:2380 \
    -e ALLOW_NONE_AUTHENTICATION=yes \
    bitnami/etcd:latest