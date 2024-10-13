package openapi3_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestIssue1012(t *testing.T) {
	spec := `
openapi: 3.0.1
info:
    description: Test API
    title: Test API
    version: "1"
servers:
    - description: Some backend server
      url: https://some.backend.server
paths:
    /v1/test:
        post:
            description: Test endpoint
            requestBody:
                content:
                    application/json:
                        schema:
                            type: object
                            properties:
                                ID:
                                    $ref: '#/components/schemas/ID'
                            required:
                              - ID
                        example:
                            ID:
                                someOtherId: test 
                required: true
            responses:
                "200":
                  description: success
components:
    schemas:
        ID:
            description: Some ID
            properties:
                someId:
                    type: string
                    readOnly: true
                someOtherId:
                  type: string
            required:
                - someId
                - someOtherId
            type: object
`[1:]

	sl := openapi3.NewLoader()
	doc, err := sl.LoadFromData([]byte(spec))
	require.NoError(t, err)
	require.NotNil(t, doc.Paths)

	err = doc.Validate(sl.Context)
	require.NoError(t, err)
	require.NotNil(t, doc.Paths)
}
