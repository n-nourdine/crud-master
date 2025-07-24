#!/bin/bash
sudo apt-get install -y rabbitmq-server
sudo systemctl enable rabbitmq-server
sudo systemctl start rabbitmq-server
sudo rabbitmqctl add_user $RABBITMQ_USER $RABBITMQ_PASSWORD
sudo rabbitmqctl set_permissions -p / $RABBITMQ_USER ".*" ".*" ".*"