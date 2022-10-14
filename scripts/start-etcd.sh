#! /usr/bin/env bash
set -xeuo pipefail
repo_root=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )/.." &> /dev/null && pwd )

exec 2> >(cat | sed 's/^/ETC | /')

exec etcd \
  --name s1 \
  --data-dir "${repo_root}/etcd-data" \
  --listen-client-urls http://0.0.0.0:2379 \
  --advertise-client-urls http://0.0.0.0:2379 \
  --listen-peer-urls http://0.0.0.0:2380 \
  --initial-advertise-peer-urls http://0.0.0.0:2380 \
  --initial-cluster s1=http://0.0.0.0:2380 \
  --initial-cluster-token tkn \
  --initial-cluster-state new \
  --log-level debug \
  --logger zap \
  --log-outputs stderr
