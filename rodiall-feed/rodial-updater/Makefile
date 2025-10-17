include $(TOPDIR)/rules.mk

PKG_NAME:=rodial-updater
PKG_VERSION:=0.1
PKG_RELEASE:=1

include $(INCLUDE_DIR)/package.mk

define Package/rodial-updater
  SECTION:=base
  CATEGORY:=Rodial
  TITLE:=Rodial OTA updater
  DEPENDS:=+wget +ca-bundle
endef

define Package/rodial-updater/install
	$(INSTALL_DIR) $(1)/usr/bin
	$(INSTALL_BIN) ./src/updater.sh $(1)/usr/bin/rodial-updater
	$(INSTALL_DIR) $(1)/etc/init.d
	$(INSTALL_BIN) ./files/etc/init.d/rodial-updater $(1)/etc/init.d/rodial-updater
endef

$(eval $(call BuildPackage,rodial-updater))
define Package/rodial-updater/description
  Rodial OTA updater script.
endef