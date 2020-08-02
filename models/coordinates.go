package models

import (
	"fmt"
	"strconv"
	"strings"
)

type Coordinates struct {
	Long string
	Lat  string
}

func (c Coordinates) ToString() string {
	return fmt.Sprintf("%s,%s", c.Long, c.Lat)
}

func NewCoordinateFromString(coordinateStr string) (Coordinates, error) {
	splitted := strings.Split(coordinateStr, ",")
	if len(splitted) != 2 {
		return Coordinates{}, fmt.Errorf("cannot read coordinates from string")
	}
	if _, err := strconv.ParseFloat(splitted[0], 32); err != nil {
		return Coordinates{}, fmt.Errorf("cannot parse coordinates: %s", err.Error())
	}
	if _, err := strconv.ParseFloat(splitted[1], 32); err != nil {
		return Coordinates{}, fmt.Errorf("cannot parse coordinates: %s", err.Error())
	}

	return Coordinates{
		Long: splitted[0],
		Lat:  splitted[1],
	}, nil
}

func NewCoordinatesCollectionFromString(coordinateStrings []string) (result []Coordinates, err error) {
	for _, v := range coordinateStrings {
		coordinate, err := NewCoordinateFromString(v)
		if err != nil {
			return nil, err
		}
		result = append(result, coordinate)
	}
	return result, nil
}
