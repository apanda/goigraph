package main
import (
    "fmt"
    "github.com/apanda/goigraph/goigraph"
    "os"
    "flag"
)

func main() {
    k4 := flag.Bool("K4", false, "Make K4")
    outfile := flag.String("out", "", "Filename to output graph")
    var graph *goigraph.GoGraph
    flag.Parse()
    if *k4 {
        graph = goigraph.CreateKn(4)
    } else {
        if flag.NArg() == 0 {
            fmt.Printf("Error: No file\n")
            os.Exit(0)
        }
        file, err := os.Open(flag.Arg(0))
        if err != nil {
            fmt.Printf("Error: Could not open file %v\n", err)
            os.Exit(0)
        }

        graph = goigraph.ReadEdgeList (file, false)
    }
    if graph == nil {
        fmt.Printf("Could not find a graph\n")
        return
    }
    if *outfile != "" {
        fmt.Printf("Writing file\n")
        ofile, err := os.Create (*outfile)
        if err != nil {
            fmt.Printf("Error: Could not open file for output %v\n", err)
            os.Exit(0)
        }
        graph.WriteGML(ofile)
        ofile.Close()
    }

    fmt.Printf("Mincut = %v\n", graph.MinCutValue())
}
