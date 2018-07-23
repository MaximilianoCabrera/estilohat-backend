package actions

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MaximilianoCabrera/estilohat/estilohat-backend/models"
)

//ProductoGet .
func ProductoGet(w http.ResponseWriter, r *http.Request) {
	var prod models.Producto

	a := r.URL.Query()
	prod.Nombre = a.Get("nombre")
	prod.Descripcion = a.Get("descripcion")
	prod.IDCategoria = 1
	prod.IDTipoProducto = 1
	prod.IDProveedor = 1
	prod.ImagenProducto = a.Get("imagenProducto")

	if prod.Nombre != "" ||
		prod.Descripcion != "" ||
		prod.IDCategoria != 0 ||
		prod.IDTipoProducto != 0 ||
		prod.IDProveedor != 0 ||
		prod.ImagenProducto != "" {
		var x models.GlobalModel

		x.Producto = prod

		globalDAO := globalDAO()
		fmt.Println("Get One Product")
		a, err := globalDAO.GetBy(x, "producto")
		if err != nil {
			fmt.Println("ERROR EN Get One Product")
			responses(w, 404, a, "producto", err)
		}

		for _, prods := range a.Producto {
			log.Println(prods)
		}
		responses(w, 200, a, "producto", nil)
	}
}
