package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"io/ioutil"
	"github.com/xeipuuv/gojsonschema"
)

func postHandler(context *gin.Context) {
	response := ResponseMessage{}
	response.Errors = make(map[string]string)

	// checking if there is a json schema for this request
	schemaFile, err := getSchema(context.Param("schema"))
	if err != nil {
		log.Errorf("Schema file %s not found", schemaFile)
		response.SuccessStatus = false
		response.Errors["schema"] = "Schema file not found for this endpoint"
		context.JSON(http.StatusInternalServerError, response)

		return
	}

	// reading request body
	requestBody, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		log.Errorf("Can not get request. Error %s", err)
		response.Errors["content"] = err.Error()
		context.JSON(http.StatusInternalServerError, response)

		return
	}

	schemaLoader := gojsonschema.NewReferenceLoader(fmt.Sprintf("file://%s", schemaFile))
	requestJsonLoader := gojsonschema.NewBytesLoader(requestBody)

	result, err := gojsonschema.Validate(schemaLoader, requestJsonLoader)

	if err != nil {
		log.Errorf("Can not validate json")
		response.Errors["invalidJson"] = err.Error()
		context.JSON(http.StatusInternalServerError, response)

		return
	}

	if result.Valid() == true {
		response.SuccessStatus = true
		context.JSON(http.StatusOK, response)

		return
	} else {
		response.SuccessStatus = false
		for _, err := range result.Errors() {
			// Err implements the ResultError interface
			response.Errors[err.Field()] = err.Description()
		}

		context.JSON(http.StatusBadRequest, response)
	}
}
