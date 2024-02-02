package apirest

import (
	"github.com/getkin/kin-openapi/openapi3"
)

type OpenAPI3 struct {
	doc *openapi3.T
}

func NewOpenAPI3() *OpenAPI3 {
	doc := &openapi3.T{
		OpenAPI: "3.0.0",
		Paths:   &openapi3.Paths{},
		Info: &openapi3.Info{
			Title:       "Golang",
			Description: "REST APIs used for interacting with the ToDo Service.",
			Version:     "1.0.0",
			License: &openapi3.License{
				Name: "MIT",
				URL:  "https://opensource.org/licenses/MIT",
			},
			Contact: &openapi3.Contact{
				Name:  "Jorge Luis Alonso",
				Email: "alonso12.dev@gmail.com",
				URL:   "https://github.com/alonsojl",
			},
		},
		Servers: openapi3.Servers{
			&openapi3.Server{
				URL:         "{Scheme}://{Host}:{Port}/api/v1",
				Description: "URL endpoints",
				Variables: map[string]*openapi3.ServerVariable{
					"Scheme": {
						Enum:    []string{"http", "https"},
						Default: "http",
					},
					"Host": {
						Enum:    []string{"localhost", "192.168.1.5"},
						Default: "localhost",
					},
					"Port": {
						Enum:    []string{"8000", "8080"},
						Default: "8000",
					},
				},
			},
		},
		Tags: openapi3.Tags{
			&openapi3.Tag{
				Name:        "Users",
				Description: "Endpoints to manage users.",
				ExternalDocs: &openapi3.ExternalDocs{
					Description: "external docs description",
					URL:         "https://swagger.io/specification/",
				},
			},
			&openapi3.Tag{
				Name:        "Login",
				Description: "Endpoint to login users.",
				ExternalDocs: &openapi3.ExternalDocs{
					Description: "external docs description",
					URL:         "https://swagger.io/specification/",
				},
			},
		},
		ExternalDocs: &openapi3.ExternalDocs{
			Description: "Find out more about OpenAPI Specification.",
			URL:         "https://swagger.io/specification/",
		},
		Components: &openapi3.Components{
			SecuritySchemes: openapi3.SecuritySchemes{
				"bearerAuth": &openapi3.SecuritySchemeRef{
					Value: openapi3.NewJWTSecurityScheme(),
				},
			},
			Schemas:       make(openapi3.Schemas),
			RequestBodies: make(openapi3.RequestBodies),
			Responses:     make(openapi3.ResponseBodies),
		},
		Security: openapi3.SecurityRequirements{
			openapi3.SecurityRequirement{
				"bearerAuth": []string{},
			},
		},
	}

	return &OpenAPI3{
		doc: doc,
	}
}

func (api *OpenAPI3) WithUser() *OpenAPI3 {
	// Schema
	api.doc.Components.Schemas["User"] = openapi3.NewSchemaRef("", openapi3.NewSchema().
		WithProperty("name", openapi3.NewStringSchema()).
		WithProperty("first_name", openapi3.NewStringSchema()).
		WithProperty("last_name", openapi3.NewStringSchema()).
		WithProperty("email", openapi3.NewStringSchema()).
		WithProperty("phone", openapi3.NewStringSchema()).
		WithProperty("age", openapi3.NewInt32Schema()))
	// Request
	api.doc.Components.RequestBodies["CreateUsersRequest"] = &openapi3.RequestBodyRef{
		Value: openapi3.NewRequestBody().
			WithDescription("Request used for creating a user.").
			WithRequired(true).
			// WithJSONSchemaRef(&openapi3.SchemaRef{Ref: "#/components/schemas/User"}),
			WithJSONSchema(openapi3.NewSchema().
				WithProperty("name", openapi3.NewStringSchema().WithDefault("Jorge Luis")).
				WithProperty("first_name", openapi3.NewStringSchema().WithDefault("Alonso")).
				WithProperty("last_name", openapi3.NewStringSchema().WithDefault("Hdez")).
				WithProperty("email", openapi3.NewStringSchema().WithDefault("alonso12.dev@gmail.com")).
				WithProperty("phone", openapi3.NewStringSchema().WithDefault("7713037204")).
				WithProperty("age", openapi3.NewInt32Schema().WithDefault("25"))),
	}
	// Response
	api.doc.Components.Responses["GetUsersResponse"] = &openapi3.ResponseRef{
		Value: openapi3.NewResponse().
			WithDescription("Response returned back after getting all users.").
			WithJSONSchema(openapi3.NewSchema().
				WithProperty("status", openapi3.NewStringSchema()).
				WithProperty("http_code", openapi3.NewInt32Schema()).
				WithProperty("datetime", openapi3.NewStringSchema()).
				WithProperty("timestamp", openapi3.NewInt64Schema()).
				WithProperty("user", openapi3.NewArraySchema().WithItems(&openapi3.Schema{
					Items: &openapi3.SchemaRef{
						Ref: "#/components/schemas/User",
					},
				}))),
	}
	api.doc.Components.Responses["CreateUsersResponse"] = &openapi3.ResponseRef{
		Value: openapi3.NewResponse().
			WithDescription("Response returned back after creating users.").
			WithJSONSchema(openapi3.NewSchema().
				WithProperty("status", openapi3.NewStringSchema()).
				WithProperty("http_code", openapi3.NewInt32Schema()).
				WithProperty("datetime", openapi3.NewStringSchema()).
				WithProperty("timestamp", openapi3.NewInt64Schema()).
				WithPropertyRef("user", &openapi3.SchemaRef{Ref: "#/components/schemas/User"})),
	}
	// Paths
	api.doc.Paths.Set("/users", &openapi3.PathItem{
		Get: &openapi3.Operation{
			Summary:     "Get all users.",
			Description: "Gets all user rows.",
			Tags:        []string{"Users"},
			OperationID: "GetUsers",
			Responses: openapi3.NewResponses(
				openapi3.WithStatus(200, &openapi3.ResponseRef{
					Ref: "#/components/responses/GetUsersResponse",
				}),
			),
		},
		Post: &openapi3.Operation{
			Summary:     "Create user.",
			Description: "Register a user with the requested fields.",
			Tags:        []string{"Users"},
			OperationID: "CreateUser",
			RequestBody: &openapi3.RequestBodyRef{
				Ref: "#/components/requestBodies/CreateUsersRequest",
			},
			Responses: openapi3.NewResponses(
				openapi3.WithStatus(201, &openapi3.ResponseRef{
					Ref: "#/components/responses/CreateUsersResponse",
				}),
			),
		},
	})

	return api
}

func (api *OpenAPI3) Generate() *openapi3.T {
	return api.doc
}
