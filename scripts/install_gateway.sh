#!/bin/bash
cd /home/vagrant/api-gateway
/usr/local/go/bin/go mod tidy
pm2 start /usr/local/go/bin/go --name api-gateway -- run main.go