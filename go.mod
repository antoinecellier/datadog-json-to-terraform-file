module github.com/antoinecellier/datadog-json-to-terraform-file

go 1.14

require (
	github.com/juliogreff/datadog-to-terraform v0.0.0-20211210143330-1d3ca0e41da6 // indirect
	github.com/rodaine/hclencoder v0.0.0-20200711024249-063a90d8fab0
)

replace github.com/rodaine/hclencoder v0.0.0-20200711024249-063a90d8fab0 => github.com/juliogreff/hclencoder v0.0.0-20201005124953-d013916b798c
