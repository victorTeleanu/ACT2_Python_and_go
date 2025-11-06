# Proyecto Python: Gestor de Contraseñas

Este proyecto implementa un **gestor de contraseñas en Python** que permite almacenar, recuperar, listar y eliminar contraseñas en memoria. Incluye pruebas unitarias para verificar el correcto funcionamiento de las funciones.

## Archivos principales
- `main.py`: ejecuta ejemplos del programa usando las funciones del gestor de contraseñas.
- `passwdmgr/passwdmgr.py`: contiene la lógica del gestor de contraseñas. Incluye:

    - Clase Entry: representa una entrada de contraseña con atributos service, username y password. El método to_dict() permite convertir la entrada en un diccionario.

    - Diccionario passwords: almacena las entradas de manera persistente en memoria, usando el nombre del servicio como clave.

    - Función add(service, username, password): añade una nueva entrada o actualiza una existente en el diccionario passwords.

    - Función get(service): devuelve las credenciales de un servicio dado. Si no existe, retorna "Servicio no encontrado".

    - Función delete(service): elimina la entrada de un servicio. Si no existe, imprime un mensaje de error.

    - Función list_services(): devuelve una lista con todos los nombres de servicios almacenados en el diccionario.

- `passwdmgr/passwdmgr_test.py`: contiene los tests unitarios para las funciones.
- `passwdmgr/__init__.py`: indica que la carpeta es un paquete Python.

## Requisitos
- Tener instalado Python 3.8 o superior.
- Tener instalado VSCode con la extensión **Python** (Microsoft).

## Cómo ejecutar el programa
```bash
python main.py
```

## Cómo ejecutar los tests
Desde la raíz del proyecto:
```bash
python -m unittest passwdmgr.test_passwdmgr
```

## Descripción de los tests
| Test | Descripción | Resultado |
|------|-------------|-----------|
| `test_basic_operations` | Verifica agregar, obtener, listar y eliminar servicios correctamente | ✅ Pasa |
| `test_add_and_get` | Verifica que se puedan agregar credenciales y recuperarlas correctamente | ✅ Pasa |
| `test_fail_intentional` | Diseñado para fallar (espera valor incorrecto en el usuario) | ❌ Falla |

## Depuración
1. Abre el proyecto en VSCode.
2. Instala la extensión **Python**.
3. Abre el archivo `main.py`.
4. Coloca breakpoints en líneas clave (por ejemplo, dentro de `divide`).
5. Pulsa `F5` y selecciona la configuración “Python: Ejecutar main.py”.
6. Observa cómo se detiene la ejecución y puedes inspeccionar variables.
