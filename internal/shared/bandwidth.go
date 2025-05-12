package shared

import "errors"

type BandWidth struct {
	downloadMbps int
	uploadMbps   int
}

func NewBandWidth(downloadMbps, uploadMbps int) (BandWidth, error) {
	if downloadMbps < 0 || uploadMbps < 0 {
		return BandWidth{}, errors.New("bandwidth values cannot be negative")
	}
	return BandWidth{
		downloadMbps: downloadMbps,
		uploadMbps:   uploadMbps,
	}, nil
}
