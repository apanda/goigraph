// A Go library wrapping igraph
package goigraph
/*
#cgo LDFLAGS: -L/usr/local/lib -ligraph
#cgo CFLAGS: -I/usr/local/include/igraph
#include <igraph.h>
#include <stdio.h>
*/
import "C"
import (
    "os"
)

// An opaque type representing an igraph graph
type GoGraph struct {
    graph C.igraph_t
}

func booltoint (in bool) C.int {
    if in {
        return C.int(1)
    }
    return C.int(0)
}


// Read an edge list file and produce a graph
func ReadEdgeList (file *os.File, directed bool) *GoGraph {
    graph := &GoGraph {}
    fstruct := C.fdopen(C.int(file.Fd()), C.CString("r"))
    err := C.igraph_read_graph_edgelist (&graph.graph, fstruct, C.igraph_integer_t(C.int(0)), C.igraph_bool_t(booltoint(directed)))
    if err == C.IGRAPH_PARSEERROR {
        return nil
    }
    return graph
}

// Read a GraphML file and produce a graph
func ReadGraphML (file *os.File, index int) *GoGraph {
    graph := &GoGraph {}
    fstruct := C.fdopen(C.int(file.Fd()), C.CString("r"))
    err := C.igraph_read_graph_graphml (&graph.graph, fstruct, C.int(index))
    if err == C.IGRAPH_PARSEERROR {
        return nil
    }
    return graph
}

// Read a GML file and produce a graph
func ReadGML (file *os.File) *GoGraph {
    graph := &GoGraph {}
    fstruct := C.fdopen(C.int(file.Fd()), C.CString("r"))
    err := C.igraph_read_graph_gml (&graph.graph, fstruct)
    if err == C.IGRAPH_PARSEERROR {
        return nil
    }
    return graph
}

// Read a Pajek file and produce a graph
func ReadPajek (file *os.File) *GoGraph {
    graph := &GoGraph {}
    fstruct := C.fdopen(C.int(file.Fd()), C.CString("r"))
    err := C.igraph_read_graph_pajek (&graph.graph, fstruct)
    if err == C.IGRAPH_PARSEERROR {
        return nil
    }
    return graph
}

// Compute the mincut for a graph, assuming all edges have unit capacity
func (graph *GoGraph) MinCutValue () float64 {
    var result C.igraph_real_t
    err := C.igraph_mincut_value(&graph.graph, &result, nil)
    _ = err
    res := float64(result)
    return res
}
