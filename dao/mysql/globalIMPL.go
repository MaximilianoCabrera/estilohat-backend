package mysql

import (
	"fmt"
	"log"

	"github.com/MaximilianoCabrera/estilohat/estilohat-backend/models"
)

type GlobalImplMysql struct{}

func checkErr(model string, accion string, err error) {
	if err != nil {
		log.Println("No se pudo "+accion+" "+model+". Error desde IMPL: ", err)
	}
}

//Create .
//TODO: Agregar los DB().BEGIN() y COMMIT
func (dao GlobalImplMysql) Create(x *models.GlobalModel, model string) (models.GlobalModel, error) {
	switch model {
	case "producto":
		err := DB().CreateAndRead(&x.Producto)
		checkErr("producto", "crear", err)
	default:
		log.Println("Modelo ingresado no existente")
	}
	return *x, nil
}

//GetAll .
func (dao GlobalImplMysql) GetAll(model string) (models.GlobalModels, error) {
	var a models.GlobalModels

	switch model {
	case "producto":
		err := DB().Read(&a.Producto, "SELECT * FROM producto")
		checkErr("producto", "getAll", err)
	default:
		log.Println("Modelo ingresado no existente")
	}
	return a, nil
}

//GetByID .
func (dao GlobalImplMysql) GetByID(id int, model string) (models.GlobalModel, error) {
	var x models.GlobalModel

	switch model {
	case "producto":
		err := DB().Read(&x.Producto, "SELECT * FROM producto WHERE id = ?", id)
		checkErr("producto", "getByID", err)
	default:
		log.Println("Modelo ingresado no existente")
	}
	return x, nil
}

//GetBy .
func (dao GlobalImplMysql) GetBy(x models.GlobalModel, model string) (models.GlobalModels, error) {
	a := models.GlobalModels{}
	switch model {
	case "producto":
		err := DB().Read(&a.Producto, "SELECT * FROM producto WHERE nombre = ? OR descripcion = ? OR idCategoria = ? OR idTipoProducto = ? OR idProveedor = ? OR imagenProducto = ?", x.Producto.Nombre, x.Producto.Descripcion, x.Producto.IDCategoria, x.Producto.IDTipoProducto, x.Producto.IDProveedor, x.Producto.ImagenProducto)
		checkErr("producto", "getOne", err)
	default:
		log.Println("Modelo ingresado no existente")
	}
	fmt.Println("MSJ: ", a)
	return a, nil
}

//Update .
func (dao GlobalImplMysql) Update(x models.GlobalModel, model string) (models.GlobalModel, error) {
	switch model {
	case "producto":
		var prod models.Producto

		err := DB().Read(&prod, "SELECT * FROM producto WHERE id = ?", x.Producto.IDProducto)
		checkErr("producto", "getByID", err)

		if x.Producto.Nombre == "" {
			x.Producto.Nombre = prod.Nombre
		}
		if x.Producto.Descripcion == "" {
			x.Producto.Descripcion = prod.Descripcion
		}
		if x.Producto.IDCategoria == 0 {
			x.Producto.IDCategoria = prod.IDCategoria
		}
		if x.Producto.IDTipoProducto == 0 {
			x.Producto.IDTipoProducto = prod.IDTipoProducto
		}
		if x.Producto.IDProveedor == 0 {
			x.Producto.IDProveedor = prod.IDProveedor
		}
		if x.Producto.ImagenProducto == "" {
			x.Producto.ImagenProducto = prod.ImagenProducto
		}

		err = DB().Update(x.Producto)
		checkErr("producto", "actualizar", err)
	default:
		log.Println("Modelo ingresado no existente")
	}
	return x, nil
}

//Delete .
func (dao GlobalImplMysql) Delete(x *models.GlobalModel, model string) (string, error) {
	var err error
	var msjReturn = ""
	switch model {
	case "producto":
		msjReturn = "Producto eliminado correctamente."
		err = DB().Delete(x.Producto)
		checkErr("producto", "eliminar", err)
	default:
		msjReturn = "Modelo ingresado no existente"
	}
	return msjReturn, err
}
