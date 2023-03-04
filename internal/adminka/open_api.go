package adminka

import "github.com/getkin/kin-openapi/openapi3"

//go:generate go run openapi-gen/main.go -path .
//go:generate swagger-codegen generate -i openapi3.yaml  -l html2 -o static/

func NewOpenAPI3() openapi3.T {
	swagger := openapi3.T{
		OpenAPI: "3.0.0",
		Info: &openapi3.Info{
			Title:       "NetRadio-Adminka API",
			Description: "REST APIs used for interacting with Adminka Service",
			Version:     "0.0.0",
			Contact: &openapi3.Contact{
				URL: "https://github.com/4el0ve4ek/netradio",
			},
		},
		Servers: openapi3.Servers{
			&openapi3.Server{
				Description: "Local development",
				URL:         "http://localhost:8080",
			},
			&openapi3.Server{
				Description: "Docker development",
				URL:         "http://localhost:8080",
			},
		},
	}

	swagger.Components = &openapi3.Components{}

	swagger.Components.Schemas = openapi3.Schemas{
		"Id": openapi3.NewSchemaRef("",
			openapi3.NewIntegerSchema()),
		"Title": openapi3.NewSchemaRef("",
			openapi3.NewStringSchema().WithNullable()),
		"Content": openapi3.NewSchemaRef("",
			openapi3.NewStringSchema().WithNullable()),
	}

	swagger.Components.RequestBodies = openapi3.RequestBodies{
		"CreateNewsRequest": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("Request used for creating a news.").
				WithRequired(true).
				WithJSONSchema(openapi3.NewSchema().
					WithPropertyRef("title", &openapi3.SchemaRef{
						Ref: "#/components/schemas/Title",
					}).
					WithPropertyRef("content", &openapi3.SchemaRef{
						Ref: "#/components/schemas/Content",
					})),
		},
		"ChangeNewsRequest": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("Request used for changing a news.").
				WithRequired(true).
				WithJSONSchema(openapi3.NewSchema().
					WithPropertyRef("id", &openapi3.SchemaRef{
						Ref: "#/components/schemas/Id",
					}).
					WithPropertyRef("title", &openapi3.SchemaRef{
						Ref: "#/components/schemas/Title",
					}).
					WithPropertyRef("content", &openapi3.SchemaRef{
						Ref: "#/components/schemas/Content",
					})),
		},
	}

	swagger.Components.Responses = openapi3.Responses{
		"ErrorResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response when errors happen.").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewSchema().
					WithProperty("error", openapi3.NewStringSchema()))),
		},
	}

	swagger.Paths = openapi3.Paths{
		"/news/add": &openapi3.PathItem{
			Post: &openapi3.Operation{
				OperationID: "Create News",
				RequestBody: &openapi3.RequestBodyRef{
					Ref: "#/components/requestBodies/CreateNewsRequest",
				},
				Responses: openapi3.Responses{
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
					"200": &openapi3.ResponseRef{
						Value: openapi3.NewResponse().WithDescription("News created"),
					},
				},
			},
		},
		"/news/change": &openapi3.PathItem{
			Post: &openapi3.Operation{
				OperationID: "Change News",
				RequestBody: &openapi3.RequestBodyRef{
					Ref: "#/components/requestBodies/ChangeNewsRequest",
				},
				Responses: openapi3.Responses{
					"200": &openapi3.ResponseRef{
						Value: openapi3.NewResponse().WithDescription("News changed"),
					},
					"404": &openapi3.ResponseRef{
						Value: openapi3.NewResponse().WithDescription("News not found"),
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
				},
			},
		},
	}

	return swagger
}
