package httpRest

import "notary-public-online/internal/configs/yaml"

type Rest interface {
	Start(cfg *yaml.Config) error
	// Shutdown() error
}
