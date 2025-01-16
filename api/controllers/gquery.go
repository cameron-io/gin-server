package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

func NewGQueryController(r *gin.RouterGroup) {
	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    GQueryType(),
			Mutation: GMutationType(),
		},
	)
	r.POST("/graphql", func(ctx *gin.Context) {
		var jsonData map[string]interface{}
		data, _ := io.ReadAll(ctx.Request.Body)
		if e := json.Unmarshal(data, &jsonData); e != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": e.Error()})
			return
		}
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: jsonData["query"].(string),
		})
		if len(result.Errors) > 0 {
			fmt.Printf("errors: %v", result.Errors)
		}
		ctx.JSON(http.StatusOK, result)
	})
}

func GQueryType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Message",
			Fields: graphql.Fields{
				"echo": &graphql.Field{
					Type:        graphql.String,
					Description: "Echo's input",
					Args: graphql.FieldConfigArgument{
						"message": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						message, isOK := params.Args["message"].(string)
						if isOK {
							return message, nil
						}
						return "no message found", nil
					},
				},
			},
		})
}

func GMutationType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Sum",
			Fields: graphql.Fields{
				"sum": &graphql.Field{
					Type:        graphql.Int,
					Description: "Sum two numbers",
					Args: graphql.FieldConfigArgument{
						"x": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
						"y": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: func(params graphql.ResolveParams) (interface{}, error) {
						x, _ := params.Args["x"].(int)
						y, _ := params.Args["y"].(int)
						return x + y, nil
					},
				},
			},
		})
}
