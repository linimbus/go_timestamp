#!/bin/bash

for((i=1;i<=10;i++));
do
sleep_time=$((RANDOM % 10))
echo "run test sleep $sleep_time"
sleep $sleep_time
done