#
# rodial-ddns - Dynamic DNS updater
#
include $(TOPDIR)/rules.mk

PKG_NAME:=rodial-ddns
PKG_VERSION:=1.0
PKG_RELEASE:=1

include $(INCLUDE_DIR)/package.mk

define Package/$(PKG_NAME)
  SECTION:=net
  CATEGORY:=Rodiall
  TITLE:=Rodiall DDNS Client
  DEPENDS:=+libcurl +libubox
endef

define Package/$(PKG_NAME)/description
  Reports the router’s current IP address to the Rodiall Cloud.
endef

define Build/Compile
	$(MAKE) -C src
endef

define Package/$(PKG_NAME)/install
	$(INSTALL_DIR) $(1)/usr/sbin
	$(INSTALL_BIN) ./src/rodial-ddns $(1)/usr/sbin/
	$(INSTALL_DIR) $(1)/etc/init.d
	$(INSTALL_BIN) ./files/etc/init.d/rodial-ddns $(1)/etc/init.d/rodial-ddns
endef

$(eval $(call BuildPackage,rodial-ddns))
