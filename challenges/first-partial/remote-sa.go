package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
)

type Point struct {
	X, Y float64
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//generatePoints array
func generatePoints(s string) ([]Point, error) {

	points := []Point{}

	s = strings.Replace(s, "(", "", -1)
	s = strings.Replace(s, ")", "", -1)
	vals := strings.Split(s, ",")
	if len(vals) < 2 {
		return []Point{}, fmt.Errorf("Point [%v] was not well defined", s)
	}

	var x, y float64

	for idx, val := range vals {

		if idx%2 == 0 {
			x, _ = strconv.ParseFloat(val, 64)
		} else {
			y, _ = strconv.ParseFloat(val, 64)
			points = append(points, Point{x, y})
		}
	}
	return points, nil
}

// getArea gets the area inside from a given shape
func getArea(points []Point) float64 {
	// Your code goes here
	var area float64
	var perimetro float64
	area = 0
	perimetro = 0

	var lados []float64
	//area = riz(sp(p-a)(p-b)(p-c))
	log.Printf("Numero de Vectores: %v\n", len(points))

	for i := 0; i < len(points)-1; i++ {
		perimetro += math.Hypot(points[i+1].X-points[i].X, points[i+1].Y-points[i].Y)
		lados = append(lados, math.Hypot(points[i+1].X-points[i].X, points[i+1].Y-points[i].Y))
	}
	perimetro += math.Hypot(points[len(points)-1].X-points[0].X, points[len(points)-1].Y-points[0].Y)
	lados = append(lados, math.Hypot(points[len(points)-1].X-points[0].X, points[len(points)-1].Y-points[0].Y))

	var semiperimetro float64
	semiperimetro = perimetro / 2

	if (len(points)) > 3 {
		if len(points) == 4 {
			var lado1 float64
			lado1 = math.Hypot(lados[0], lados[1])

			area += math.Sqrt(semiperimetro * (semiperimetro - lados[0]) * (semiperimetro - lados[1]) * (semiperimetro - lado1))
			area += math.Sqrt(semiperimetro * (semiperimetro - lados[2]) * (semiperimetro - lados[3]) * (semiperimetro - lado1))

			log.Printf("Lado faltante %v", lado1)

			return area
		}
		if len(points) == 5 {
			var lado1, lado2 float64
			lado1 = math.Hypot(lados[0], lados[1])
			lado2 = math.Hypot(lados[2], lados[3])

			area += math.Sqrt(semiperimetro * (semiperimetro - lados[0]) * (semiperimetro - lados[1]) * (semiperimetro - lado1))
			area += math.Sqrt(semiperimetro * (semiperimetro - lados[2]) * (semiperimetro - lados[3]) * (semiperimetro - lado2))
			area += math.Sqrt(semiperimetro * (semiperimetro - lados[4]) * (semiperimetro - lado1) * (semiperimetro - lado2))

			log.Printf("Lado faltante %v", lado1)

			return area
		}
	}

	area = math.Sqrt(semiperimetro * (semiperimetro - lados[0]) * (semiperimetro - lados[1]) * (semiperimetro - lados[2]))

	return area
}

// getPerimeter gets the perimeter from a given array of connected points
func getPerimeter(points []Point) float64 {
	// Your code goes here
	var perimetro float64
	perimetro = 0

	for i := 0; i < len(points)-1; i++ {
		perimetro += math.Hypot(points[i+1].X-points[i].X, points[i+1].Y-points[i].Y)
	}
	perimetro += math.Hypot(points[len(points)-1].X-points[0].X, points[len(points)-1].Y-points[0].Y)

	return perimetro
}

// handler handles the web request and reponds it
func handler(w http.ResponseWriter, r *http.Request) {

	var vertices []Point
	for k, v := range r.URL.Query() {
		if k == "vertices" {
			points, err := generatePoints(v[0])
			if err != nil {
				fmt.Fprintf(w, fmt.Sprintf("error: %v", err))
				return
			}
			vertices = points
			break
		}
	}

	// Results gathering
	area := getArea(vertices)
	perimeter := getPerimeter(vertices)

	// Logging in the server side
	log.Printf("Received vertices array: %v", vertices)

	// Response construction
	response := fmt.Sprintf("Welcome to the Remote Shapes Analyzer\n")
	if len(vertices) > 2 {
		response += fmt.Sprintf(" - Your figure has : [%v] vertices\n", len(vertices))
		response += fmt.Sprintf(" - Vertices        : %v\n", vertices)
		response += fmt.Sprintf(" - Perimeter       : %v\n", perimeter)
		response += fmt.Sprintf(" - Area            : %v\n", area)
	} else {
		response += fmt.Sprintf("Error: [%v] vertices\n", len(vertices))
		response += fmt.Sprintf("You need more to make a triangle\n")

	}

	// Send response to client
	fmt.Fprintf(w, response)
}
