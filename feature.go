package geo

import (
	"strings"
)

type GeometryType string

const (
	PolygonFeature GeometryType = "polygon"
	LineFeature    GeometryType = "line"
	PointFeature   GeometryType = "point"
)

type Feature struct {
	ID         interface{}
	Geometry   []*Shape
	Type       string
	Properties map[string]interface{}
}

func NewFeature(geometryType string, geometry ...*Shape) *Feature {
	geometryType = strings.ToLower(geometryType)
	return &Feature{Geometry: geometry, Type: geometryType}
}

func NewPolygonFeature(geometry ...*Shape) *Feature {
	return NewFeature(string(PolygonFeature), geometry...)
}

func NewLineFeature(geometry ...*Shape) *Feature {
	return NewFeature(string(LineFeature), geometry...)
}

func NewPointFeature(geometry ...*Shape) *Feature {
	return NewFeature(string(PointFeature), geometry...)
}

func MakeFeature(length int) *Feature {
	return &Feature{Geometry: make([]*Shape, length)}
}

func (f *Feature) AddShape(s *Shape) {
	f.Geometry = append(f.Geometry, s)
}

func (f *Feature) Tags(key string) string {
	return f.Properties[key].(string)
}

func (f *Feature) Center() (avg Coordinate) {
	div := 0.0
	avg = Coordinate{Lat: 0, Lon: 0}
	for _, shape := range f.Geometry {
		for _, c := range shape.Coordinates {
			avg.Lat += c.Lat
			avg.Lon += c.Lon
			div += 1
		}
	}
	avg.Lat /= div
	avg.Lon /= div
	return
}

//Only checks as exterior ring
func (f *Feature) Contains(c Coordinate) bool {
	for _, shp := range f.Geometry {
		if shp.Contains(c) {
			return true
		}
	}
	return false
}
