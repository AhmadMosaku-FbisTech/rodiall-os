include $(TOPDIR)/rules.mk

PKG_NAME:=rodial-dt
PKG_VERSION:=0.1
PKG_RELEASE:=1

include $(INCLUDE_DIR)/package.mk

define Package/rodial-dt
  SECTION:=utils
  CATEGORY:=Rodial
  TITLE:=Rodial Digital Twin agent
  DEPENDS:=+python3
endef

define Package/rodial-dt/description
  Digital Twin lightweight agent (publishes telemetry via MQTT).
endef

define Package/rodial-dt/install
	$(INSTALL_DIR) $(1)/usr/bin
	$(INSTALL_BIN) ./src/dt_agent.py $(1)/usr/bin/rodial-dt
	$(INSTALL_DIR) $(1)/etc/init.d
	$(INSTALL_BIN) ./files/etc/init.d/rodial-dt $(1)/etc/init.d/rodial-dt
	$(INSTALL_DIR) $(1)/etc/rodial
	$(INSTALL_DATA) ./config/schema.json $(1)/etc/rodial/schema.json
endef

$(eval $(call BuildPackage,rodial-dt))
