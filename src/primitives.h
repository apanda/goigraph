/* Some primitive operations for creating igraph structures of interest
 *
 */
#include <igraph.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdint.h>
#include <gsl/gsl_rng.h>
gsl_rng *r = NULL;
//Init the rng
inline void InitRng () {
    const gsl_rng_type *T;
    unsigned long int seed;
    FILE* random = fopen("/dev/random", "r");
    fread(&seed, 1, sizeof(seed), random);
    fclose(random);
    if (r == NULL) {
        gsl_rng_env_setup();
        T = gsl_rng_default;
        r = gsl_rng_alloc (T);
        gsl_rng_set (r, seed);
    }
}

// Split one edge and add a new edge connecting the new vertex to a third vertex
void splitEdge (igraph_t* graph) {
    // # of edges
    int32_t edges = igraph_ecount(graph);
    // # of vertices
    int32_t vertices = igraph_vcount(graph);
    // Choose edge randomly
    int32_t chosen_edge = (int32_t)(gsl_rng_get (r) % (unsigned long int)edges);
    int32_t err;
    igraph_vector_t added_edges;
    // Choose a vertex
    int32_t chosen_vertex = (int32_t)(gsl_rng_get (r) % (unsigned long int)vertices);
    int32_t i = 0;
    int32_t small = 0;
    igraph_integer_t v[3];

    // Find edge endpoints
    igraph_edge(graph, chosen_edge, &v[0], &v[1]);
    v[2] = igraph_vcount(graph);
    // Add a vertex
    err = igraph_add_vertices(graph, 1, NULL);
    // Delete old edge
    igraph_delete_edges(graph, igraph_ess_1(chosen_edge));
    if (v[0] > v[1]) {
        small = 1;
    }
    
    for (i = 0; i < 2; i++) {
        // Skip over v0
        if (chosen_vertex >= v[(small + i) % 2]) {
            chosen_vertex = (chosen_vertex + 1) % vertices;  
        }
        
    }
    
    // Add new edges
    igraph_vector_init (&added_edges, 3 * 2);
    for (i = 0; i < 2; i++) {
        VECTOR(added_edges)[i * 2] = v[i]; VECTOR(added_edges)[(i * 2) + 1] = v[2];
    }
    VECTOR(added_edges)[4] = chosen_vertex; VECTOR(added_edges)[5] = v[2];

    igraph_add_edges (graph, &added_edges, NULL);
    igraph_vector_destroy (&added_edges);
}

// Split two edges by adding vertices, connect the two new vertices
void addTwoVertices (igraph_t* graph) {
    // # of edges
    int32_t edges = igraph_ecount(graph);
    // # of vertices
    int32_t vertices = igraph_vcount(graph);
    int32_t i = 0;
    igraph_integer_t new_vertices[2] = {vertices, vertices + 1};
    igraph_add_vertices(graph, 2, NULL);
    igraph_vector_t added_edges;
    igraph_vector_t deleted_edges;
    // Choose a vertex
    igraph_integer_t v[2];
    int32_t smaller = 0;
    igraph_vector_init (&deleted_edges, 2);
    // Choose edge randomly
    VECTOR(deleted_edges)[0] = (int32_t)(gsl_rng_get (r) % (unsigned long int)edges);
    VECTOR(deleted_edges)[1] = (int32_t)(gsl_rng_get (r) % (unsigned long int)edges);
    if (VECTOR(deleted_edges)[0] > VECTOR(deleted_edges)[1]) {
        smaller = 1;
    }
    VECTOR(deleted_edges)[smaller] = ((int32_t)(VECTOR(deleted_edges)[1]) + 1) % edges;
    igraph_vector_init (&added_edges, 5 * 2);
    for (i = 0; i < 2; i++) {
        igraph_edge(graph, VECTOR(deleted_edges)[i], &v[0], &v[1]);
        VECTOR(added_edges)[i*4 + 0] = v[0]; VECTOR(added_edges)[i*4 + 1] = new_vertices[i];
        VECTOR(added_edges)[i*4 + 2] = v[1]; VECTOR(added_edges)[i*4 + 3] = new_vertices[i];
    }
    igraph_delete_edges(graph, igraph_ess_vector(&deleted_edges));
    VECTOR(added_edges)[8] = new_vertices[0]; VECTOR(added_edges)[9] = new_vertices[1];
    igraph_add_edges (graph, &added_edges, NULL);
    igraph_vector_destroy (&deleted_edges);
    igraph_vector_destroy (&added_edges);
}

// Add a new edge between two vertexes. This might result in a vertex connecting two already connected
// vertexes.

void addNewEdge (igraph_t* graph) {
    // # of edges
    // # of vertices
    int32_t vertices = igraph_vcount(graph);
    int32_t i = 0;
    igraph_integer_t verticesToConnect[2];
    igraph_vector_t added_edges; 
    verticesToConnect[0] = (int32_t)(gsl_rng_get (r) % (unsigned long int)vertices);
    verticesToConnect[1] = (int32_t)(gsl_rng_get (r) % (unsigned long int)vertices);
    if (verticesToConnect[0] == verticesToConnect[1]) {
        verticesToConnect[1] = (verticesToConnect[1] + 1) % vertices;
    }
    igraph_vector_init(&added_edges, 2);
    for (i = 0; i < 2; i++) {
        VECTOR(added_edges)[i] = verticesToConnect[i];
    }
    igraph_add_edges (graph, &added_edges, NULL);
    igraph_vector_destroy (&added_edges);
}

// Generate random graph using BG-operations
void evolve3ConnectedGraph (igraph_t* graph, uint32_t verticesToAdd) {
    InitRng();
    uint32_t added = 0;
    while (added < verticesToAdd) {
        int32_t operation = (int32_t)(gsl_rng_get(r) % 3);
        printf("%d", operation);
        switch (operation) {
            case 0:
                splitEdge (graph);
                added += 1;
                break;
            case 1:
                addTwoVertices (graph);
                added += 2;
                break;
            case 2:
                addNewEdge (graph);
                break;
        };
    }
    printf("\n");
}

void construct3ConnectedGraph (igraph_t* graph, uint32_t vertices) {
    // Start with K4
    int err = igraph_full (graph,
                 4, // Vertices
                 false, // Undirected
                 false); // No self loops
    if (err != 0) {
        printf ("Failed to construct graph\n");
        return;
    }
    evolve3ConnectedGraph (graph, vertices - 4);
}

