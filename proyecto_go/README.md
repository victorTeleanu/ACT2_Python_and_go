# Proyecto GO: Conversor de Divisas

Este proyecto implementa un **conversor de monedas en Go** que permite convertir cantidades entre diferentes monedas usando tasas fijas. Incluye pruebas unitarias para verificar que las conversiones funcionen correctamente.

## Archivos principales
- `main.go`: ejecuta un ejemplo del programa usando la función `Convert`.
- `converter/converter.go`: contiene la función `Convert` y el mapa de tasas de cambio.
- `converter/converter_test.go`: contiene los tests unitarios para la función `Convert`.
- `go.mod`: define el módulo Go del proyecto y gestiona las dependencias.

## Requisitos
- Tener instalado **Go** (1.18 o superior recomendado).
- Tener instalado VSCode con la extensión (aunque no sea del todo necesaria) **Go** (Go Team at Google).

## Cómo ejecutar el programa
Desde la raíz del proyecto:

```bash
go run main.go
```

## Cómo ejecutar los tests
Desde la raíz del proyecto:
```bash
go test ./...
```

## Descripción de los tests
| Test | Descripción | Resultado |
|------|-------------|-----------|
| `TestConvertEURtoUSD` | Verifica la conversión de 100 EUR a USD | ✅ Pasa |
| `TestConvertUSDtoEUR` | Verifica la conversión de 110 USD a EUR | ✅ Pasa |
| `TestFailIntentional` | Diseñado para fallar (espera valor incorrecto en conversión EUR a GBP) | ❌ Falla |

## Depuración
1. Abre el proyecto en VSCode.
2. Instala la extensión **Go**.
3. Abre el archivo `main.go`.
4. Coloca breakpoints en líneas clave (por ejemplo, dentro de `Convert`).
5. Pulsa `F5` y selecciona la configuración “Go: Launch File”.
6. Observa cómo se detiene la ejecución y puedes inspeccionar variables.