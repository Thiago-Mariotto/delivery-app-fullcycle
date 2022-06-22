package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type Route struct {
	Id       string     `json:"routeId"`
	ClientId string     `json:"clientId"`
	Position []Position `json:"position"`
}

type PartialRoutePostion struct {
	Id       string    `json:"routeId"`
	ClientId string    `json:"ClientId"`
	Position []float64 `json:"position"`
	Finished bool      `json:"finished"`
}

func (r *Route) LoadPositions() error {
	if r.Id == "" {
		return errors.New("route id not informed")
	}

	file, err := os.Open("destination/" + r.Id + ".txt")
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

		r.Position = append(r.Position, Position{
			Lat: lat, Long: long,
		})
	}

	return nil
}

func (r *Route) ExportJsonPosition() ([]string, error) {
	var route PartialRoutePostion
	var result []string
	total := len(r.Position)

	for k, v := range r.Position {
		route.Id = r.Id
		route.ClientId = r.ClientId
		route.Position = []float64{v.Lat, v.Long}
		route.Finished = false

		if total-1 == k {
			route.Finished = true
		}

		jsonRoute, err := json.Marshal(route)
		if err != nil {
			return nil, err
		}

		result = append(result, string(jsonRoute)) // parse slice de byte para string
	}

	return result, nil
}
