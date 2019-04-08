# refactor_controller

Este proyecto es un script desarrollado en python 2.7 para realizar los ajustes de los micro servicios desarrollados en el framework Beego de forma masiva y automática.

Refactoriza los métodos POST, GETONE, GETALL, PUT, DELETE, en los controladores de una API; especificando el [código de estado HTTP](https://es.wikipedia.org/wiki/Anexo:C%C3%B3digos_de_estado_HTTP) de cada solicitud y retornando siempre un JSON. Esto conforme a los estandares de la Oficina Asesora de Sistemas.

la única restricción que existe, es que **solo realiza los ajustes en micro servicios que nos se han personalizado o modificado en sus líneas**.


A la izquierda el método por defecto en el Framewrok a la derecha el refactor


### POST
![Refactor Metodo Post](/images/post.png)

### GETONE
![Refactor Metodo GetOne](/images/getone.png)

### GETALL
![Refactor Metodo GetAll](/images/getall.png)

### PUT
![Refactor Metodo Post](/images/put.png)

### DELETE
![Refactor Metodo Post](/images/delete.png)

## Ejecutar Script

- Clonar repositorio

      git clone https://github.com/jotavargas/refactor_controller.git

- Ingresar al proyecto para ejecutar main.py

      cd refactor_controller

- Ejecutar Script

      # Ejecutar con python 2.7
      python main.py -F ruta_controladores_del_api_a_refactoring

- Indentar e importar package en Controladores *.go

      cd ruta_controladores_del_api_a_refactoring
      gofmt -w *.go
      goimports -w *.go


## Refactoring API Controllers

Editar el **main.go** de la API a Ajustar. Agregar las plantillas de errores que se encuentran en el repositorio **[utils_oas](https://github.com/udistrital/utils_oas)** de la siguiente forma.

- Importar paquete:

      import (
        "github.com/udistrital/utils_oas/customerror"
      )

- Implementación en **func main()**:

      beego.ErrorController(&customerror.CustomErrorController{})

- El **main.go** Lucirá de la siguiente forma:

      package main

      import (
          "github.com/astaxie/beego"
          "github.com/astaxie/beego/orm"
          "github.com/astaxie/beego/plugins/cors"
          _ "github.com/jotavargas/debug_beego_request/routers"
          _ "github.com/lib/pq"
          "github.com/udistrital/utils_oas/customerror"
      )

      func init() {
          orm.RegisterDataBase("default", "postgres", "postgres://postgres:postgres@127.0.0.1/test?sslmode=disable")
      }

      func main() {
          if beego.BConfig.RunMode == "dev" {
              beego.BConfig.WebConfig.DirectoryIndex = true
              beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
          }

          beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
              AllowOrigins: []string{"*"},
              AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
              AllowHeaders: []string{"Origin", "x-requested-with",
                  "content-type",
                  "accept",
                  "origin",
                  "authorization",
                  "x-csrftoken"},
              ExposeHeaders:    []string{"Content-Length"},
              AllowCredentials: true,
          }))
          beego.ErrorController(&customerror.CustomErrorController{})
          beego.Run()
      }



## Estructura JSON

Obtenemos la siguiente estructura cuando el framework a controlado un error de bd

  ![Json1](/images/json01.png)

Obtenemos este Json cuando es desarrollador ha personalizado el servicio y estructura del error en el atributo development

  ![Json2](/images/json02.png)

Obtenemos este Json cuando ingresamos a una servicio  que no existe

  ![Json3](/images/json03.png)

## Licencia

This file is part of refactor_controller.

refactor_controller is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

Foobar is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with Foobar.  If not, see <https://www.gnu.org/licenses/>.
