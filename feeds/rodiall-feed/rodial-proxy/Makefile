#
# rodial-proxy - SOCKS/HTTP proxy service
#
include $(TOPDIR)/rules.mk

PKG_NAME:=rodial-proxy
PKG_VERSION:=1.0
PKG_RELEASE:=1

include $(INCLUDE_DIR)/package.mk

define Package/$(PKG_NAME)
  SECTION:=net
  CATEGORY:=Rodiall
  TITLE:=Rodiall Proxy Service
  DEPENDS:=+libpthread +libopenssl
endef

define Package/$(PKG_NAME)/description
  SOCKS/HTTP proxy component for Rodiall OS networking.
endef

define Build/Compile
	$(MAKE) -C src
endef

define Package/$(PKG_NAME)/install
	$(INSTALL_DIR) $(1)/usr/sbin
	$(INSTALL_BIN) ./src/rodial-proxy $(1)/usr/sbin/
	$(INSTALL_DIR) $(1)/etc/init.d
	$(INSTALL_BIN) ./files/etc/init.d/rodial-proxy $(1)/etc/init.d/rodial-proxy
endef

$(eval $(call BuildPackage,rodial-proxy))
