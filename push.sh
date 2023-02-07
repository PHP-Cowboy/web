#!/usr/bin/env bash

echo 'rm...'
rm -rf crmPkg
echo 'building...'
GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o crmPkg

echo 'COPY...'
scp crmPkg root@121.196.60.92:/app/crm/temp/
echo 'COPY DONE'
