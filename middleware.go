package ginpongo2

import (
	pongo2 "gopkg.in/flosch/pongo2.v3"
	. "github.com/gin-gonic/gin"
	"net/http"
)

func Pongo2() HandlerFunc {
	return func(c *Context) {
		c.Next()

		templateName, templateExists := c.Get("template")
		templateNameValue, isString := templateName.(string)

		if templateExists && isString {
			templateData, templateDataExists := c.Get("data")
			var template = pongo2.Must(pongo2.FromFile(templateNameValue))
			err := template.ExecuteWriter(getContext(templateData, templateDataExists), c.Writer)
			if err != nil {
				http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}

func getContext(templateData interface{}, exists bool) pongo2.Context {
	if templateData == nil || !exists {
		return nil
	}
	contextData, isMap := templateData.(map[string]interface{})
	if isMap {
		return contextData
	}
	return nil
}
