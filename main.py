#!/usr/bin/python
# -*- coding: utf-8 -*-

import logging
from optparse import OptionParser
from os import listdir

logging.basicConfig()
_logger = logging.getLogger('Oas')


def reemplazo(linea_actual, flags_linea, numero_linea_tope, archivo_linea, reemplazo):
    '''
    funcion reemplazo
        @linea_actual linea de iteracion del documento
        @flags_linea linea donde se encontro la validación del metodo
        @numero_linea_tope numero de lineas tope para hacer cambios en el metodo
        @archivo_linea strin del la line actual
        @reemplazo diccionario donde se encuentran los cambios
    '''
    if (linea_actual) > (flags_linea + numero_linea_tope):
        return archivo_linea, False, 0
    for i, co in reemplazo[reemplazo.keys()[0]].items():
        #print i, ":", co
        if i in archivo_linea:
            archivo_linea = co
            _logger.info(" Linea {0}: ## {1} ##: {2} -> Refactor".format(linea_actual, reemplazo.keys()[0],i,co ))
    return archivo_linea, True, flags_linea

def escribir_archivo_reemplazo(nombre_archivo, lineas):
    '''
    funcion escribir_archivo_reemplazo
        @nombre_archivo nombre de nuevo archivo
        @lineas contenido del nuevo archivo
    '''
    with open(nombre_archivo,'w') as g:
        for linea in lineas:
            g.write(linea)

def main():
    usage = "Refactoring of Controllers in API Beego"
    parser = OptionParser(usage)
    # Input Parameters
    parser.add_option("-F", "--file_route", dest="file_route", help="Ruta del Archivo")
    parser.add_option("-d", "--debug", dest="debug", help="Mostrar mensajes de debug utilize 10", default=10)

    (options, args) = parser.parse_args()
    _logger.setLevel(int(options.debug))

    if not options.file_route:
        parser.error('Espesifique ruta de los controladores')

    # Especificacion de los Cambios por Metodos
    reemplazos_post = {
        "@Title Post": {
            "@Failure 403 body is empty": "// @Failure 400 the request contains incorrect syntax\n",
            "c.Data[\"json\"] = err.Error()": "beego.Error(err)\n //c.Data[\"development\"] = map[string]interface{}{\"Code\": \"000\", \"Body\": err.Error(), \"Type\": \"error\"}\nc.Data[\"system\"] = err\n c.Abort(\"400\")\n"
        },
    }

    reemplazos_get_one = {
        "@Title Get One": {
            "@Failure 403 :id is empty": "// @Failure 404 not found resource\n",
            "c.Data[\"json\"] = err.Error()": "beego.Error(err)\n //c.Data[\"development\"] = map[string]interface{}{\"Code\": \"000\", \"Body\": err.Error(), \"Type\": \"error\"}\nc.Data[\"system\"] = err\n c.Abort(\"404\")\n"
        },
    }

    reemplazos_get_all = {
        "@Title Get All": {
            "@Failure 403": "// @Failure 404 not found resource\n",
            "c.Data[\"json\"] = err.Error()": "beego.Error(err)\n //c.Data[\"development\"] = map[string]interface{}{\"Code\": \"000\", \"Body\": err.Error(), \"Type\": \"error\"}\nc.Data[\"system\"] = err\n c.Abort(\"404\")\n",
            "c.Data[\"json\"] = l": "if l == nil {\n l = append(l, map[string]interface{}{})\n }\n c.Data[\"json\"] = l\n",
        },
    }

    reemplazos_put = {
        "@Title Put": {
            "@Failure 403 :id is not int": "// @Failure 400 the request contains incorrect syntax\n",
            "c.Data[\"json\"] = err.Error()": "beego.Error(err)\n //c.Data[\"development\"] = map[string]interface{}{\"Code\": \"000\", \"Body\": err.Error(), \"Type\": \"error\"}\nc.Data[\"system\"] = err\n c.Abort(\"400\")\n",
            "c.Data[\"json\"] = \"OK\"": "c.Data[\"json\"] = v\n",
        },
    }

    reemplazos_delete = {
        "@Title Delete": {
            "@Failure 403 id is empty": "// @Failure 404 not found resource\n",
            "c.Data[\"json\"] = \"OK\"": "c.Data[\"json\"] = map[string]interface{}{\"Id\": id}\n",
            "c.Data[\"json\"] = err.Error()": "beego.Error(err)\n //c.Data[\"development\"] = map[string]interface{}{\"Code\": \"000\", \"Body\": err.Error(), \"Type\": \"error\"}\nc.Data[\"system\"] = err\n c.Abort(\"404\")\n"
        },
    }

    # Recorrer Archivos
    for archivo in listdir(options.file_route):
        ruta_archivo =  options.file_route + "/" + archivo
        _logger.info("#####################################################")
        _logger.info("######## {0}".format(ruta_archivo))
        _logger.info("#####################################################")

        with open(ruta_archivo) as f:
            lineas = f.readlines()
            #Flags
            flags_post = False
            flags_get_one = False
            flags_get_all = False
            flags_put = False
            flags_delete = False
            flags_linea_metodo = 0
            liena_actual = 1
            new_fiel = []
            #Recorrer Archivo por salto de lineas
            for linea in lineas:
                #Validacion Post
                if reemplazos_post.keys()[0] in linea:
                    flags_post = True
                    flags_linea_metodo = liena_actual
                if flags_post:
                    linea, flags_post, flags_linea_metodo = reemplazo(liena_actual,flags_linea_metodo, 20, linea, reemplazos_post)

                #Validacion GetOne
                if reemplazos_get_one.keys()[0] in linea:
                    flags_get_one = True
                    flags_linea_metodo = liena_actual
                if flags_get_one:
                    linea, flags_get_one, flags_linea_metodo = reemplazo(liena_actual,flags_linea_metodo, 16, linea, reemplazos_get_one)

                #Validacion GetAll
                if reemplazos_get_all.keys()[0] in linea:
                    flags_get_all = True
                    flags_linea_metodo = liena_actual
                if flags_get_all:
                    linea, flags_get_all, flags_linea_metodo = reemplazo(liena_actual,flags_linea_metodo, 60, linea, reemplazos_get_all)

                #Validacion Put
                if reemplazos_put.keys()[0] in linea:
                    flags_put = True
                    flags_linea_metodo = liena_actual
                if flags_put:
                    linea, flags_put, flags_linea_metodo = reemplazo(liena_actual,flags_linea_metodo, 21, linea, reemplazos_put)

                #Validacion Delete
                if reemplazos_delete.keys()[0] in linea:
                    flags_delete = True
                    flags_linea_metodo = liena_actual
                if flags_delete:
                    linea, flags_delete, flags_linea_metodo = reemplazo(liena_actual,flags_linea_metodo, 15, linea, reemplazos_delete)
                # Construir Nuevo Archivo
                #print linea
                new_fiel.append(linea)
                liena_actual += 1
        f.close()
        #Escritura de las Cambios
        escribir_archivo_reemplazo(ruta_archivo, new_fiel)

if __name__ == '__main__':
    main()