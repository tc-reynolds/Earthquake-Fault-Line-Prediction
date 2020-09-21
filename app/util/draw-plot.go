package util

import (
    "gonum.org/v1/plot"
    "gonum.org/v1/plot/plotter"
    "gonum.org/v1/plot/plotutil"
    "gonum.org/v1/plot/vg"
    "fmt"
)

func DrawPlot() {

    p, err := plot.New()
    if err != nil {
        panic(err)
    }

    p.Title.Text = "Plotutil example"
    p.X.Label.Text = "X"
    p.Y.Label.Text = "Y"

    err = plotutil.AddScatters(p,
        "Earthquake Report", GraphPoints(LoadDatabase()))
    if err != nil {
        panic(err)
    }

    // Save the plot to a PNG file.
    if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
        panic(err)
    }
}

// Graph Points returns a scatter graph of data specified in file-reader.
func GraphPoints(data [][]string) plotter.XYs {
    pts := make(plotter.XYs, 15)
    for i := range pts {
        if i == 0 {
            pts[i].X = rand.Float64()
        } else {
            pts[i].X = pts[i-1].X + rand.Float64()
        }
        pts[i].Y = pts[i].X + 10*rand.Float64()
    }
    return pts
}
