#!/bin/bash
cd /home/vagrant/billing-app
/usr/local/go/bin/go mod tidy
pm2 start /usr/local/go/bin/go --name billing-app -- run main.go