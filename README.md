# forms_management_crud


--Api de personas con CI--
CI deploy with lambda - S3
Drone 0.8 
forms_management_crud master/develop

## Requirements
Go version >= 1.8.

## Preparación:
    Para usar el API, usar el comando:
        - go get github.com/udistrital/forms_management_crud

## Run

Definir los valores de las siguientes variables de entorno:

 - `API_FORMS_MANAGEMENT_HTTP_PORT`: Puerto asignado para la ejecución del API
 - `FORMS_MANAGEMENT_CRUD__PGUSER`: Usuario de la base de datos
 - `FORMS_MANAGEMENT_CRUD__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `FORMS_MANAGEMENT_CRUD__PGURLS`: Host de conexión
 - `FORMS_MANAGEMENT_CRUD__PGDB`: Nombre de la base de datos
 - `FORMS_MANAGEMENT_CRUD__SCHEMA`: Esquema a utilizar en la base de datos

Ejemplo: API_FORMS_MANAGEMENT_HTTP_PORT=8083 FORMS_MANAGEMENT_CRUD__PGUSER=user FORMS_MANAGEMENT_CRUD__PGPASS=password FORMS_MANAGEMENT_CRUD__PGURLS=localhost FORMS_MANAGEMENT_CRUD__PGDB=udistrital_core_db FORMS_MANAGEMENT_CRUD__SCHEMA=core_new bee run

## MODELO
![image](https://user-images.githubusercontent.com/14035745/42359402-4cd653f4-80a7-11e8-8b90-61e30a20bbaf.png).
