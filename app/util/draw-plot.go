package util

import (
    "gonum.org/v1/plot"
    "gonum.org/v1/plot/plotter"
    "gonum.org/v1/plot/plotutil"
    "gonum.org/v1/plot/vg"
    "fmt"
	"strconv"
    "log"
)


var data [][]string
var minX = 25
var maxX = 50
var minY = 75
var maxY = 130

func DrawPlot() {
	data = LoadDatabase()
	fmt.Println("DrawPlot entered.")
    p, err := plot.New()
	fmt.Println("Plot Object created")
	if err != nil {
        log.Fatal(err)
    }

    p.Title.Text = "Plotutil example"
    p.X.Label.Text = "Longitude"
    p.Y.Label.Text = "Latitude"
	fmt.Println("Plot Axis Labels initialized")

	err = plotutil.AddScatters(p,
        "Earthquake Report", GraphPoints(data))
    if err != nil {
        log.Fatal(err)
    }
	fmt.Println("Earthquake Report init")
    // Save the plot to a PNG file.
    if err := p.Save(4*vg.Inch, 4*vg.Inch, "earthquake.png"); err != nil {
        log.Fatal(err)
    }
}

// Graph Points returns a scatter graph of data specified in file-reader.
func GraphPoints(data [][]string) plotter.XYs {
	fmt.Println("GraphPoints called")
	pts := make(plotter.XYs, len(data))
	fmt.Println("pts initialized")
	for i := 1; i < len(pts); i++ {
		//fmt.Println("data[i][2]")
		//fmt.Println(data[i][2])
		//fmt.Println("data[i][3]")
		//fmt.Println(data[i][3])
		X, err := strconv.ParseFloat(data[i][2], 64)
		if err != nil{
			fmt.Println("float X parsing error")
			log.Fatal(err)
		}
		Y, err := strconv.ParseFloat(data[i][3], 64)
		if err != nil{
			fmt.Println("float Y parsing error")
			log.Fatal(err)
		}
		onMainland := parseData(X, Y)
		if bool(onMainland) {
			if i == 0 {
				pts[i].X = X
			} else {
				pts[i].X = pts[i-1].X + X
			}
			pts[i].Y = pts[i].Y + Y
		}
	}

    return pts
}
func parseData(X float64, Y float64) bool{
	//fmt.Println(int(X))
	if int(X) > minX && int(X) < maxX{
		if int(Y) > minY && int(Y) < maxY{
			//fmt.Println(X, Y)
			return true
		}
	}
	return false
}
