
---

## Package skeletons & files

> Note: These Makefiles are OpenWRT package skeletons (they follow the `include $(TOPDIR)/rules.mk` convention). They are minimal — expand as needed (add dependencies, CFLAGS, Go packaging options, compiling steps, install lines).

---

### `luci-app-rodial/Makefile`
```make
#
# luci-app-rodial - LuCI Web UI for Rodial
#
include $(TOPDIR)/rules.mk

LUCI_TITLE:=LuCI - Rodial
LUCI_DEPENDS:=+rpcd +uhttpd +luci-base

PKG_NAME:=luci-app-rodial
PKG_VERSION:=1.0
PKG_RELEASE:=1

include $(INCLUDE_DIR)/package.mk

define Package/$(PKG_NAME)
  SECTION:=luci
  CATEGORY:=LuCI
  SUBMENU:=3. Applications
  TITLE:=Rodial Web UI
  DEPENDS:=$(LUCI_DEPENDS)
endef

define Package/$(PKG_NAME)/description
  LuCI web interface for the Rodial agent and configuration.
endef

define Build/Prepare
	mkdir -p $(PKG_BUILD_DIR)
endef

define Package/$(PKG_NAME)/install
	$(INSTALL_DIR) $(1)/www/rodial
	$(CP) -a $(CURDIR)/htdocs/* $(1)/www/rodial/
	$(INSTALL_DIR) $(1)/usr/lib/lua/luci/controller
	$(INSTALL_DATA) $(CURDIR)/luasrc/controller/rodial.lua $(1)/usr/lib/lua/luci/controller/rodial.lua
	$(INSTALL_DIR) $(1)/etc/config
	$(INSTALL_DATA) $(CURDIR)/root/etc/config/rodial $(1)/etc/config/rodial
endef

$(eval $(call BuildPackage,luci-app-rodial))
