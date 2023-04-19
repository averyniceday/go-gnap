package main

import (
	pubgn "github.com/averyniceday/go-mpath-proto/genome-nexus-public-api"
   	"fmt"
	gognap "github.com/averyniceday/go-gnap/gognap"
)

func main() {
	genomicLocation1 := pubgn.GenomicLocation{}
	genomicLocation1.Chromosome = "17"
	genomicLocation1.Start = 7577121
	genomicLocation1.End = 7577121
	genomicLocation1.ReferenceAllele = "G"
	genomicLocation1.VariantAllele = "A"
	genomicLocations := []pubgn.GenomicLocation{genomicLocation1}
	variantAnnotations := gognap.GetVariantAnnotations(genomicLocations)
	fmt.Println(variantAnnotations[0].SuccessfullyAnnotated)
}

