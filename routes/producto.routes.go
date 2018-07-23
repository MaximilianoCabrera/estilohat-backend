package routes

import "github.com/MaximilianoCabrera/estilohat/estilohat-backend/actions"

//Producto .
var producto = Routes{
	//ProductoGet
	Route{
		"productoGet",
		"GET",
		"/producto/{id}",
		actions.ProductoGet
	},
}
