
# 🛰️ Rodiall OS

**A Smart, Digital-Twin-Enabled Router Firmware for Next-Generation Network Access**

---

## 📖 Overview

**Rodiall OS** is a custom firmware inspired by **OpenWrt**, designed for routers, gateways, and edge devices that interact with a **Digital Twin Network**, **VPN Mesh**, and **centralized control server**.
It combines real-world network infrastructure with virtual twin synchronization — allowing internet connectivity, control, and full 5G-level performance *even outside physical router range*, via a VPN-based mobile application called **Rodial**.

Rodiall OS transforms routers into **intelligent network nodes** that can broadcast, slice, and tunnel internet access through a hybrid virtual layer — enabling remote full-strength connectivity, low latency, and adaptive VPN routing.

---

## 🧩 Key Features

| Feature                                   | Description                                                                                 |
| ----------------------------------------- | ------------------------------------------------------------------------------------------- |
| **OpenWrt Core**                          | Built atop the robust OpenWrt platform for modularity and hardware support                  |
| **Digital Twin Daemon (`rodial-dt`)**     | Mirrors physical router network states to a virtual twin for synchronization with the cloud |
| **Rodial VPN Integration**                | Connects the Rodial mobile app and VPN backend for virtualized internet access              |
| **Dynamic Proxy Layer (`rodial-proxy`)**  | Provides SOCKS5/HTTP proxy tunnels to route data securely                                   |
| **Secure OTA Updater (`rodial-updater`)** | Handles signed firmware distribution and over-the-air updates                               |
| **Custom Web UI (`luci-app-rodial`)**     | LuCI-based dashboard for monitoring and managing Rodial services                            |
| **DDNS Agent (`rodial-ddns`)**            | Registers devices dynamically to central servers with rotating endpoints                    |
| **CLI Utilities (`rodial-utils`)**        | Handy command-line tools for diagnostics, configuration, and network slicing                |

---

## 🏗️ Repository Structure

```
rodiall-os/
│
├── openwrt/                            # Core OpenWRT source (git submodule or clone)
│   ├── feeds.conf                      # Feed definitions (include Rodiall feed)
│   ├── target/                         # Platform targets (x86, arm, etc.)
│   ├── package/                        # Base OpenWRT packages
│   ├── include/                        # Common build includes
│   ├── files/                          # Overlay rootfs (e.g., /etc/banner, /etc/config)
│   └── README.md
│
├── feeds/                              # Custom external feeds
│   ├── rodiall-feed/                   # Primary Rodiall package feed
│   │   ├── luci-app-rodial/            # Custom LuCI WebUI for Rodiall control/status
│   │   │   ├── luasrc/
│   │   │   ├── htdocs/
│   │   │   ├── Makefile
│   │   │   └── root/etc/config/rodial
│   │   │
│   │   ├── rodial-agent/               # Core device agent (registers router, manages tunnels)
│   │   │   ├── src/
│   │   │   │   ├── main.c              # or main.go / main.rs
│   │   │   │   └── utils/
│   │   │   ├── files/etc/init.d/rodial-agent
│   │   │   ├── Makefile
│   │   │   └── README.md
│   │   │
│   │   ├── rodial-proxy/               # Proxy service (SOCKS/HTTP tunnel)
│   │   │   ├── src/
│   │   │   ├── Makefile
│   │   │   └── files/etc/init.d/rodial-proxy
│   │   │
│   │   ├── rodial-dt/                  # Digital Twin daemon
│   │   │   ├── src/
│   │   │   ├── config/schema.json
│   │   │   ├── Makefile
│   │   │   └── files/etc/init.d/rodial-dt
│   │   │
│   │   ├── rodial-updater/             # OTA updater and signed firmware manager
│   │   │   ├── src/
│   │   │   ├── Makefile
│   │   │   └── files/etc/init.d/rodial-updater
│   │   │
│   │   ├── rodial-ddns/                # DDNS client for router registration
│   │   │   ├── src/
│   │   │   ├── Makefile
│   │   │   └── files/etc/init.d/rodial-ddns
│   │   │
│   │   └── rodial-utils/               # Common helper libraries and CLI tools
│   │       ├── src/
│   │       ├── Makefile
│   │       └── README.md
│   │
│   └── others/                         # Any 3rd-party or external feed
│
├── digital-twin/                       # Cloud digital twin SDK and schema definitions
│   ├── backend-api/                    # Go/Python backend for FBIS twin services
│   ├── grpc-protos/                    # Proto definitions for agent <-> twin
│   ├── mqtt-schema/                    # JSON schemas for telemetry topics
│   └── dashboard-ui/                   # Admin dashboard (React/Vue)
│
├── mobile-app/                         # "Rodial" VPN mobile client app
│   ├── android/
│   ├── ios/
│   ├── shared-core/                    # Shared logic, encryption, session handling
│   └── README.md
│
├── build/                              # Build tools and CI/CD pipeline
│   ├── configs/
│   │   ├── x86_64.config               # For PC/router-class hardware
│   │   ├── armv8.config                # For ARM boards (MT7986, BPi-R3, etc.)
│   │   ├── riscv.config                # For RISC-V targets
│   │   └── default.config
│   │
│   ├── scripts/
│   │   ├── build_iso.sh                # Build x86_64 ISO (PC router)
│   │   ├── build_arm_img.sh            # Build ARM firmware image (.img.gz)
│   │   ├── build_riscv_itb.sh          # Build RISC-V firmware (.itb)
│   │   ├── sign_firmware.sh            # Sign builds with private key
│   │   ├── verify_firmware.sh          # Verify firmware signature
│   │   ├── publish_release.sh          # Push to GitHub/S3 releases
│   │   └── ci_hooks.sh                 # CI hooks
│   │
│   └── pipeline.yml                    # CI/CD definition (GitHub Actions, GitLab CI)
│
├── docs/                               # Developer and integration documentation
│   ├── architecture-overview.md        # Explains control/data plane flow
│   ├── digital-twin-spec.md            # JSON schema + MQTT topics
│   ├── vpn-handshake-sequence.md       # Mapped to the image you uploaded
│   ├── carrier-integration.md          # Explains APN, zero-rating, network slicing
│   ├── firmware-build-guide.md         # Step-by-step OpenWRT build & signing
│   ├── mobile-app-protocol.md          # Handshake, packet format
│   └── security-model.md               # Device identity, PKI, signed updates
│
├── ci-keys/                            # Firmware signing keys (private key never committed)
│   ├── rodial.pub
│   └── rodial.key (excluded via .gitignore)
│
├── outputs/                            # Auto-generated build artifacts
│   ├── rodiall-os-x86_64.iso
│   ├── rodiall-armv8.img.gz
│   ├── rodiall-riscv.itb
│   ├── checksums/sha256sums.txt
│   └── latest/                         # Symlink to latest release
│
├── LICENSE
├── README.md
└── .gitignore
```

