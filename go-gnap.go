package main

import "fmt"
import (
	pubgn "github.com/averyniceday/go-mpath-proto/genome-nexus-public-api"
)

func main() {
	testloc := pubgn.GenomicLocation{"X", 1000, 2000, "AC", "AG"}
	fmt.Println(testloc.Chromosome)
}

