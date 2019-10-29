// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

/*
 * Following the approach fromm the Lissajous example in Section 1.7, construct
 * a webserver that computes surfaces and writes SVG data to the client. The server
 * must set the Content-Type header like this:
 *
 * w.Header().Set("Content-Type", "image/svg+xml")
 *
 * (This step was not required in the Lissajous example because the server uses
 * standard heuristics to recognise common formats like PNG from the first 512
 * bytes of the response, and generates the proper header.) Allow the client to
 * specify values like the height, width, and color as HTTP request parameters.
 */

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
	color         = "ffffff"            // colour of polygon fill
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)
var svgParams = &svgParameters{}

type svgParameters struct {
	Height int
	Width  int
	Color  string
}

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		query := r.URL.Query()
		if newHeight, ok := fetchIntFromQuery(query, "height"); ok {
			svgParams.Height = newHeight
		} else {
			svgParams.Height = height
		}

		if newWidth, ok := fetchIntFromQuery(query, "width"); ok {
			svgParams.Width = newWidth
		} else {
			svgParams.Width = width
		}

		if newColor, ok := fetchStringFromQuery(query, "color"); ok {
			svgParams.Color = newColor
		} else {
			svgParams.Color = color
		}

		fmt.Fprintf(w, generateSVG(svgParams))
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func fetchIntFromQuery(q url.Values, param string) (int, bool) {
	values, ok := q[param]
	if !ok {
		return 0, false
	}

	value, err := strconv.Atoi(values[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing %s: %v", param, err)
	}
	return value, true
}

func fetchStringFromQuery(q url.Values, param string) (string, bool) {
	values, ok := q[param]
	if !ok {
		return "", false
	}

	return values[0], true
}
func generateSVG(p *svgParameters) string {
	var result string
	result += fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: #%s; stroke-width: 0.7' "+
		"width='%d' height='%d'>", p.Color, p.Width, p.Height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			var ax, ay, bx, by, cx, cy, dx, dy float64
			var ok bool
			if ax, ay, ok = corner(i+1, j); !ok {
				continue
			}
			if bx, by, ok = corner(i, j); !ok {
				continue
			}
			if cx, cy, ok = corner(i, j+1); !ok {
				continue
			}
			if dx, dy, ok = corner(i+1, j+1); !ok {
				continue
			}
			result += fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	result += fmt.Sprintf("</svg>")
	return result
}

func corner(i, j int) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, false
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
