#!/usr/bin/env bash

set -e

APP_NAME="frontdoor"
INSTALL_DIR="/opt/frontdoor"
SERVICE_FILE="frontdoor.service"


echo "Creating install directory..."
sudo mkdir -p "$INSTALL_DIR"

echo "Copying files..."

sudo cp "./$APP_NAME" "$INSTALL_DIR/"
sudo cp -r templates "$INSTALL_DIR/"
sudo cp services.json "$INSTALL_DIR/"

echo "🔧 Installing systemd service..."

if [[ -f "$SERVICE_FILE" ]]; then
    sudo cp "$SERVICE_FILE" /etc/systemd/system/
    sudo systemctl daemon-reload
    sudo systemctl enable frontdoor
    sudo systemctl restart frontdoor
    echo "✅ Service installed and restarted"
else
    echo "⚠️ No systemd service file found, skipping"
fi

