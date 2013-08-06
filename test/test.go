package main
import (
    "fmt"
    "github.com/apanda/goigraph/goigraph"
    "os"
    "flag"
)

func main() {
    flag.Parse()
    if flag.NArg() == 0 {
        fmt.Printf("Error: No file\n")
        os.Exit(0)
    }
    file, err := os.Open(flag.Arg(0))
    if err != nil {
        fmt.Printf("Error: Could not open file %v\n", err)
        os.Exit(0)
    }

    graph := goigraph.ReadEdgeList (file, false)
    if graph == nil {
        fmt.Printf("ReadEdgeList returned nil, error parsing\n")
    }
    fmt.Printf("Mincut = %v\n", graph.MinCutValue())
}
