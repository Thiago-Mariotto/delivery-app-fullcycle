package route

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	Lat  float64
	Long float64
}

type Route struct {
	Id       string
	ClientId string
	Position []Position
}

type PartialRoutePostion struct {
	Id       string
	ClientId string
	Position []float64
	Finished bool
}

func (route *Route) LoadPositions() error {
	if route.Id == "" {
		return errors.New("route id not informed")
	}

	file, err := os.Open("destination/" + route.Id + ".txt")
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		lat, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return nil
		}

		long, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return nil
		}

		route.Position = append(route.Position, Position{
			Lat: lat, Long: long,
		})
	}

	return nil
}
