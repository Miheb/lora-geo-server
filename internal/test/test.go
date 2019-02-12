//go:generate protoc -I=. -I=../../vendor -I=$GOPATH/src --go_out=plugins=grpc:. models.proto

package test

import (
	"context"
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/brocaar/lora-geo-server/internal/backends/collos"
	"github.com/brocaar/lora-geo-server/internal/backends/willy"
	"github.com/brocaar/lora-geo-server/internal/config"
	"github.com/brocaar/loraserver/api/geo"
)

// ResolveTDOA runs the given Resolve TDOA test-suite.
func ResolveTDOA(ts ResolveTDOATestSuite) error {
	var backend geo.GeolocationServerServiceServer
	switch config.C.GeoServer.Backend.Name {
    case "willy":
        backend = willy.NewAPIWilly(config.C.GeoServer.Backend.Willy)
	default:
        backend = collos.NewAPICollos(config.C.GeoServer.Backend.Collos)
    }
	
	w := csv.NewWriter(os.Stdout)
	if err := w.Write([]string{
		"id",
		"exp_lat",
		"exp_long",
		"exp_alt",
		"collos_out_lat",
		"collos_out_long",
		"collos_out_alt",
		"collos_diff_meters",
	}); err != nil {
		log.Fatal(err)
	}

	for i, t := range ts.Tests {
		res, err := backend.ResolveTDOA(context.Background(), t.Request)
		if err != nil {
			log.Printf("%d - %s", i, err)
			continue
		}

		if res.Result == nil {
			log.Printf("%d - nil result", i)
			continue
		}

		if res.Result.Location == nil {
			log.Printf("%d - nil location", i)
			continue
		}

		p1 := NewPoint(t.ExpectedResult.Location.Latitude, t.ExpectedResult.Location.Longitude, 0)
		p2 := NewPoint(res.Result.Location.Latitude, res.Result.Location.Longitude, 0)

		if err := w.Write([]string{
			strconv.FormatInt(int64(i), 10),
			strconv.FormatFloat(t.ExpectedResult.Location.Latitude, 'f', 6, 64),
			strconv.FormatFloat(t.ExpectedResult.Location.Longitude, 'f', 6, 64),
			strconv.FormatFloat(t.ExpectedResult.Location.Altitude, 'f', 6, 64),
			strconv.FormatFloat(res.Result.Location.Latitude, 'f', 6, 64),
			strconv.FormatFloat(res.Result.Location.Longitude, 'f', 6, 64),
			strconv.FormatFloat(res.Result.Location.Altitude, 'f', 6, 64),
			strconv.FormatFloat(p1.Distance(p2), 'f', 6, 64),
		}); err != nil {
			log.Fatal(err)
		}

	}

	w.Flush()

	return nil
}
