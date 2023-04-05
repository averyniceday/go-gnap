package gognap

import (
	pubgn "github.com/averyniceday/go-mpath-proto/genome-nexus-public-api"
   	"fmt"
	"context"
    	"os"
)


// TODO add summary statistics
// Java call: annotator.getAnnotatedRecordsUsingPOST(summaryStatistics, records, "mskcc", true, postIntervalSize, reannotate);

func GetVariantAnnotations(genomicLocations []pubgn.GenomicLocation) []pubgn.VariantAnnotation{
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
	for i := 0; i < len(variantAnnotations); i++ {
		if !*variantAnnotations[i].SuccessfullyAnnotated {
			fmt.Println("ERROR: Failed to annotate: ", variantAnnotations[i].Variant)
		} else {
			fmt.Println("Success: ", *variantAnnotations[i].SuccessfullyAnnotated)
		}
	}
	return variantAnnotations
}	
