# forms_management_crud
API de gestión de formularios dinámicos

Integración con

 - `CI`
 - `AWS Lambda - S3`
 - `Drone 1.x`
 - `forms_management_crud master/develop`

## Requerimientos
Go version >= 1.8.

## Preparación
Para usar el API, usar el comando:

 - `go get github.com/planesticud/forms_management_crud`

## Ejecución
Definir los valores de las siguientes variables de entorno:

 - `API_FORMS_MANAGEMENT_HTTP_PORT`: Puerto asignado para la ejecución del API
 - `FORMS_MANAGEMENT_CRUD__PGUSER`: Usuario de la base de datos
 - `FORMS_MANAGEMENT_CRUD__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `FORMS_MANAGEMENT_CRUD__PGURLS`: Host de conexión
 - `FORMS_MANAGEMENT_CRUD__PGDB`: Nombre de la base de datos
 - `FORMS_MANAGEMENT_CRUD__SCHEMA`: Esquema a utilizar en la base de datos

## Ejemplo
API_FORMS_MANAGEMENT_HTTP_PORT=9011 FORMS_MANAGEMENT_CRUD__PGUSER=user FORMS_MANAGEMENT_CRUD__PGPASS=password FORMS_MANAGEMENT_CRUD__PGURLS=localhost FORMS_MANAGEMENT_CRUD__PGDB=bd FORMS_MANAGEMENT_CRUD__SCHEMA=schema_new bee run

## Modelo BD
![image](https://github.com/planesticud/forms_management_crud/blob/develop/modelo_forms_management_crud.png).