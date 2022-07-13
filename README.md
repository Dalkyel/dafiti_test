# Examen Técnico Dafiti Alejandro Castañeda

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

## Ejercicio Técnico

Para ejecutar el validador podemos hacerlo de dos formas diferentes:

1. ejecutar el comando ```go run main.go``` en la raíz del proyecto. Por defecto tiene una mano predefinida la cual validará.
2. se pueden correr los tests con el comando ```go test -v -cover``` y se ejecutarán los tests para las manos predefinidas.


## MySQL

A continuación se resuelve el apartado de preguntas ***SQL***.

* ¿Cuál es el jugador más viejo de cada equipo?

```sql
select
    e.nombre as equipo,
    (
        select
        j.nombre
        from jugadores j
        where j.fk_equipos = e.id_equipos
        order by j.fecha_nacimiento desc
        limit 1
    ) as "jugador más viejo"
from equipos e;
```

* ¿Cuántos partidos jugó de visitante cada equipo? (nota: hay equipos no jugaron ningún partido)

```sql
select
    e.id_equipos,
    e.nombre as equipo,
    count(p.id_partidos) as "partidos jugados cómo visitante"
from equipos e
left join partidos p on e.id_equipos = p.fk_equipo_visitante
group by e.id_equipos, e.nombre;
```

* ¿Qué equipos jugaron el 01/01/2016 y el 12/02/2016?

```sql
select
    e.*
from equipos e
where e.id_equipos in
    (
        select
            p.fk_equipo_local
        from partidos p
        where DATE(p.fecha_partido) = STR_TO_DATE("01/01/2016", '%d/%m/%Y') or DATE(p.fecha_partido) = STR_TO_DATE("12/02/2016", '%d/%m/%Y')
        group by p.fk_equipo_local
    )
and e.id_equipos in
    (
        select
        p.fk_equipo_visitante
        from partidos p
        where DATE(p.fecha_partido) = STR_TO_DATE("01/01/2016", '%d/%m/%Y') or DATE(p.fecha_partido) = STR_TO_DATE("12/02/2016", '%d/%m/%Y')
        group by p.fk_equipo_visitante
);
```

* Diga el total de goles que hizo el equipo “Chacarita” en su historia (como local o visitante)

```sql
select
    e.id_equipos,
    e.nombre,
    (
        select
        sum(p.goles_local)
        from partidos p
        where p.fk_equipo_local = e.id_equipos
        group by p.fk_equipo_local
    )
    +
    (
        select
        sum(p.goles_visitante)
        from partidos p
        where p.fk_equipo_visitante = e.id_equipos
        group by p.fk_equipo_visitante
    ) as " total goles"
from equipos e
where nombre = "Chacarita";
```

## Extra

Cuéntanos en pocas líneas cuál crees que es la mayor innovación en el mundo del desarrollo en los últimos 5 años, y por qué.

* Para mi, el paradigma de la computación sin servidor (Serverless Computing) como enfoque arquitectónico se ha vuelto muy popular en los últimos años. Hoy es uno de los modelos de servicios en la nube de más rápido crecimiento para el desarrollo de software, modelos que continúan dominando la industria. Esto se debe a que permite el desarrollo rápido de aplicaciones sin el dolor de cabeza de la administración del servidor, y las operaciones se pueden escalar fácilmente.
