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
func FindMinElement(mtx DistanceMatrix) (int, int, float64) {
    if len(mtx) <= 1 || len(mtx[0]) <= 0 {
        panic("One row or one column!")
    }

    // assume that matrix is at least 2 x 2
    row := 0
    col := 1
    minVal := mtx[row][col]

    // range over matrix, and see if we can do better than minVal.
    for i := 0; i < len(mtx)-1; i++ {
        // start at column ranging at i + 1
        for j := i+1; j < len(mtx[i]); j++ {
            //do we have a winner?
            if mtx[i][j] < minVal {
                // update all three variables
                minVal = mtx[i][j]
                row = i
                col = j
            }
        }
    }
    return row, col, minVal
}
//DelRowCol takes a distance matrix and a row/col index and deletes the row and column indicated, returning the resulting matrix

//DelClusters takes a slice of Node pointers along with a row/col index and deletes the clusters in the slice corresponding to these indices.
//Assume col > row

//AddRowCol takes a DistanceMatrix, a slice of current clusters, and a row/col index (col > row).
//It returns the matrix corresponding to "gluing" clusters[row] and clusters[col] together and forming a new row/col of the matrix (no deletions yet).

//CountLeaves is a recursive Node function that counts the number of leaves in the tree rooted at the node. It returns 1 at a leaf.

//InitializeClusters is a Tree method that returns a slice of pointers to the leaves of the Tree
func (t Tree) InitializeClusters() []*Node {
    numNodes := len(t)
    numLeaves := (numNodes+1)/2

    clusters := make([]*Node, numLeaves)
    // clusters[i] should point to the i-th leaf node of t
    for i := range clusters {
        clusters[i] = t[i]
    }

    return clusters
}
//InitializeTree takes the n names of our present-day species (leaves) and returns a rooted binary tree with 2n-1 total nodes, where the leaves are the first n and have the associated species names.
func InitializeTree(speciesNames []string) Tree {
    numLeaves := len(speciesNames)
    var t Tree // a Tree is []*Node

    t = make([]*Node, 2*numLeaves-1)
    // all of these pointers have default value of nil; we need to point them at nodes

    // we should create 2n-1 nodes.
    for i := range t {
        var vx Node
        // let's label the first numLeaves nodes with the appropriate species name.
        // by default, vx.age = 0.0, and its children are nil.
        if i < numLeaves {
            //at a leaf ... let's assign its label.
            vx.label = speciesNames[i]
        } else {
            // let's just give it an unspecific name
            vx.label = "Ancestor species " + strconv.Itoa(i)
        }
        // one thing to do: point t[i] at vx
        t[i] = &vx
    }

    return t
}
