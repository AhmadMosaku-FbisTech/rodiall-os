include $(TOPDIR)/rules.mk

PKG_NAME:=rodial-proxy
PKG_VERSION:=0.1
PKG_RELEASE:=1

include $(INCLUDE_DIR)/package.mk

define Package/rodial-proxy
  SECTION:=net
  CATEGORY:=Rodial
  TITLE:=Rodial Proxy (placeholder)
  DEPENDS:=+libc
endef

define Package/rodial-proxy/description
  Minimal proxy placeholder for Rodial (SOCKS/HTTP).
endef

define Package/rodial-proxy/install
	$(INSTALL_DIR) $(1)/usr/bin
	$(INSTALL_BIN) $(PKG_BUILD_DIR)/proxy $(1)/usr/bin/rodial-proxy || true
	$(INSTALL_DIR) $(1)/etc/init.d
	$(INSTALL_BIN) ./files/etc/init.d/rodial-proxy $(1)/etc/init.d/rodial-proxy
endef

$(eval $(call BuildPackage,rodial-proxy))
