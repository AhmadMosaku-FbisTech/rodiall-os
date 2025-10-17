# Rodiall OpenWRT Feed

This feed contains OpenWRT-style packages for the Rodiall project:
- luci-app-rodial
- rodial-agent
- rodial-proxy
- rodial-dt
- rodial-updater
- rodial-ddns
- rodial-utils

How to use:
1. Put this directory in your repo at `feeds/rodiall-feed`.
2. In `openwrt/feeds.conf`, add:
   `src-link rodiall ../feeds/rodiall-feed`
3. From `openwrt/` run:
   ```bash
   ./scripts/feeds update rodiall
   ./scripts/feeds install -a -p rodiall
   make menuconfig   # select packages under 'rodiall' or refresh package list
   make -j$(nproc)
