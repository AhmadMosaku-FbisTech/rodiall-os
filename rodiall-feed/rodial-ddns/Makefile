include $(TOPDIR)/rules.mk

PKG_NAME:=rodial-ddns
PKG_VERSION:=0.1
PKG_RELEASE:=1

include $(INCLUDE_DIR)/package.mk

define Package/rodial-ddns
  SECTION:=net
  CATEGORY:=Rodial
  TITLE:=Rodial DDNS client
  DEPENDS:=+curl
endef

define Package/rodial-ddns/install
	$(INSTALL_DIR) $(1)/usr/bin
	$(INSTALL_BIN) ./src/ddns.sh $(1)/usr/bin/rodial-ddns
	$(INSTALL_DIR) $(1)/etc/init.d
	$(INSTALL_BIN) ./files/etc/init.d/rodial-ddns $(1)/etc/init.d/rodial-ddns
endef

$(eval $(call BuildPackage,rodial-ddns))