---

## ⚙️ Build Instructions

1. **Clone OpenWrt source:**

   ```bash
   git clone https://github.com/openwrt/openwrt.git openwrt
   cd openwrt
   ```

2. **Link the Rodiall feed:**

   ```bash
   echo "src-link rodiall-feed ../feeds/rodiall-feed" >> feeds.conf
   ./scripts/feeds update -a
   ./scripts/feeds install -a
   ```

3. **Select your target configuration:**

   ```bash
   cp ../build/configs/x86_64.config .config
   make defconfig
   ```

4. **Build the firmware:**

   ```bash
   make -j$(nproc)
   ```

5. **Locate output images:**

   ```
   bin/targets/x86/64/rodiall-os-x86_64.iso
   ```

---

## 🌐 Rodial VPN & Digital Twin

Rodiall OS works in tandem with the **Rodial mobile VPN app** and a **cloud Digital Twin network**:

* The **physical router** broadcasts limited seed data (≈5 MB initialization packets).
* The **digital twin** mirrors network performance parameters (latency, throughput, QoS).
* The **Rodial app** receives full-strength, low-latency internet access through virtual slicing.
* VPN tunnels synchronize telemetry, authentication, and usage states between the physical and virtual networks.

---

## 🔐 Security Architecture

* **AES-256 & Curve25519** encryption for tunnel sessions
* **TLS-based control channels** between router ↔ digital twin ↔ mobile
* **Signed firmware updates** validated by `rodial-updater`
* **Firewall hooks** integrated with OpenWrt’s `netfilter` and `ubus`

---

## 🚀 Build Targets

| Target   | Output    | Description                                    |
| -------- | --------- | ---------------------------------------------- |
| `x86_64` | `.iso`    | PC or x86 router firmware                      |
| `armv8`  | `.img.gz` | ARM64 embedded boards (e.g. MT7986, Banana Pi) |
| `riscv`  | `.itb`    | RISC-V dev boards and SoCs                     |

---

## 🧠 Development Notes

* Each package (`rodial-agent`, `rodial-proxy`, etc.) is self-contained in `feeds/rodiall-feed/`.
* CI/CD builds are automated using `build/pipeline.yml`.
* The LuCI WebUI (`luci-app-rodial`) offers real-time control and visualization of all services.

---

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/awesome`)
3. Commit changes (`git commit -m 'Add new feature'`)
4. Push and open a Pull Request

---

## 🧩 License

Rodiall OS © 2025 — Licensed under the **GNU GPL v3**.
Includes components derived from **OpenWrt** and other open-source projects.

---

## 🛰️ Credits

* **Lead Firmware Engineer:** Malik of Codes
* **Core Technologies:** OpenWrt, LuCI, VPN Mesh, Digital Twin Synchronization
* **Special Thanks:** OpenWrt community & embedded systems contributors

---

