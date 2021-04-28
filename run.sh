#!/bin/bash
DOCKER_USER=distracted_feynman
ROOT_HOST=root@192.168.101.231
docker exec $DOCKER_USER  ssh $ROOT_HOST killall MinerServer-linux-arm64
echo 'kill server'
docker exec $DOCKER_USER  ssh $ROOT_HOST '/root/build_test/MinerServer-linux-arm64'
