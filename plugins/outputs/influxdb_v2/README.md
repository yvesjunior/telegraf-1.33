# InfluxDB v2.x Output Plugin

This plugin writes metrics to a [InfluxDB v2.x][influxdb_v2] instance via HTTP.

⭐ Telegraf v1.8.0
🏷️ datastore
💻 all

[influxdb_v2]: https://docs.influxdata.com/influxdb/v2

## Global configuration options <!-- @/docs/includes/plugin_config.md -->

In addition to the plugin-specific configuration settings, plugins support
additional global and plugin configuration settings. These settings are used to
modify metrics, tags, and field or create aliases and configure ordering, etc.
See the [CONFIGURATION.md][CONFIGURATION.md] for more details.

[CONFIGURATION.md]: ../../../docs/CONFIGURATION.md#plugins

## Secret-store support

This plugin supports secrets from secret-stores for the `token` option.
See the [secret-store documentation][SECRETSTORE] for more details on how
to use them.

[SECRETSTORE]: ../../../docs/CONFIGURATION.md#secret-store-secrets

## Configuration

```toml @sample.conf
# Configuration for sending metrics to InfluxDB 2.0
[[outputs.influxdb_v2]]
  ## The URLs of the InfluxDB cluster nodes.
  ##
  ## Multiple URLs can be specified for a single cluster, only ONE of the
  ## urls will be written to each interval.
  ##   ex: urls = ["https://us-west-2-1.aws.cloud2.influxdata.com"]
  urls = ["http://127.0.0.1:8086"]

  ## Local address to bind when connecting to the server
  ## If empty or not set, the local address is automatically chosen.
  # local_address = ""

  ## Token for authentication.
  token = ""

  ## Organization is the name of the organization you wish to write to.
  organization = ""

  ## Destination bucket to write into.
  bucket = ""

  ## The value of this tag will be used to determine the bucket.  If this
  ## tag is not set the 'bucket' option is used as the default.
  # bucket_tag = ""

  ## If true, the bucket tag will not be added to the metric.
  # exclude_bucket_tag = false

  ## Timeout for HTTP messages.
  # timeout = "5s"

  ## Additional HTTP headers
  # http_headers = {"X-Special-Header" = "Special-Value"}

  ## HTTP Proxy override, if unset values the standard proxy environment
  ## variables are consulted to determine which proxy, if any, should be used.
  # http_proxy = "http://corporate.proxy:3128"

  ## HTTP User-Agent
  # user_agent = "telegraf"

  ## Content-Encoding for write request body, can be set to "gzip" to
  ## compress body or "identity" to apply no encoding.
  # content_encoding = "gzip"

  ## Enable or disable uint support for writing uints influxdb 2.0.
  # influx_uint_support = false

  ## When true, Telegraf will omit the timestamp on data to allow InfluxDB
  ## to set the timestamp of the data during ingestion. This is generally NOT
  ## what you want as it can lead to data points captured at different times
  ## getting omitted due to similar data.
  # influx_omit_timestamp = false

  ## HTTP/2 Timeouts
  ## The following values control the HTTP/2 client's timeouts. These settings
  ## are generally not required unless a user is seeing issues with client
  ## disconnects. If a user does see issues, then it is suggested to set these
  ## values to "15s" for ping timeout and "30s" for read idle timeout and
  ## retry.
  ##
  ## Note that the timer for read_idle_timeout begins at the end of the last
  ## successful write and not at the beginning of the next write.
  # ping_timeout = "0s"
  # read_idle_timeout = "0s"

  ## Optional TLS Config for use on HTTP connections.
  # tls_ca = "/etc/telegraf/ca.pem"
  # tls_cert = "/etc/telegraf/cert.pem"
  # tls_key = "/etc/telegraf/key.pem"
  ## Use TLS but skip chain & host verification
  # insecure_skip_verify = false

  ## Rate limits for sending data (disabled by default)
  ## Available, uncompressed payload size e.g. "5Mb"
  # rate_limit = "unlimited"
  ## Fixed time-window for the available payload size e.g. "5m"
  # rate_limit_period = "0s"
```

## Metrics

Reference the [influx serializer][] for details about metric production.

[influx serializer]: /plugins/serializers/influx/README.md#Metrics