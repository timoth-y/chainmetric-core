package models

import (
	"github.com/kellydunn/golang-geo"
)

// Location defines the geospatial location data model.
type Location struct {
	Name      string `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// IsNearBy determines whether two geo points is located within `d` distance radius (m).
func (l Location) IsNearBy(l2 Location, d float64) bool {
	return l.Distance(l2) <= d
}

// Distance calculates distance between two geo points (in m).
func (l Location) Distance(l2 Location) float64 {
	return l.toGeoPoint().GreatCircleDistance(l2.toGeoPoint()) * 1000
}

func (l Location) toGeoPoint() *geo.Point {
	return geo.NewPoint(l.Latitude, l.Longitude)
}
