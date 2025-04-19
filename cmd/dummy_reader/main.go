package main

import (
	"context"
	"fmt"
	"log"
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func main() {
	token := os.Getenv("INFLUXDB_TOKEN")
	url := "http://localhost:8086"
	client := influxdb2.NewClient(url, token)
	org := os.Getenv("INFLUXDB_ORG")
	queryAPI := client.QueryAPI(org)
	query := `from(bucket: "foo")
              |> range(start: -30m)
              |> filter(fn: (r) => r.tagKey_10 =~ /foobar/)`

	// |> filter(fn: (r) => r.tagKey_10 == "a_location_at_181")`
	// |> filter(fn: (r) => r.tagKey_10 =~ /a_loc/)`

	for {
		results, err := queryAPI.Query(context.Background(), query)
		if err != nil {
			log.Fatal(err)
		}
		for results.Next() {
			fmt.Println(results.Record())
		}
		if err := results.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
