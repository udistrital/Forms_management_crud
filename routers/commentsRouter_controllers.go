package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/forms_management_crud/controllers:CampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/forms_management_crud/controllers:CampoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/forms_management_crud/controllers:CampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/forms_management_crud/controllers:CampoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/forms_management_crud/controllers:CampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/forms_management_crud/controllers:CampoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/forms_management_crud/controllers:CampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/forms_management_crud/controllers:CampoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/forms_management_crud/controllers:CampoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/forms_management_crud/controllers:CampoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/forms_management_crud/controllers:FormularioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/forms_management_crud/controllers:FormularioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/forms_management_crud/controllers:FormularioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/forms_management_crud/controllers:FormularioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/forms_management_crud/controllers:FormularioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/forms_management_crud/controllers:FormularioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/forms_management_crud/controllers:FormularioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/forms_management_crud/controllers:FormularioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/forms_management_crud/controllers:FormularioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/forms_management_crud/controllers:FormularioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/forms_management_crud/controllers:TipoEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/forms_management_crud/controllers:TipoEstadoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/forms_management_crud/controllers:TipoEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/forms_management_crud/controllers:TipoEstadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/forms_management_crud/controllers:TipoEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/forms_management_crud/controllers:TipoEstadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/forms_management_crud/controllers:TipoEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/forms_management_crud/controllers:TipoEstadoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/forms_management_crud/controllers:TipoEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/forms_management_crud/controllers:TipoEstadoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
