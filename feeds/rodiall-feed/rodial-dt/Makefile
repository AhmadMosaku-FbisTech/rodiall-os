#
# rodial-dt - Digital Twin daemon
#
include $(TOPDIR)/rules.mk

PKG_NAME:=rodial-dt
PKG_VERSION:=1.0
PKG_RELEASE:=1

include $(INCLUDE_DIR)/package.mk

define Package/$(PKG_NAME)
  SECTION:=utils
  CATEGORY:=Rodiall
  TITLE:=Rodiall Digital Twin Daemon
  DEPENDS:=+libjson-c +libubus +libuci
endef

define Package/$(PKG_NAME)/description
  Collects device telemetry and syncs to cloud twin model.
endef

define Build/Compile
	$(MAKE) -C src
endef

define Package/$(PKG_NAME)/install
	$(INSTALL_DIR) $(1)/usr/sbin
	$(INSTALL_BIN) ./src/rodial-dt $(1)/usr/sbin/
	$(INSTALL_DIR) $(1)/etc/init.d
	$(INSTALL_BIN) ./files/etc/init.d/rodial-dt $(1)/etc/init.d/rodial-dt
endef

$(eval $(call BuildPackage,rodial-dt))
