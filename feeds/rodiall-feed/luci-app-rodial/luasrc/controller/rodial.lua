module("luci.controller.rodial", package.seeall)

function index()
  entry({"admin", "services", "rodial"}, firstchild(), "Rodial", 60)
  entry({"admin", "services", "rodial", "status"}, call("action_status"))
end

function action_status()
  luci.http.prepare_content("application/json")
  luci.http.write_json({status="ok", agent="rodial-agent", uptime=os.time()})
end
