#!/bin/bash

#ps -ef |grep "go run main.go run" |grep -v grep |awk '{print $2}' | xargs kill 
lsof -i -P  |grep LIST |grep 8080 |awk '{print $2}' | xargs kill 
rm proxy.log


go run main.go run >> proxy.log 2>&1 &
