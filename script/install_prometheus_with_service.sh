#!/bin/bash

# Download Prometheus
wget https://github.com/prometheus/prometheus/releases/download/v2.53.2/prometheus-2.53.2.linux-amd64.tar.gz

# Extract the downloaded file
tar -xvzf prometheus-2.53.2.linux-amd64.tar.gz

# Move into the extracted directory
cd prometheus-2.53.2.linux-amd64/

# Create prometheus group
sudo groupadd --system prometheus

# Create prometheus user and assign to prometheus group
sudo useradd --system -s /sbin/nologin -g prometheus prometheus

# Move the binary files to /usr/local/bin/
sudo mv prometheus promtool /usr/local/bin/

# Create directory for Prometheus config files
sudo mkdir /etc/prometheus

# Create directory for Prometheus data
sudo mkdir /var/lib/prometheus

# Change ownership of the data directory to prometheus user and group
sudo chown -R prometheus:prometheus /var/lib/prometheus/

# Move the config files to /etc/prometheus/
sudo mv consoles/ console_libraries/ prometheus.yml /etc/prometheus/

# Change to /etc/prometheus directory
cd /etc/prometheus/

# Modify the prometheus.yml configuration
sudo tee prometheus.yml > /dev/null <<EOL
global:
  scrape_interval: 15s
  evaluation_interval: 15s

scrape_configs:
  - job_name: "prometheus"
    static_configs:
      - targets: ["localhost:9090"]
EOL

# Create Prometheus systemd service
sudo tee /etc/systemd/system/prometheus.service > /dev/null <<EOL
[Unit]
Description=Prometheus
Documentation=https://prometheus.io/docs/introduction/overview/
Wants=network-online.target
After=network-online.target

[Service]
User=prometheus
Group=prometheus
Type=simple
ExecStart=/usr/local/bin/prometheus \\
  --config.file /etc/prometheus/prometheus.yml \\
  --storage.tsdb.path /var/lib/prometheus/ \\
  --web.console.templates=/etc/prometheus/consoles \\
  --web.console.libraries=/etc/prometheus/console_libraries

[Install]
WantedBy=multi-user.target
EOL

# Reload systemd daemon
sudo systemctl daemon-reload

# Enable and start Prometheus service
sudo systemctl enable --now prometheus

echo "Prometheus installation and service setup completed."
