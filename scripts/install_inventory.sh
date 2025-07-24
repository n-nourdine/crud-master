#!/bin/bash
cd /home/vagrant/inventory-app
/usr/local/go/bin/go mod tidy
pm2 start /usr/local/go/bin/go --name inventory-app -- run main.go