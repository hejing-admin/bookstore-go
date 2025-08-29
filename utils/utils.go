package utils

import (
	"bookstore-go/web/response"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"reflect"
)

// Bind creates a middleware that binds request data to a struct.
// The struct must be passed by value (not a pointer).
// It supports JSON, XML, and form data based on Content-Type.
// The bound object is stored in context under gin.BindKey.
func Bind(obj interface{}) gin.HandlerFunc {
	// Get the reflective value of the input object
	value := reflect.ValueOf(obj)

	// Check if the input is a pointer type
	if value.Kind() == reflect.Ptr {
		// Log fatal error if pointer is passed (should be value type)
		log.Fatalf(`Bind struct can not be a pointer. Example:
        Use: gin.ParamBind(Struct{}) instead of gin.ParamBind(&Struct{})`)
	}

	// Get the type information of the input struct
	typ := value.Type()

	// Return the handler function that will be executed for each request
	return func(c *gin.Context) {
		// Create a new instance of the struct (as pointer) for binding
		obj := reflect.New(typ).Interface()

		// Get the Content-Type header from the request
		contentType := c.ContentType()

		// Set default content type to JSON if not specified
		if contentType == "" {
			contentType = "application/json"
		}

		// Get the appropriate binder based on method and content type
		b := binding.Default(c.Request.Method, contentType)

		// Bind the request data to the newly created struct instance
		err := b.Bind(c.Request, obj)

		// Check if binding encountered any errors
		if err != nil {
			// Return bad request response if binding fails
			response.BadRequest(c, "something is wrong about request body", err)
			// Stop further processing in the handler chain
			return
		}

		// Store the successfully bound object in context for later use
		c.Set(gin.BindKey, obj)
	}
}
