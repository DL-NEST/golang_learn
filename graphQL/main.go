package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"log"
	"net/http"
)

// QueryRequest GraphQL 查询语法字段
type QueryRequest struct {
	Query     string                 `json:"query" form:"query"`          //GraphQL query 语句
	Variables map[string]interface{} `json:"variables"  form:"variables"` //变量
}

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	},
})

func main() {
	srv := gin.Default()
	srv.GET("/graphql", func(ctx *gin.Context) {

		var reqData QueryRequest
		if err := ctx.Bind(&reqData); err != nil {
			return
		}

		log.Printf("aa%v\n\n", reqData)

		var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
			Query: queryType,
		})
		params := graphql.Params{
			Schema:         Schema,
			RequestString:  reqData.Query,
			Context:        ctx,
			VariableValues: reqData.Variables,
		}
		r := graphql.Do(params)
		if len(r.Errors) > 0 {
			log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
		}
		rJSON, _ := json.Marshal(r)
		fmt.Printf("%s \n", rJSON)
		ctx.JSON(http.StatusOK, r)
		return
	})
	err := srv.Run(":9099")
	if err != nil {
		return
	}

}
