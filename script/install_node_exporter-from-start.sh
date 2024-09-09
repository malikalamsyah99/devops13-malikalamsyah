#!/bin/bash

# Create prometheus group
sudo groupadd --system prometheus

# Create prometheus user and assign to prometheus group
sudo useradd --system -s /sbin/nologin -g prometheus prometheus

# Download Node Exporter
wget https://github.com/prometheus/node_exporter/releases/download/v1.8.2/node_exporter-1.8.2.linux-amd64.tar.gz

# Extract the downloaded file
tar -xvzf node_exporter-1.8.2.linux-amd64.tar.gz

# Move into the extracted directory
cd node_exporter-1.8.2.linux-amd64/

# Move the Node Exporter binary file to /usr/local/bin/
sudo mv node_exporter /usr/local/bin/

# Create systemd service for Node Exporter
sudo tee /etc/systemd/system/node-exporter.service > /dev/null <<EOL
[Unit]
Description=Prometheus exporter for machine metrics

[Service]
Restart=always
User=prometheus
ExecStart=/usr/local/bin/node_exporter
ExecReload=/bin/kill -HUP \$MAINPID
TimeoutStopSec=20s
SendSIGKILL=no

[Install]
WantedBy=multi-user.target
EOL

# Reload systemd daemon
sudo systemctl daemon-reload

# Enable and start Node Exporter service
sudo systemctl enable --now node-exporter.service

echo "Node Exporter installation and service setup completed."

