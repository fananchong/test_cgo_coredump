#!/bin/bash

(gdb ./main ./core >a.log) <<GDBEOF
thread apply all bt
GDBEOF

n=$(grep "runtime.rt0_go" a.log | awk -F' ' '{printf $1 }')
n=${n:1}
line=$(grep -n "runtime.rt0_go" a.log | awk -F':' '{printf $1 }')

no=$((line - n - 2))
sed -n "$no,${line}p" a.log
