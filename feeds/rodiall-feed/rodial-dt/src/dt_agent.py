#!/usr/bin/env python3
import time, json, os
def main():
    # Very small DT agent that prints telemetry periodically
    while True:
        telemetry = {
            "uptime": int(time.time()),
            "device": os.uname().nodename
        }
        print("DT-TELEMETRY:", json.dumps(telemetry))
        time.sleep(30)

if __name__=="__main__":
    main()
