/* Some primitive operations for creating igraph structures of interest
 *
 */
#include <igraph.h>
#include <stdbool.h>
#include <stdio.h>

/// Create a 4-clique (the base 3-connected graph).
int createK4 (igraph_t* graph) {
    const int VERTICES = 4;
    const int EDGES = 4 * 3; // C(4, 2)
    int idx = 0;
    int err = 0;
    int v0 = 0;
    int v1 = 0;
    igraph_vector_t edges;
    igraph_vector_init (&edges, EDGES);
    for (v0 = 0; v0 < VERTICES; v0++) {
        for (v1 = v0 + 1; v1 < VERTICES; v1++) {
            VECTOR(edges)[idx] = v0;
            idx++;
            VECTOR(edges)[idx] = v1;
            idx++;
        }
    }
    err = igraph_create (graph, &edges, VERTICES, false);
    igraph_vector_destroy (&edges);
    if (err != 0) {
        return err; // Could not create a graph for some reason
    }
    return 0;
}

void testWriteGraphMl (igraph_t* graph) {
    const char* fname = "temp.graphml";
    int result;
    FILE* ofile;
    printf("Graph has %d vertices, %d edges\n", igraph_vcount(graph), igraph_ecount(graph));
    ofile = fopen(fname, "w+");
    if (ofile) {
        if (!(result = igraph_write_graph_graphml(graph, ofile))) {
            printf("Written graph\n");
            fclose(ofile);
        } else {
            printf("Write failed %d\n", result);
        }
    } else {
        printf("fopen failed\n");
    }
    ofile = fopen("temp.gml", "w+");
    if (ofile) {
        if (!(result = igraph_write_graph_gml(graph, ofile, NULL, NULL))) {
            printf("Written graph\n");
            fclose(ofile);
        } else {
            printf("Write failed %d\n", result);
        }

    } else {
        printf("fopen failed\n");
    }
}
