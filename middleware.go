package ginpongo2

import (
	pongo2 "gopkg.in/flosch/pongo2.v3"
	. "github.com/gin-gonic/gin"
	"net/http"
)

func Pongo2() HandlerFunc {
	return func(c *Context) {
		c.Next()

		templateName, templateNameExists := c.Get("template")
		templateData, templateDataExists := c.Get("data")
		templateNameValue, isString := templateName.(string)

		if templateNameExists && templateDataExists && isString {
			template := pongo2.Must(pongo2.FromFile(templateNameValue))
			err := template.ExecuteWriter(getContext(templateData), c.Writer)
			if err != nil {
				http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}

func getContext(templateData interface{}) pongo2.Context {
	if templateData == nil {
		return nil
	}
	contextData, isMap := templateData.(map[string]interface{})
	if isMap {
		return contextData
	}
	return nil
}
