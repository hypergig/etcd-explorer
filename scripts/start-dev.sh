#! /usr/bin/env bash
set -xeuo pipefail
repo_root=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )/.." &> /dev/null && pwd )

"${repo_root}/scripts/start-etcd.sh" &
etcd_pid=$!
trap "kill ${etcd_pid}" EXIT

exec npm run dev
