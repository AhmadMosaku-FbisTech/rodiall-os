#
# rodial-utils - Shared CLI tools and libraries
#
include $(TOPDIR)/rules.mk

PKG_NAME:=rodial-utils
PKG_VERSION:=1.0
PKG_RELEASE:=1

include $(INCLUDE_DIR)/package.mk

define Package/$(PKG_NAME)
  SECTION:=utils
  CATEGORY:=Rodiall
  TITLE:=Rodiall Utilities
  DEPENDS:=+libubus +libuci
endef

define Package/$(PKG_NAME)/description
  Common utilities and helper binaries for Rodiall OS services.
endef

define Build/Compile
	$(MAKE) -C src
endef

define Package/$(PKG_NAME)/install
	$(INSTALL_DIR) $(1)/usr/bin
	$(INSTALL_BIN) ./src/rodial-cli $(1)/usr/bin/
endef

$(eval $(call BuildPackage,rodial-utils))
