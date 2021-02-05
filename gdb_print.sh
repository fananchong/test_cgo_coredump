#!/bin/bash

binname=$1
corefile=$2

(gdb $binname $corefile &> ./$corefile.log) <<GDBEOF
thread apply all bt
GDBEOF

n=$(grep "runtime.asmcgocall" ./$corefile.log | awk -F' ' '{printf $1 }')
echo "n:$n"
if [ "$n" == "" ]; then
    exit 0
fi
n=${n:1}
line=$(grep -n "runtime.asmcgocall" ./$corefile.log | awk -F':' '{printf $1 }')

no=$((line - n - 2))
sed -n "$no,${line}p" ./$corefile.log
