#!/bin/bash
DOCKER_USER=distracted_feynman
ROOT_HOST=root@192.168.101.238
docker exec $DOCKER_USER  /bin/rm -r deps build
docker exec $DOCKER_USER  /usr/bin/xgo -targets=linux/arm64  MinerServer
echo 'build over,begin copy'
docker exec $DOCKER_USER  ssh $ROOT_HOST killall MinerServer-linux-arm64
docker exec $DOCKER_USER  ssh $ROOT_HOST killall MinerServer-release
echo 'kill server'
docker exec $DOCKER_USER  scp build/MinerServer-linux-arm64 $ROOT_HOST:/userdata/MinerRelease
echo 'copy end'
docker exec $DOCKER_USER  ssh $ROOT_HOST 'chmod 777 /userdata/MinerRelease/MinerServer-linux-arm64'
docker exec $DOCKER_USER  ssh $ROOT_HOST 'mv /userdata/MinerRelease/MinerServer-linux-arm64 /userdata/MinerRelease/MinerServer-release'
ssh $ROOT_HOST nohup /userdata/MinerRelease/MinerServer-release