#!/bin/bash

repo=("kitex" "kitex-mux")

# build
. ./scripts/build_thrift.sh

# benchmark
. ./scripts/kill_servers.sh
core=0
for ((i = 0; i < ${#repo[@]}; i++)); do
  rp=${repo[i]}

  # server start
  nohup taskset -c $core-$(($core + 3)) ./output/bin/${rp}_reciever >> output/log/nohup.log 2>&1 &
  echo "server $rp running at cpu $core-$(($core + 3))"
  core=$(($core + 4))
  sleep 1
done
