#
# rodial-agent - device agent
#
include $(TOPDIR)/rules.mk

PKG_NAME:=rodial-agent
PKG_VERSION:=0.1
PKG_RELEASE:=1

PKG_BUILD_DIR:=$(BUILD_DIR)/$(PKG_NAME)
PKG_SOURCE_PROTO:=git
PKG_SOURCE_URL:=file://$(TOPDIR)/feeds/rodiall-feed/rodial-agent/src
PKG_SOURCE_VERSION:=local

include $(INCLUDE_DIR)/package.mk

define Package/rodial-agent
  SECTION:=utils
  CATEGORY:=Rodial
  TITLE:=Rodial device agent
  DEPENDS:=+libc +ca-bundle
endef

define Package/rodial-agent/description
  The Rodial agent maintains connection to FBIS and manages tunnels.
endef

define Build/Prepare
	mkdir -p $(PKG_BUILD_DIR)
endef

define Build/Compile
	# For the feed skeleton, we'll copy prebuilt binary or Go source.
	# If using Go, cross-build instructions belong here.
	true
endef

define Package/rodial-agent/install
	$(INSTALL_DIR) $(1)/usr/bin
	$(INSTALL_BIN) $(PKG_BUILD_DIR)/rodial-agent $(1)/usr/bin/rodial-agent || true
	$(INSTALL_DIR) $(1)/etc/init.d
	$(INSTALL_BIN) ./files/etc/init.d/rodial-agent $(1)/etc/init.d/rodial-agent
endef

$(eval $(call BuildPackage,rodial-agent))
