# Como planee el juego

## De que trata
Es un juego de ahorcado en consola, lo hice en Go para el examen de
Estructura de Datos.

## Como lo fui pensando

Primero anote las reglas basicas del juego para tener claro que
necesitaba programar:
- una palabra secreta
- el jugador adivina letra por letra
- maximo 6 errores
- el muñeco se va dibujando con cada error

## Que datos iba a necesitar
- Un slice con las palabras posibles
- Otro slice para guardar las letras que ya use
- Un numero para contar los errores

## Que funciones hice
Separe el codigo en funciones pequeñas para no tener todo en el main:
- una para checar si una letra ya se uso
- una para mostrar la palabra con guiones
- una para saber si ya gane

## El muñeco
Hice el arte ASCII a mano, una etapa por cada error y las guarde
en un slice.

## Como probe el juego
- Intente meter la misma letra dos veces
- Deje que me ahorcaran para ver si mostraba bien el mensaje
- Adivine una palabra completa para probar el mensaje de victoria
- Verifique que la pista correspondiera a la palabra correcta

## Que aprendi
Aprendi a usar slices para guardar cosas, a recorrer strings con range,
a usar el indice para relacionar dos slices y a dividir el codigo
en funciones en lugar de poner todo en main.
