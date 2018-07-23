package models

//Producto ...
type Producto struct {
	IDProducto     int    `json: "idProducto"`
	Nombre         string `json: "nombre"`
	Descripcion    string `json: "descripcion"`
	IDCategoria    int    `json: "idCategoria"`
	IDTipoProducto int    `json: "idTipo"`
	IDProveedor    int    `json: "idProveedor"`
	ImagenProducto string `json: "imagenProducto"`
}

//TipoProducto ...
type TipoProducto struct {
	IDTipoProducto          int //`json: "idProducto"`
	NombreTipoProducto      string
	DescripcionTipoProducto string
}

//Proveedor .
type Proveedor struct {
	IDProveedor      int
	NombreProveedor  string
	PuntajeProveedor float32
}

//Repuesto - ver nombre ...
type Repuesto struct {
	IDProducto int //`json: "idProducto"`
}

//Herramienta .
type Herramienta struct {
	IDHerramienta int
	IDProducto    int //`json: "idProducto"`
	IDEstado      int //1-Nuevo, 2-usado, 3-gastado, 4-repuesto, 5-roto
}

//Asignacion ...
type Asignacion struct {
	IDAsignacion int
	IDProducto   int
	IDEmpleado   int
}

//Empleado ...
type Empleado struct {
	IDEmpleado       int
	NombreEmpleado   string
	ApellidoEmpleado string
	IDDireccion      int
}

//Obrero ...
type Obrero struct {
	IDObrero   int
	IDEmpleado int
	IDPuesto   int
}

//Categoria ...
type Categoria struct {
	IDCategoria          int    //`json: "idCategoria"`
	NombreCategoria      string //`json: "nombreCategoria"`
	DescripcionCategoria string //`json: "descripcionCategoria"`
}

//Stock ...
type Stock struct {
	IDStock    int     //`json: "idStock"`
	IDProducto int     //`json: "idProducto"`
	Cantidad   float32 //`json: "cantidad"`

}
