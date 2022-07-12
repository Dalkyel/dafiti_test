# Examen Técnico Dafiti

A continuación se resuelven los puntos del examen técnico de Dafiti para el puesto de Backend Developer.

## Preguntas

1. Supone que en un repositorio GIT hiciste un commit y te olvidaste un archivo. Explicar como se soluciona si hiciste push, y como si aun no hiciste. De ser posible, buscar que quede solo un commit con los cambios

Existen diferentes formas de abordar este escenario y cada uno con sus pros y sus contras.

* Si no se ha realizado el push:

  * Se podría eliminar el commit con ```git revert``` para crear un nuevo commit y asignar los archivos faltantes. Esta opción se podría utilizar si se quiere cambiar los datos del commit como la fecha o los comentarios etc.

  * Otra opción es utilizar el ```git commit -a --amend``` para modificar el ultimo commit y solo agregar el archivo faltante.

* Si ya se realizó el push:

  * La opción del ```git commit -a --amend``` funcionaría para el escenario en el que ya se hizo un push al repositorio, pero se debe tener en cuenta si la rama del repositorio está siendo usada por otros desarrolladores porque esto podría traer confusiones en el proceso de desarrollo. En mi caso, crearía otro commit nuevo para evitar errores de conflicto.

2. Tenes un sitio en tu computadora de desarrollo, y cuando entras desde el navegador, en la consola te aparece esto: ***<https://site.local/favicon.ico> Failed to load resource: net::ERR_INSECURE_RESPONSE*** El archivo existe, el sitio debe cargar por https. Como lo solucionas?

    * Lo primero a revisar sería el correcto funcionamiento de los certificados **SSL** para el sitio al que se está tratando de acceder, podría haber alguna falla con estos o una mala configuración. Dependiendo del Motor que se este usando, la configuración podría ser diferente.

3. Tenes un archivo comiteado en un repositorio GIT que deseas que quede ignorado. Que pasos debes realizar?

   * Primero debemos asegurarnos de eliminar el archivo del repositorio remoto si este ya se encuentra allí y también eliminarlo localmente. Una forma de asegurar esto es con el comando ```git rm --cached filename```.
   * Una vez eliminemos el archivo debemos incluirlo en el archivo ***.gitignore*** del proyecto si queremos que no se vuelva a agregar al repositorio.

4. Explica las ventajas de cargar en tu sitio las librerías de terceros por GTM.

   * En esta pregunta voy a ser sincero, nunca he usado GTM para gestión de librerías de terceros. Intenté consultarlo en Google para ponerme al día con el funcionamiento pero no encontré nada referente a GTM.
