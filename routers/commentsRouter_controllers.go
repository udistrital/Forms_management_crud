package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/planesticud/forms_management_crud/controllers:CampoProgramaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/forms_management_crud/controllers:CampoProgramaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/forms_management_crud/controllers:CampoProgramaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/forms_management_crud/controllers:CampoProgramaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/forms_management_crud/controllers:CampoProgramaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/forms_management_crud/controllers:CampoProgramaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/forms_management_crud/controllers:CampoProgramaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/forms_management_crud/controllers:CampoProgramaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/forms_management_crud/controllers:CampoProgramaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/forms_management_crud/controllers:CampoProgramaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/forms_management_crud/controllers:FormularioProgramaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/forms_management_crud/controllers:FormularioProgramaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/forms_management_crud/controllers:FormularioProgramaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/forms_management_crud/controllers:FormularioProgramaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/forms_management_crud/controllers:FormularioProgramaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/forms_management_crud/controllers:FormularioProgramaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/forms_management_crud/controllers:FormularioProgramaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/forms_management_crud/controllers:FormularioProgramaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/forms_management_crud/controllers:FormularioProgramaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/forms_management_crud/controllers:FormularioProgramaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/forms_management_crud/controllers:TipoEstadoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/forms_management_crud/controllers:TipoEstadoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/forms_management_crud/controllers:TipoEstadoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/forms_management_crud/controllers:TipoEstadoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/forms_management_crud/controllers:TipoEstadoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/forms_management_crud/controllers:TipoEstadoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/forms_management_crud/controllers:TipoEstadoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/forms_management_crud/controllers:TipoEstadoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/planesticud/forms_management_crud/controllers:TipoEstadoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/forms_management_crud/controllers:TipoEstadoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
