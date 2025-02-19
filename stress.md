

### **Telegraf Stress Testing with k6**  
This guide will help you **stress test Telegraf** using [`k6`](https://k6.io/), a powerful load testing tool. The goal is to send a large number of metrics to **Telegraf’s HTTP listener** and measure its performance.

---

## **1. Set Up Telegraf for HTTP Input**  
Ensure Telegraf can receive metrics via HTTP by enabling `inputs.http_listener_v2` in your config file (`/etc/telegraf/telegraf.conf`):

```toml
[[inputs.http_listener_v2]]
  service_address = ":8186"
  data_format = "influx"
```
- This opens **port 8186** for Telegraf to receive metrics.  
- `data_format = "influx"` ensures data is formatted correctly.

Restart Telegraf:  
```bash
sudo systemctl restart telegraf
```

---

## **2. Install k6**
If you don’t have `k6`, install it:

- **Ubuntu/Debian**  
  ```bash
  sudo apt install k6
  ```
- **MacOS (Homebrew)**  
  ```bash
  brew install k6
  ```
- **Windows**  
  Download from [k6.io](https://k6.io/docs/getting-started/installation/).

---

## **3. Create a k6 Test Script**
Create a test script (`telegraf_stress_test.js`) to simulate **high metric ingestion**:

```javascript
import http from "k6/http";
import { sleep } from "k6";

export let options = {
    vus: 100,  // Number of virtual users (simultaneous connections)
    duration: "30s",  // Test duration
};

export default function () {
    let payload = "test_metric,host=server01 value=42 " + Date.now() * 1000000;
    let params = { headers: { "Content-Type": "text/plain" } };

    http.post("http://localhost:8186/metrics", payload, params);
    sleep(0.01);  // Short delay to simulate real-world load
}
```

This script:
- Simulates **100 concurrent users**.
- Sends a metric (`test_metric,host=server01 value=42`) to Telegraf’s HTTP input.
- Runs for **30 seconds**.
- Uses **10ms sleep** to control load.

---

## **4. Run the k6 Stress Test**
Execute the test:  
```bash
k6 run telegraf_stress_test.js
```

Output Example:
```
  checks..............: 100% ✓  (5000/5000)
  http_reqs..........: 5000    (166.667/s)
  http_req_failed....: 0% ✓  (0/5000)
  http_req_duration..: avg=50ms, min=10ms, max=100ms
```

**What to check:**
- `http_req_failed`: Should be **0%** (if nonzero, Telegraf is dropping requests).
- `http_reqs`: Total HTTP requests sent (higher is better).
- `http_req_duration`: Response time from Telegraf.

---

## **5. Monitor Telegraf Performance**
While running `k6`, **monitor system performance** in another terminal:

```bash
htop       # CPU/memory usage
iotop      # Disk I/O
ifstat     # Network traffic
```

Check **Telegraf’s internal metrics** by enabling this plugin:
```toml
[[inputs.internal]]
  collect_memstats = true
```
Then inspect:
```bash
telegraf --test --config /etc/telegraf/telegraf.conf
```

---

## **6. Optimize Telegraf Performance**
If Telegraf slows down, modify its settings:
1. **Increase Buffer Size** (prevents dropped metrics):  
   ```toml
   metric_buffer_limit = 20000
   ```
2. **Increase Batch Size** (reduces write frequency):  
   ```toml
   batch_size = 5000
   flush_interval = "5s"
   ```
3. **Reduce HTTP Overhead** by increasing the request timeout:  
   ```toml
   read_timeout = "10s"
   write_timeout = "10s"
   ```

---

## **7. Analyze Results**
After the test, analyze:
- **Telegraf logs** (`/var/log/telegraf/telegraf.log`):
  ```bash
  grep "error" /var/log/telegraf/telegraf.log
  ```
- **InfluxDB or Prometheus ingestion rate**:
  ```bash
  curl -G http://localhost:8086/query --data-urlencode "q=SELECT * FROM test_metric LIMIT 10"
  ```

---

# #############################
# #############################
# #############################
# #############################


### **Telegraf Stress Testing Guide**
Stress testing **Telegraf** ensures it can handle high loads without performance degradation. Here’s how you can effectively stress test Telegraf:

---

## **1. Define Your Stress Test Goals**
Before running a stress test, determine:
- Maximum number of metrics per second Telegraf should handle.
- The impact of increasing input/output plugins.
- CPU, memory, disk, and network utilization thresholds.
- How efficiently Telegraf processes and sends data to the backend.

---

## **2. Generate High Metric Volume**
### **Method 1: Using `inputs.exec` to Simulate High Load**
You can generate a large number of metrics using a simple script:
```bash
while true; do
  echo "test_metric,value=$RANDOM count=$((RANDOM % 1000)) $(date +%s%N)"
done
```
Then configure Telegraf to ingest this data using:
```toml
[[inputs.exec]]
  commands = ["/path/to/your/script.sh"]
  interval = "1s"
  data_format = "influx"
```
This will flood Telegraf with random metrics every second.

---

### **Method 2: Using `k6` for HTTP Stress Testing**
If Telegraf is ingesting metrics via HTTP (e.g., `inputs.http_listener_v2`), use [`k6`](https://k6.io/) to stress test:
1. Install k6:
   ```bash
   sudo apt install k6  # For Ubuntu
   ```
2. Write a k6 test script (e.g., `stress_test.js`):
   ```javascript
   import http from "k6/http";
   import { sleep } from "k6";

   export default function () {
       http.post("http://localhost:8186/metrics", "test_metric,value=42 count=1");
       sleep(0.01);
   }
   ```
3. Run the stress test:
   ```bash
   k6 run --vus 100 --duration 30s stress_test.js
   ```
   This simulates 100 virtual users sending metrics for 30 seconds.

---

## **3. Monitor Performance**
### **Telegraf Internal Metrics**
Enable Telegraf’s internal metrics:
```toml
[[inputs.internal]]
  collect_memstats = true
```
Monitor:
- `telegraf_agent_gather_seconds`: Collection time per cycle.
- `telegraf_outputs_write_seconds`: Time to send metrics to the backend.
- `telegraf_metrics_dropped_total`: Number of dropped metrics.

### **System Performance**
Run these in parallel:
```bash
htop       # Monitor CPU and memory
iotop      # Check disk usage
ifstat     # Monitor network traffic
```
Check if Telegraf consumes excessive resources or drops metrics.

---

## **4. Optimize Telegraf Performance**
If performance degrades, try:
- Increasing `batch_size` and `metric_buffer_limit` in output plugins:
  ```toml
  [[outputs.influxdb]]
    batch_size = 5000
    metric_buffer_limit = 20000
  ```
- Using parallel processing:
  ```toml
  [agent]
    metric_batch_size = 5000
    metric_buffer_limit = 10000
    flush_interval = "5s"
  ```
- Reducing unnecessary metrics with `namepass` or `namedrop`.

---

## **5. Analyze and Tune**
- Identify slowest plugins using:
  ```bash
  telegraf --test --config /etc/telegraf/telegraf.conf
  ```
- Profile CPU and memory usage:
  ```bash
  TELEGRAF_PROFILE=6060 telegraf
  go tool pprof http://localhost:6060/debug/pprof/profile
  ```
- Review backend ingestion rates (e.g., InfluxDB or Prometheus).

---


# Global Agent Configuration
[agent]
  interval = "10s"           # Data collection interval
  round_interval = true
  metric_batch_size = 1000
  metric_buffer_limit = 10000
  collection_jitter = "0s"
  flush_interval = "10s"
  flush_jitter = "0s"
  precision = ""
  debug = false
  quiet = false
  hostname = ""              # Set to override the hostname

# Input Plugins for Host Metrics
[[inputs.cpu]]
  percpu = true
  totalcpu = true
  collect_cpu_time = false
  report_active = true

[[inputs.mem]]

[[inputs.swap]]

[[inputs.disk]]
  ignore_fs = ["tmpfs", "devtmpfs", "overlay"]

[[inputs.diskio]]

[[inputs.net]]
  interfaces = ["eth0", "eth1"]  # Specify interfaces or leave empty for all

[[inputs.netstat]]

[[inputs.system]]

[[inputs.processes]]

[[inputs.kernel]]

# Output Plugin (Example: InfluxDB)
[[outputs.influxdb]]
  urls = ["http://localhost:8086"]
  database = "telegraf"
  username = "telegraf"
  password = "telegraf_password"
  retention_policy = ""
  timeout = "5s"
  precision = "s"


