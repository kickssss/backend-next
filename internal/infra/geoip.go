package infra

import (
	_ "embed"

	"github.com/oschwald/geoip2-golang"

	"exusiai.dev/backend-next/internal/app/appconfig"
	"exusiai.dev/backend-next/internal/pkg/geoip"
)

func GeoIPDatabase(conf *appconfig.Config) (*geoip2.Reader, error) {
	db, err := geoip2.FromBytes(geoip.Database)
	if err != nil {
		return nil, err
	}

	return db, nil
}
