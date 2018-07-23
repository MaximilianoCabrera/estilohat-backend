package factory

import (
	"github.com/MaximilianoCabrera/estilohat/estilohat-backend/dao/mysql"
	"github.com/maximilianoCabrera/estilohat/estilohat-backend/dao/interfaces"
)

func GlobalFactoryDAO(e string) interfaces.GlobalDAO {
	var i interfaces.GlobalDAO
	i = mysql.GlobalImplMysql{}
	return i
}
