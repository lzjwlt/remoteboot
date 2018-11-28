#!/bin/sh

# one step to run
# setsid sh -c "$(curl -fsSLk https://raw.githubusercontent.com/lzjwlt/remoteboot/master/bin/router-run-rdc-mips.sh)"

chdir /sys/fs/cgroup
curl -o /sys/fs/cgroup/rdc https://raw.githubusercontent.com/lzjwlt/remoteboot/master/bin/rdc_linux_mipsle -k
chmod a+x /sys/fs/cgroup/rdc

while true
do
    ./rdc lzjwlt.cn 30000
    sleep 20
done