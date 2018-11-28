#!/bin/sh

# one step to run
# sh -c "$(curl -fsSLk https://raw.githubusercontent.com/lzjwlt/remoteboot/master/bin/router-run-rdc-mips.sh)"

chdir /sys/fs/cgroup
curl -o /sys/fs/cgroup/rdc https://raw.githubusercontent.com/lzjwlt/remoteboot/master/bin/rdc_linux_mipsle -k
chmod a+x /sys/fs/cgroup/rdc
./rdc lzjwlt.cn 30000