#!/bin/bash

# Update package list and install dependencies
sudo apt-get update
sudo apt-get install -y adduser libfontconfig1 musl

# Download Grafana package
wget https://dl.grafana.com/oss/release/grafana_11.2.0_amd64.deb

# Install Grafana package
sudo dpkg -i grafana_11.2.0_amd64.deb

# Reload systemd daemon
sudo /bin/systemctl daemon-reload

# Enable Grafana to start on boot
sudo /bin/systemctl enable grafana-server

# Start Grafana server
sudo /bin/systemctl start grafana-server

echo "Grafana installation completed and service started."