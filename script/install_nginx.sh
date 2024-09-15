#!/bin/bash

## Install the nginx
apt update
apt -y install nginx

## Running the service
systemctl enable  nginx.service
systemctl start  nginx.service
a=`systemctl is-active nginx.service`

if [ $a == "active" ]; then echo  "Nginx has been started successfully."; 
else echo "Failed to start Nginx" && exit 1 
fi