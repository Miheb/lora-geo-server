package cmd

import (
	"os"
	"text/template"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/brocaar/lora-geo-server/internal/config"
)

const configTemplate = `[general]
# Log level
#
# debug=5, info=4, warning=3, error=2, fatal=1, panic=0
log_level={{ .General.LogLevel }}

# Geolocation-server configuration.
[geo_server]
  # Geolocation API.
  #
  # This is the geolocation API that can be used by LoRa Server.
  [geo_server.api]
  # ip:port to bind the api server
  bind="{{ .GeoServer.API.Bind }}"

  # CA certificate used by the api server (optional)
  ca_cert="{{ .GeoServer.API.CACert }}"

  # TLS certificate used by the api server (optional)
  tls_cert="{{ .GeoServer.API.TLSCert }}"

  # TLS key used by the api server (optional)
  tls_key="{{ .GeoServer.API.TLSKey }}"


  # Geolocation backend configuration.
  [geo_server.backend]
  # Name.
  #
  # The name of the geolocation backend to use.
  name="{{ .GeoServer.Backend.Name }}"

  [geo_server.backend.collos]
  # Collos subscription key.
  #
  # This key can be retrieved after creating a Collos account at:
  # http://preview.collos.org/
  subscription_key="{{ .GeoServer.Backend.Collos.SubscriptionKey }}"

  # Request timeout.
  #
  # This defines the request timeout when making calls to the Collos API.
  request_timeout="{{ .GeoServer.Backend.Collos.RequestTimeout }}"
`

var configfileCmd = &cobra.Command{
	Use:   "configfile",
	Short: "Print the LoRa Geolocation Server configuration file",
	RunE: func(cmd *cobra.Command, args []string) error {
		t := template.Must(template.New("config").Parse(configTemplate))
		err := t.Execute(os.Stdout, &config.C)
		if err != nil {
			return errors.Wrap(err, "execute config template error")
		}
		return nil
	},
}
