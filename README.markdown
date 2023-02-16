This repo is a fork of https://github.com/juliogreff/datadog-to-terraform.

# datadog to terraform

A simple utility to format a Datadog JSON export into a Terraform file.

Commande to run:
```
go run . monitor "my_arlerting" test.json
```

# Getting started


-  Export a monitor in Datadog

![Export a monitor](export_monitor.png)

- Copy paste the export in a json file and run the `main.go` file with appropriate arguments
```
go run . [dashboard|monitor] [ressource_name] [json_path]
```

- The output
```
resource "datadog_monitor" "ressource_name" {
  type  = "query alert"
  query = "sum(last_1d):(sum:trace.rack.request.errors{env:production,resource_name:paymentcontroller*}.as_count() * 100) / sum:trace.rack.request.hits{env:production,resource_name:paymentcontroller*}.as_count() > 0.05"
  name  = "Third party error rate"

  message = "@slack-my_team
 
Third party error rate reached {{eval "value*100"}}%."

  tags = [
    "team",
  ]

  no_data_timeframe = 0
  notify_audit      = false
  notify_no_data    = false
  renotify_interval = 0

  monitor_thresholds {
    critical = "0.05"
    warning  = "0.005"
  }

  include_tags        = false
  require_full_window = false
}
```
