#!/bin/sh
# Simple DDNS placeholder
DDNS_ENDPOINT="https://example.com/api/ddns/update"
HOSTNAME=${1:-"rodial-device.local"}
curl -X POST -d "host=$HOSTNAME&ip=$(hostname -I | awk '{print $1}')" $DDNS_ENDPOINT
echo "DDNS update request sent for $HOSTNAME"