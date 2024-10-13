package openapi3

import "context"

func validateExampleValue(ctx context.Context, input any, schema *Schema) error {
	opts := make([]SchemaValidationOption, 0, 2)

	// FIXME: for some reason ctx has not validation options
	// so we dont validate as neither request nor response
	if vo := getValidationOptions(ctx); vo.examplesValidationAsReq {
		opts = append(opts, VisitAsRequest())
	} else if vo.examplesValidationAsRes {
		opts = append(opts, VisitAsResponse())
	}
	opts = append(opts, MultiErrors())

	return schema.VisitJSON(input, opts...)
}
