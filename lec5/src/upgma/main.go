package main

import (
	"fmt"
	"os"
	"strconv"
)

type DistanceMatrix [][]float64

type Tree []*Node

// we also think of a cluster as a *Node

type Node struct {
	age            float64
	label          string
	child1, child2 *Node // if at leaf, both will be nil
}

func main() {
	fmt.Println("UPGMA.")

	fileName := os.Args[1]
  mtx, species := ReadMatrixFromFile(fileName)

  var t Tree = UPGMA(mtx, speciesNames) // or t := UPGMA(mtx, speciesNames)

  t.PrintGraphViz()
}

//UPGMA function goes here
func UPGMA(mtx DistanceMatrix, speciesNames []string) Tree {

      t := InitializeTree(speciesNames)
      // will create all nodes needed, and not point any node at any node as a child. 

      clusters := t.InitializeClusters()
      //clusters will start out as just the leaves of t (slice of a node pointers)
      // a cluster is a pointer to a node!

      // now for the UPGMA engine ... apply steps of the algorithm numLeaves - 1 steps

      numLeaves := len(speciesNames)
      for p := numLeaves; p < 2*numLeaves - 1; p++ {
          // first, find minimum element of mtx
          row, col, val := FindMinElement(mtx)
          // big assumption: col > row
          // set the age of current node.
          t[p].age = val/2.0

          // set children of t[p]
          t[p].child1 = clusters[row]
          t[p].child2 = clusters[col]

          // now, update matrix and clusters
          mtx = AddRowCol(mtx, clusters, row, col)

          // add t[p] to the end of clusters
          clusters = append(clusters, t[p])

          // now, hack out row and col from distance matrix and form clusters

          mtx = DelRowCol(mtx, row, col)
          clusters = DelClusters(clusters, row, col)
      }
      // now, we are ready to return tree
      return t
}

//FindMinElement takes a DistanceMatrix and returns the row index, column index, and value corresponding to a minimum element.
//Assumption that col > row
