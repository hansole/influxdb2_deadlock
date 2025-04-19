# InfluxDB dummy reader and writer to trigger deadlock
We have observed a bucket become unresponsive. It looks like a
deadlock. This Go program is some dummy code that triggers this
deadlock quite "instantly".

## Setup
Create a bucket called `foo`.

```
term1$
export INFLUXDB_TOKEN=<TOKEN>
export INFLUXDB_ORG=<ORG>

term2$
export INFLUXDB_TOKEN=<TOKEN>
export INFLUXDB_ORG=<ORG>
```

## Compile
```
go build -o dummy_reader cmd/dummy_reader/main.go
go build -o dummy_writer cmd/dummy_writer/main.go
```

## Running
In one terminal start one of the programs. In another terminal start the other program. 



`term1$ ./dummy_writer`

`term2$ ./dummy_reader`

The deadlock happen within a few seconds.

InfluxDB need to be restarted to break the deadlock.

(This has been tested on Ubuntu 22.04)
