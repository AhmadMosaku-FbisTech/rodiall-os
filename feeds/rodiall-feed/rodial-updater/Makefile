#
# rodial-updater - OTA updater and firmware manager
#
include $(TOPDIR)/rules.mk

PKG_NAME:=rodial-updater
PKG_VERSION:=1.0
PKG_RELEASE:=1

include $(INCLUDE_DIR)/package.mk

define Package/$(PKG_NAME)
  SECTION:=utils
  CATEGORY:=Rodiall
  TITLE:=Rodiall OTA Updater
  DEPENDS:=+libopenssl +libcurl
endef

define Package/$(PKG_NAME)/description
  Handles firmware updates and signed image verification for Rodiall OS.
endef

define Build/Compile
	$(MAKE) -C src
endef

define Package/$(PKG_NAME)/install
	$(INSTALL_DIR) $(1)/usr/sbin
	$(INSTALL_BIN) ./src/rodial-updater $(1)/usr/sbin/
	$(INSTALL_DIR) $(1)/etc/init.d
	$(INSTALL_BIN) ./files/etc/init.d/rodial-updater $(1)/etc/init.d/rodial-updater
endef

$(eval $(call BuildPackage,rodial-updater))
