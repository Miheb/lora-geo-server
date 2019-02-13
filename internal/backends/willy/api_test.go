package willy

import (
	"context"
	//"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/brocaar/loraserver/api/common"
	"github.com/brocaar/loraserver/api/geo"
	"github.com/brocaar/loraserver/api/gw"
)


type WillyTestSuite struct {
	suite.Suite

	apiResponse string
	apiRequest  string
	apiServer   *httptest.Server

	client geo.GeolocationServerServiceServer
}

func (ts *WillyTestSuite) SetupSuite() {
	log.SetLevel(log.ErrorLevel)

	ts.apiServer = httptest.NewServer(http.HandlerFunc(ts.apiHandler))

	ts.client = &APIWilly{
		config: Config{
			RequestTimeout: time.Second,
		},
	}

	//fmt.Printf("ts.apiServer.URL %+v\n", ts.apiServer.URL)
	//tdoaEndpoint = ts.apiServer.URL
}

func (ts *WillyTestSuite) TearDownSuite() {
	ts.apiServer.Close()
}

func (ts *WillyTestSuite) TestResolveTDOA() {
	ts.apiResponse = `
		{
			"result": {
			"latitude": 1.12345,
			"longitude": 1.22345,
			"altitude": 1.32345,
			"accuracy": 4.5,
			"algorithmType": "a-algorithm",
			"numberOfGatewaysReceived": 4,
			"numberOfGatewaysUsed": 3
			},
			"warnings": [
			],
			"errors": [
			],
			"correlationId": "abcde"
		}
	`

	now := time.Now()
	nowPB, _ := ptypes.TimestampProto(now)

	testTable := []struct {
		Name    string
		Request geo.ResolveTDOARequest

		ExpectedError    error
		ExpectedRequest  *tdoaRequest
		ExpectedResponse *geo.ResolveTDOAResponse
	}{
		{
			Name: "not enough meta-data",
			Request: geo.ResolveTDOARequest{
				DevEui: []byte{1, 2, 3, 4, 5, 6, 7, 8},
				FrameRxInfo: &geo.FrameRXInfo{
					RxInfo: []*gw.UplinkRXInfo{
						{
							GatewayId:         []byte{1, 1, 1, 1, 1, 1, 1, 1},
							FineTimestampType: gw.FineTimestampType_NONE,
						},
						{
							GatewayId:         []byte{2, 1, 1, 1, 1, 1, 1, 1},
							FineTimestampType: gw.FineTimestampType_PLAIN,
							Location:          &common.Location{},
						},
						{
							GatewayId:         []byte{3, 1, 1, 1, 1, 1, 1, 1},
							FineTimestampType: gw.FineTimestampType_ENCRYPTED,
							Location:          &common.Location{},
						},
						{
							GatewayId:         []byte{4, 1, 1, 1, 1, 1, 1, 1},
							FineTimestampType: gw.FineTimestampType_ENCRYPTED,
							Location:          &common.Location{},
							FineTimestamp: &gw.UplinkRXInfo_EncryptedFineTimestamp{
								EncryptedFineTimestamp: &gw.EncryptedFineTimestamp{
									EncryptedNs: []byte{1, 2, 3, 4, 5},
									FpgaId:      []byte{},
								},
							},
						},
					},
				},
			},
			ExpectedError: grpc.Errorf(codes.InvalidArgument, "not enough meta-data for geolocation"),
		},
		{
			Name: "valid decrypted timestamp request",
			Request: geo.ResolveTDOARequest{
				DevEui: []byte{1, 2, 3, 4, 5, 6, 7, 8},
				FrameRxInfo: &geo.FrameRXInfo{
					RxInfo: []*gw.UplinkRXInfo{
						{
							GatewayId: []byte{1, 1, 1, 1, 1, 1, 1, 1},
							Location: &common.Location{
								Latitude:  1.1,
								Longitude: 1.2,
								Altitude:  1.3,
							},
							FineTimestampType: gw.FineTimestampType_PLAIN,
							FineTimestamp: &gw.UplinkRXInfo_PlainFineTimestamp{
								PlainFineTimestamp: &gw.PlainFineTimestamp{
									Time: nowPB,
								},
							},
						},
						{
							GatewayId: []byte{2, 1, 1, 1, 1, 1, 1, 1},
							Location: &common.Location{
								Latitude:  2.1,
								Longitude: 2.2,
								Altitude:  2.3,
							},
							FineTimestampType: gw.FineTimestampType_PLAIN,
							FineTimestamp: &gw.UplinkRXInfo_PlainFineTimestamp{
								PlainFineTimestamp: &gw.PlainFineTimestamp{
									Time: nowPB,
								},
							},
						},
						{
							GatewayId: []byte{3, 1, 1, 1, 1, 1, 1, 1},
							Location: &common.Location{
								Latitude:  3.1,
								Longitude: 3.2,
								Altitude:  3.3,
							},
							FineTimestampType: gw.FineTimestampType_PLAIN,
							FineTimestamp: &gw.UplinkRXInfo_PlainFineTimestamp{
								PlainFineTimestamp: &gw.PlainFineTimestamp{
									Time: nowPB,
								},
							},
						},
					},
				},
			},
			ExpectedRequest: &tdoaRequest{
				LoRaWAN: []loRaWANRX{
					{
						GatewayID: "0101010101010101",
						TOA:       now.Nanosecond(),
						AntennaLocation: antennaLocation{
							Latitude:  1.1,
							Longitude: 1.2,
							Altitude:  1.3,
						},
					},
					{
						GatewayID: "0201010101010101",
						TOA:       now.Nanosecond(),
						AntennaLocation: antennaLocation{
							Latitude:  2.1,
							Longitude: 2.2,
							Altitude:  2.3,
						},
					},
					{
						GatewayID: "0301010101010101",
						TOA:       now.Nanosecond(),
						AntennaLocation: antennaLocation{
							Latitude:  3.1,
							Longitude: 3.2,
							Altitude:  3.3,
						},
					},
				},
			},
			ExpectedResponse: &geo.ResolveTDOAResponse{
				Result: &geo.ResolveResult{
					Location: &common.Location{
						Latitude:  1.12345,
						Longitude: 1.22345,
						Altitude:  1.32345,
						Source:    common.LocationSource_GEO_RESOLVER,
						Accuracy:  4,
					},
				},
			},
		},
		{
			Name: "valid encrypted timestamp request",
			Request: geo.ResolveTDOARequest{
				DevEui: []byte{1, 2, 3, 4, 5, 6, 7, 8},
				FrameRxInfo: &geo.FrameRXInfo{
					RxInfo: []*gw.UplinkRXInfo{
						{
							GatewayId: []byte{1, 1, 1, 1, 1, 1, 1, 1},
							Location: &common.Location{
								Latitude:  1.1,
								Longitude: 1.2,
								Altitude:  1.3,
							},
							FineTimestampType: gw.FineTimestampType_ENCRYPTED,
							FineTimestamp: &gw.UplinkRXInfo_EncryptedFineTimestamp{
								EncryptedFineTimestamp: &gw.EncryptedFineTimestamp{
									FpgaId:      []byte{1},
									EncryptedNs: []byte{1, 1, 1, 1},
								},
							},
						},
						{
							GatewayId: []byte{2, 1, 1, 1, 1, 1, 1, 1},
							Location: &common.Location{
								Latitude:  2.1,
								Longitude: 2.2,
								Altitude:  2.3,
							},
							FineTimestampType: gw.FineTimestampType_ENCRYPTED,
							FineTimestamp: &gw.UplinkRXInfo_EncryptedFineTimestamp{
								EncryptedFineTimestamp: &gw.EncryptedFineTimestamp{
									FpgaId:      []byte{2},
									EncryptedNs: []byte{2, 1, 1, 1},
								},
							},
						},
						{
							GatewayId: []byte{3, 1, 1, 1, 1, 1, 1, 1},
							Location: &common.Location{
								Latitude:  3.1,
								Longitude: 3.2,
								Altitude:  3.3,
							},
							FineTimestampType: gw.FineTimestampType_ENCRYPTED,
							FineTimestamp: &gw.UplinkRXInfo_EncryptedFineTimestamp{
								EncryptedFineTimestamp: &gw.EncryptedFineTimestamp{
									FpgaId:      []byte{3},
									EncryptedNs: []byte{3, 1, 1, 1},
								},
							},
						},
					},
				},
			},
			ExpectedRequest: &tdoaRequest{
				LoRaWAN: []loRaWANRX{
					{
						GatewayID:    "0x01",
						EncryptedTOA: "AQEBAQ==",
						AntennaLocation: antennaLocation{
							Latitude:  1.1,
							Longitude: 1.2,
							Altitude:  1.3,
						},
					},
					{
						GatewayID:    "0x02",
						EncryptedTOA: "AgEBAQ==",
						AntennaLocation: antennaLocation{
							Latitude:  2.1,
							Longitude: 2.2,
							Altitude:  2.3,
						},
					},
					{
						GatewayID:    "0x03",
						EncryptedTOA: "AwEBAQ==",
						AntennaLocation: antennaLocation{
							Latitude:  3.1,
							Longitude: 3.2,
							Altitude:  3.3,
						},
					},
				},
			},
			ExpectedResponse: &geo.ResolveTDOAResponse{
				Result: &geo.ResolveResult{
					Location: &common.Location{
						Latitude:  1.12345,
						Longitude: 1.22345,
						Altitude:  1.32345,
						Source:    common.LocationSource_GEO_RESOLVER,
						Accuracy:  4,
					},
				},
			},
		},
	}

	for _, test := range testTable {
		ts.T().Run(test.Name, func(t *testing.T) {
			assert := require.New(t)

			resp, err := ts.client.ResolveTDOA(context.Background(), &test.Request)
			assert.Equal(test.ExpectedError, err)

			if test.ExpectedResponse != nil {
				assert.Equal(test.ExpectedResponse, resp)
			}

			// TODO : would be necessary once the request is made from an actual
			// server
			/*if test.ExpectedRequest != nil {
				var req tdoaRequest
				assert.NoError(json.Unmarshal([]byte(ts.apiRequest), &req))
				assert.Equal(test.ExpectedRequest, &req)
			}*/
		})
	}
}

func (ts *WillyTestSuite) apiHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	ts.apiRequest = string(b)
	w.Write([]byte(ts.apiResponse))
}

func TestWilly(t *testing.T) {
	suite.Run(t, new(WillyTestSuite))
}
