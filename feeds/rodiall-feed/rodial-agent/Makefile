#
# rodial-agent - Core device agent (handles registration, tunnels, and telemetry)
#
include $(TOPDIR)/rules.mk

PKG_NAME:=rodial-agent
PKG_VERSION:=1.0
PKG_RELEASE:=1

PKG_BUILD_DIR := $(BUILD_DIR)/$(PKG_NAME)

include $(INCLUDE_DIR)/package.mk

define Package/$(PKG_NAME)
  SECTION:=net
  CATEGORY:=Rodiall
  TITLE:=Rodiall Agent
  DEPENDS:=+libpthread +libuci +libubus +libjson-c
endef

define Package/$(PKG_NAME)/description
  Core Rodiall device agent. Handles registration with the cloud and maintains tunnels.
endef

define Build/Prepare
	mkdir -p $(PKG_BUILD_DIR)
	$(CP) ./src/* $(PKG_BUILD_DIR)/
endef

define Build/Compile
	$(MAKE) -C $(PKG_BUILD_DIR) CC="$(TARGET_CC)" CFLAGS="$(TARGET_CFLAGS)"
endef

define Package/$(PKG_NAME)/install
	$(INSTALL_DIR) $(1)/usr/sbin
	$(INSTALL_BIN) $(PKG_BUILD_DIR)/rodial-agent $(1)/usr/sbin/
	$(INSTALL_DIR) $(1)/etc/init.d
	$(INSTALL_BIN) ./files/etc/init.d/rodial-agent $(1)/etc/init.d/rodial-agent
endef

$(eval $(call BuildPackage,rodial-agent))
