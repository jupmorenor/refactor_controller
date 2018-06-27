# refactor_controller

Este proyecto permite refactorizar los métodos POST, GETONE, GETALL, PUT, DELETE, de los controladores de una API generada por el Framework Beego.

Especifica el código de estatus de la respuesta a cada uno de los métodos y siempre retorna un JSON.

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

Clonamos este repositorio

    git clone https://github.com/jotavargas/refactor_controller.git

Ingresamos al proyecto para ejecutar main.py

    cd refactor_controller

Ejecutamos Script

    # Ejecutar con python 2.7
    python main.py -F ruta_controladores_del_api_a_refactoring


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
