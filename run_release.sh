#!/bin/bash
DOCKER_USER=distracted_feynman
ROOT_HOST=root@192.168.101.238
docker exec $DOCKER_USER  ssh $ROOT_HOST killall MinerServer-release
echo 'kill server'
docker exec $DOCKER_USER  ssh $ROOT_HOST '/userdata/MinerRelease/MinerServer-release'
