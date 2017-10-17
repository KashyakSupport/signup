package desgin

import . "github.com/goadesign/goa/design"
import . "github.com/goadesign/goa/design/apidsl"

var _ = API("signup", func() {
	Description("The User signup service")
	Host("localhost:8080")
})

//UserPayload is a inputtype/Request
var UserPayload = Type("UserPayload", func() {
	Description("UserPayload is type used to create users")

	Attribute("FirstName", String, "FirstName of the user", func() {
		MinLength(2)
	})

	Required("FirstName")
})

//UserMedia is a outputtype/Response
var UserMedia = MediaType("application/vnd.goa.example.user+json", func() {

	TypeName("user")
	Reference(UserPayload)

	Attributes(func() {
		Attribute("ID", Integer, "Unique bottle ID")
		Attribute("FirstName")

		Required("ID", "FirstName")
	})
	View("default", func() {
		Attribute("ID")
		Attribute("FirstName")

	})

})

var _ = Resource("user", func() {
	Description("A user Account")
	BasePath("/signup")

	Action("create", func() {
		Description("creates a user")
		Routing(POST("/"))
		Payload(UserPayload)
		Response(Created)
	})

	Action("show", func() {
		Description("shows a user")
		Routing(GET("/:ID"))
		Params(func() {
			Param("ID", Integer)
		})
		Response(OK, UserMedia)
	})
})

var _ = Resource("swagger", func() {
	Description("The API Swagger specification")
	Files("/swagger.json", "swagger/swagger.json")
	Files("/swagger-ui/*filepath", "swagger-ui/")
})
