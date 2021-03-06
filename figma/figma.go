package figma

import (
	"net/http"

	"github.com/spf13/viper"
)

// Figma wraps the incomplete figma client
type Figma struct {
	config     *viper.Viper
	httpClient *http.Client
}

// New returns a new instace of Figma
func New(config *viper.Viper, httpClient *http.Client) Figma {
	return Figma{
		config:     config,
		httpClient: httpClient,
	}
}
