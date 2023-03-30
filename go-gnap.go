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
	//genomicLocation := "7,140453136,140453136,A,T"
	//genomicLocations := []string{"7,140453136,140453136,A,T"}
	// func NewGenomicLocation(chromosome string, start int32, end int32, referenceAllele string, variantAllele string) *GenomicLocation {
	genomicLocation1 := pubgn.NewGenomicLocation("7", 140453136, 140453136, "A", "T")
	genomicLocation2 := pubgn.NewGenomicLocation("17", 7577121, 7577121, "G", "A")
	genomicLocation3 := pubgn.NewGenomicLocation("X", 1000, 2000, "AC", "AG")
	genomicLocations := []pubgn.GenomicLocation{*genomicLocation1, *genomicLocation2, *genomicLocation3}
	GetVariantAnnotation(genomicLocations)
	/*token := ""
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
	*/
}

// TODO add summary statistics
// Java call: annotator.getAnnotatedRecordsUsingPOST(summaryStatistics, records, "mskcc", true, postIntervalSize, reannotate);
func GetVariantAnnotation(genomicLocations []pubgn.GenomicLocation) {
	token := ""
	isoformOverrideSource := ""
	fields := []string{"hotspots"}
	configuration := pubgn.NewConfiguration()
    	apiClient := pubgn.NewAPIClient(configuration)
    
	variantAnnotations, r, err := apiClient.AnnotationControllerApi.FetchVariantAnnotationByGenomicLocationPOST(context.Background()).GenomicLocations(genomicLocations).IsoformOverrideSource(isoformOverrideSource).Token(token).Fields(fields).Execute()
    	if err != nil {
        	fmt.Fprintf(os.Stderr, "Error when calling `AnnotationControllerApi.FetchVariantAnnotationByGenomicLocationGET``: %v\n", err)
        	fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    	}
	// https://github.com/averyniceday/go-mpath-proto/blob/16d6085cee926ad85fc29c6c62319ee3697edcf2/genome-nexus-public-api/api_annotation_controller.go#L218
    	// response from `FetchVariantAnnotationByGenomicLocationPOST`: []VariantAnnotation, *http.Response, error
	// VariantAnnotation defined here: https://github.com/averyniceday/go-mpath-proto/blob/16d6085cee926ad85fc29c6c62319ee3697edcf2/genome-nexus-internal-api/model_variant_annotation.go#L21
    	fmt.Println(r)
	for i := 0; i < len(variantAnnotations); i++ {
		fmt.Println("Success: ", *variantAnnotations[i].SuccessfullyAnnotated)
		if !*variantAnnotations[i].SuccessfullyAnnotated {
			fmt.Println("ERROR: Failed to annotate: ", variantAnnotations[i].Variant)
		}
	}
}	
