# HDDtemp Input Plugin

This plugin reads data from hddtemp daemon.

Hddtemp should be installed and its daemon running.

## OS Support & Alternatives

This plugin depends on the availability of the `hddtemp` binary. The upstream
project is not active and Debian made the decision to remove it in Debian
Bookworm. This means the rest of the Debian ecosystem no longer has this binary
in later releases, like Ubuntu 22.04.

As an alternative consider using the [`smartctl` plugin]. This parses the full
JSON output from `smartctl`, which includes temperature data, in addition to
much more data about devices in a system.

[`smartctl` plugin]: ../smartctl/README.md

## Global configuration options <!-- @/docs/includes/plugin_config.md -->

In addition to the plugin-specific configuration settings, plugins support
additional global and plugin configuration settings. These settings are used to
modify metrics, tags, and field or create aliases and configure ordering, etc.
See the [CONFIGURATION.md][CONFIGURATION.md] for more details.

[CONFIGURATION.md]: ../../../docs/CONFIGURATION.md#plugins

## Configuration

```toml @sample.conf
# Monitor disks' temperatures using hddtemp
[[inputs.hddtemp]]
  ## By default, telegraf gathers temps data from all disks detected by the
  ## hddtemp.
  ##
  ## Only collect temps from the selected disks.
  ##
  ## A * as the device name will return the temperature values of all disks.
  ##
  # address = "127.0.0.1:7634"
  # devices = ["sda", "*"]
```

## Metrics

- hddtemp
  - tags:
    - device
    - model
    - unit
    - status
    - source
  - fields:
    - temperature

## Example Output

```text
hddtemp,source=server1,unit=C,status=,device=sdb,model=WDC\ WD740GD-00FLA1 temperature=43i 1481655647000000000
hddtemp,device=sdc,model=SAMSUNG\ HD103UI,unit=C,source=server1,status= temperature=38i 148165564700000000
hddtemp,device=sdd,model=SAMSUNG\ HD103UI,unit=C,source=server1,status= temperature=36i 1481655647000000000
```