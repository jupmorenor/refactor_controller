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
