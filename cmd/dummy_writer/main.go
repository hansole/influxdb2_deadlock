package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

func main() {
	token := os.Getenv("INFLUXDB_TOKEN")
	url := "http://localhost:8086"
	client := influxdb2.NewClient(url, token)
	org := os.Getenv("INFLUXDB_ORG")
	bucket := "foo"
	writeAPI := client.WriteAPIBlocking(org, bucket)
	count := 0
	for {
		tags := make(map[string]string)
		fields := make(map[string]any)
		j := rand.Intn(30) + 10 // More measueremets hits the deadlock faster
		measurement := fmt.Sprintf("measurement_%d", j)
		for n := 10; n < 30; n++ { // More tags hits the deadlock faster
			key := fmt.Sprintf("tagKey_%d", n)
			tags[key] = fmt.Sprintf("a_location_at_%d", int64(rand.Intn(200))+100)
		}
		fields["fieldKey"] = int64(rand.Intn(100))
		point := write.NewPoint(measurement, tags, fields, time.Now())
		if err := writeAPI.WritePoint(context.Background(), point); err != nil {
			log.Fatal(err)
		}
		count += 1
		if count > 1000 {
			fmt.Printf(".")
			count = 0
		}
	}
}
