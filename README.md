# refactor_controller
Este proyecto es un script actualizado a python 3.9 para realizar ajustes en la estructura de respuesta JSON y `status code` de los micro servicios desarrollados en el framework `Beego` de forma masiva y automática. En este caso se incluye el formateo de los campos fecha_creacion y fecha_modificacion para que no pierdan consistencia usando la libreria time_bogota de [utils_oas](https://github.com/udistrital/utils_oas).

Refactoriza los métodos POST, GETONE, GETALL, PUT, DELETE, en los controladores de una API; especificando el [código de estado HTTP](https://es.wikipedia.org/wiki/Anexo:C%C3%B3digos_de_estado_HTTP) de cada solicitud y retornando siempre un JSON. Esto conforme a los estandares de la Oficina Asesora de Sistemas.

La única restricción que existe, es que **solo realiza los ajustes en micro servicios que no han sido personalizado o modificado a cómo los genera el framework**.

## Procedimiento

### 1 Configurar paquete utils_oas

Se debe implementar las plantilla de error que se encuentra en [utils_oas](https://github.com/udistrital/utils_oas) como se indica a continuación.

##### 1.1. Importar paquete:
Para esto Editar el `main.go` de la API a Ajustar.
```golang
import (
  "github.com/udistrital/utils_oas/customerrorv2"
)
```
##### 1.2 Implementación en `func main()`
```golang
beego.ErrorController(&customerrorv2.CustomErrorController{})
```

##### El **main.go** Lucirá de la siguiente forma:
```golang
package main

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "github.com/astaxie/beego/plugins/cors"
    _ "github.com/jotavargas/debug_beego_request/routers"
    _ "github.com/lib/pq"
    "github.com/udistrital/utils_oas/customerrorv2"
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
    beego.ErrorController(&customerrorv2.CustomErrorController{})
    beego.Run()
}
```

### 2 Ejecución de Script para Refactorizar los Controladores

##### 2.1  Clonar repositorio
```golang
git clone git@github.com:udistrital/refactor_controller.git
```
##### 2.2 Ejecución del script
```bash
#ir al proyecto
cd refactor_controller

# establecer la version 2
git checkout version/0.0.2

#como ejecutar (con python 2.7)
python2.7 main.py -F ruta_controladores_del_api_a_refactoring

#Ejemplo:
python2.7 main.py -F /home/jjvargass/go/src/github.com/udistrital/api_financiera/controllers
```
##### 2.3 Indentar e importar package en Controladores *.go
```bash
cd ruta_controladores_del_api_a_refactoring
gofmt -w *.go
goimports -w *.go
```

##### 2.4 Ajustar el tipo de dato en fecha_creacion y fecha_modificacion

Para finalizar lo que corresponde a los campos de fecha, es necesario que en los modelos el tipo de dato de estos campos sea string para su correcta inserción y modificación en la base de datos.

```golang

type ModeloEjemplo struct {
	FechaCreacion     string `orm:"column(fecha_creacion);type(timestamp without time zone);null"`
	FechaModificacion string `orm:"column(fecha_modificacion);type(timestamp without time zone);null"`
}
```

### 3 Conparación de cambios (Antes y Despues)
Los cambios específicos en cada uno de los microservicios se definirán a continuación.

#### 3.1 Solicitud POST
<table>
 <tr>
  <td colspan="2"><img src="/images/post.png">
</td>
 </tr>
 <tr>
  <td>A la izquierda el método por defecto creados por el Framewrok. </td>
  <td>A la derecha el refactor por el Script <a href="https://github.com/udistrital/refactor_controller">refactor_controller</a> </td>
 </tr>
</table>

#### 3.2 Solicitud GETONE
<table>
 <tr>
  <td colspan="2"><img src="/images/getOne.png">
</td>
 </tr>
 <tr>
  <td>A la izquierda el método por defecto creados por el Framewrok. </td>
  <td>A la derecha el refactor por el Script <a href="https://github.com/udistrital/refactor_controller">refactor_controller</a> </td>
 </tr>
</table>


#### 3.3 Solicitud GETALL
<table>
 <tr>
  <td colspan="2">
    <img src="/images/getAll-1.png"><br><br>
    <img src="/images/getAll-2.png">
  </td>
 </tr>
 <tr>
  <td>A la izquierda el método por defecto creados por el Framewrok. </td>
  <td>A la derecha el refactor por el Script <a href="https://github.com/udistrital/refactor_controller">refactor_controller</a> </td>
 </tr>
</table>


#### 3.4 Solicitud PUT
<table>
 <tr>
  <td colspan="2"><img src="/images/put.png">
</td>
 </tr>
 <tr>
  <td>A la izquierda el método por defecto creados por el Framewrok. </td>
  <td>A la derecha el refactor por el Script <a href="https://github.com/udistrital/refactor_controller">refactor_controller</a> </td>
 </tr>
</table>

#### 3.5 Solicitud DELETE
<table>
 <tr>
  <td colspan="2"><img src="/images/delete.png">
</td>
 </tr>
 <tr>
  <td>A la izquierda el método por defecto creados por el Framewrok. </td>
  <td>A la derecha el refactor por el Script <a href="https://github.com/udistrital/refactor_controller">refactor_controller</a> </td>
 </tr>
</table>


### 4 Estructura Estandar de Respuestas (JSON)
A continuación se detalla la estructura JSON que responderá el api, al aplicar el refactor.

<table>
 <tr>
  <td><img src="/images/jsonUnicoRegistros.png">
  <td><img src="/images/jsonMultiplesRegistros.png">
</td>
 </tr>
 <tr>
  <td>Respuesta con único registro </td>
  <td>Resultado con múltiples registros</td>
 </tr>
</table>


#### POST
![Refactor Metodo Post](/images/post_test.png)

##### Post Correcto
<table>
 <tr>
  <td><img src="/images/post-correcto-1.png">
  <td><img src="/images/post-correcto-2.png">
</td>
 </tr>
 <tr>
  <td>Cuerpo de la respuesta </td>
  <td>Resultado de la solicitud</td>
 </tr>
</table>

<br>

##### Post con Parametros Incorrecto
<table>
 <tr>
  <td><img src="/images/post-incorrecto-1.png">
  <td><img src="/images/post-incorrecto-2.png">
</td>
 </tr>
 <tr>
  <td>Cuerpo de la respuesta </td>
  <td>Resultado de la solicitud</td>
 </tr>
</table>



### 5 Pruebas con JMeter
Se realizaron diferentes escenarios alternos por cada uno de los servicios expuesto por el API generado en beego.  Adjunto se encontrará el código fuente.

![Test Jmeter](/images/test_jmeter.png)
#### [Link Testing JMeter](/generacion_de_apis/src/beegoTodasLasSolicitudes.jmx)


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
