#!/bin/sh
# Very small placeholder OTA updater
FIRMWARE_URL=${1:-"https://example.com/rodiall/latest.img"}
TMP=/tmp/rodial_update.img

echo "Downloading $FIRMWARE_URL..."
/usr/bin/wget -O $TMP $FIRMWARE_URL || exit 1
echo "Downloaded to $TMP (not flashing in this skeleton)"
# Real implementation: verify signature, call sysupgrade --force $TMP
echo "Update process completed (not really, this is a placeholder)."
exit 0
