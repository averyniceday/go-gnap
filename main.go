package main

import (
	pubgn "github.com/averyniceday/go-mpath-proto/genome-nexus-public-api"
   	"fmt"
	gognap "github.com/averyniceday/go-gnap/gognap"
)

func main() {
	genomicLocation1 := pubgn.NewGenomicLocation("7", 140453136, 140453136, "A", "T")
	genomicLocation2 := pubgn.NewGenomicLocation("17", 7577121, 7577121, "G", "A")
	genomicLocation3 := pubgn.NewGenomicLocation("X", 1000, 2000, "AC", "AG")
	genomicLocations := []pubgn.GenomicLocation{*genomicLocation1, *genomicLocation2, *genomicLocation3}
	variantAnnotations := gognap.GetVariantAnnotations(genomicLocations)
	fmt.Println(*variantAnnotations[0].SuccessfullyAnnotated)
}

