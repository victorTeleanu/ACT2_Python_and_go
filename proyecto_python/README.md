# Proyecto Python: Gestor de Contraseñas

Este proyecto implementa un **gestor de contraseñas en Python** que permite almacenar, recuperar, listar y eliminar contraseñas en memoria. Incluye pruebas unitarias para verificar el correcto funcionamiento de las funciones.

## Archivos principales
- `main.py`: ejecuta ejemplos del programa usando las funciones del gestor de contraseñas.
- `passwdmgr/passwdmgr.py`: contiene las funciones (`add`, `get`, `delete`, `list_services`) y el diccionario en memoria.
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
python -m unittest passwdmgr.passwdmgr_test
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
