package main

import (
	pubgn "github.com/averyniceday/go-mpath-proto/genome-nexus-public-api"
   	"fmt"
	"context"
    	"os"
)

func main() {
	testloc := pubgn.GenomicLocation{"X", 1000, 2000, "AC", "AG"}
	fmt.Println(testloc.Chromosome)
	genomicLocation := "7,140453136,140453136,A,T"
	token := ""
	isoformOverrideSource := ""
	fields := []string{"hotspots"}

	configuration := pubgn.NewConfiguration()
    	apiClient := pubgn.NewAPIClient(configuration)
    
	resp, r, err := apiClient.AnnotationControllerApi.FetchVariantAnnotationByGenomicLocationGET(context.Background(), genomicLocation).IsoformOverrideSource(isoformOverrideSource).Token(token).Fields(fields).Execute()
    	if err != nil {
        	fmt.Fprintf(os.Stderr, "Error when calling `AnnotationControllerApi.FetchVariantAnnotationByGenomicLocationGET``: %v\n", err)
        	fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    	}
    	// response from `FetchVariantAnnotationByGenomicLocationGET`: VariantAnnotation
    	fmt.Println(r)
	fmt.Println(resp)
}

func GetVariantAnnotation(genomicLocation string) {
	fmt.Println("Annotation for location " + genomicLocation)
	token := ""
	isoformOverrideSoruce := ""
	fields := []string{"hotspots"}
	configuration := pubgn.NewConfiguration()
    	apiClient := pubgn.NewAPIClient(configuration)
    
	_, r, err := apiClient.AnnotationControllerApi.FetchVariantAnnotationByGenomicLocationGET(context.Background(), genomicLocation).IsoformOverrideSource(isoformOverrideSource).Token(token).Fields(fields).Execute()
    	if err != nil {
        	fmt.Fprintf(os.Stderr, "Error when calling `AnnotationControllerApi.FetchVariantAnnotationByGenomicLocationGET``: %v\n", err)
        	fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    	}
    	// response from `FetchVariantAnnotationByGenomicLocationGET`: VariantAnnotation
    	fmt.Println(r)
}	
