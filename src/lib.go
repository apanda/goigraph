// A Go library wrapping igraph
package goigraph
// #cgo pkg-config: igraph gsl
// #include <igraph.h>
// #include <stdio.h>
// #include "primitives.h"
import "C"
import (
    "os"
    "errors"
)

// An opaque type representing an igraph graph
type GoGraph struct {
    graph C.igraph_t
}

type GoVector struct {
    vector C.igraph_vector_t
}

func booltoint (in bool) C.int {
    if in {
        return C.int(1)
    }
    return C.int(0)
}

// Create a random 3 connected graph using BG-operations
func CreateRandom3Connected (vertices uint32) *GoGraph {
    graph := &GoGraph {}
    C.construct3ConnectedGraph (&graph.graph, C.uint32_t(vertices))
    return graph
}

// Evolve a 3 connected graph; i.e. starting with a graph produce another
func (graph *GoGraph) Evolve3Connected (verticesToAdd uint32) {
    C.evolve3ConnectedGraph (&graph.graph, C.uint32_t(verticesToAdd));
}

// Create and return a K4
func CreateKn (n int) *GoGraph {
    graph := &GoGraph {}
    err := C.igraph_full(&graph.graph, C.igraph_integer_t(C.int(n)), C.igraph_bool_t(booltoint(false)), C.igraph_bool_t(booltoint(false)))
    if err != 0 {
        return nil
    }
    return graph
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

// Write GraphML file
func (graph *GoGraph) WriteGraphML (file *os.File) error {
    //C.testWriteGraphMl(&graph.graph) 
    //return nil
    fstruct := C.fdopen(C.int(file.Fd()), C.CString("w"))
    err := C.igraph_write_graph_graphml(&graph.graph, fstruct, 0)
    if err != 0 {
        return errors.New("Write failed")
    }
    C.fflush(fstruct)
    return nil
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

// Write GML file
func (graph *GoGraph) WriteGML (file *os.File) error {
    //C.testWriteGraphMl(&graph.graph) 
    //return nil
    fstruct := C.fdopen(C.int(file.Fd()), C.CString("w"))
    err := C.igraph_write_graph_gml(&graph.graph, fstruct, nil, nil)
    if err != 0 {
        return errors.New("Write failed")
    }
    C.fflush(fstruct)
    return nil
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

