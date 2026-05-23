# Testing y QA

## 1. Objetivo del documento

Este documento define la estrategia de testing y control de calidad para la red social desarrollada con Go, htmx y templ.

El objetivo es garantizar que la aplicacion cumple los requisitos definidos en el PRD, que las funcionalidades principales funcionan correctamente y que el codigo generado con apoyo de IA se revisa y valida antes de considerarse terminado.

## 2. Enfoque general

La estrategia de calidad combina:

- Tests automaticos.
- Pruebas manuales.
- Revision de codigo.
- Revision con IA.
- Validacion de criterios de aceptacion.
- Comprobacion de seguridad basica.
- Pruebas de interfaz y usabilidad.

No todas las partes del proyecto requieren el mismo nivel de prueba. Las funcionalidades criticas, como autenticacion, permisos, publicaciones, likes, comentarios y follows, tendran mayor prioridad.

## 3. Objetivos de calidad

- Comprobar que la aplicacion cumple los requisitos funcionales.
- Evitar errores en autenticacion y autorizacion.
- Validar reglas de negocio importantes.
- Detectar errores generados por IA.
- Mantener el codigo entendible y mantenible.
- Garantizar que la aplicacion puede ejecutarse localmente.
- Preparar el proyecto para integracion continua.
- Documentar las pruebas realizadas para la memoria final.

## 4. Tipos de pruebas

### Tests unitarios

Validan funciones o servicios de forma aislada.

Ejemplos:

- Validacion de email.
- Validacion de contrasena.
- Hash y verificacion de contrasena.
- Reglas de permisos.
- Logica para dar o quitar like.
- Logica para seguir o dejar de seguir usuarios.

### Tests de servicios

Validan casos de uso de negocio.

Ejemplos:

- Registrar usuario.
- Crear publicacion.
- Editar publicacion propia.
- Impedir editar publicacion ajena.
- Crear comentario.
- Evitar likes duplicados.
- Evitar follows duplicados.

### Tests de handlers HTTP

Validan la capa HTTP usando peticiones simuladas.

Ejemplos:

- `GET /login` devuelve la pagina de login.
- `POST /login` autentica correctamente.
- `POST /posts` crea una publicacion si el usuario esta autenticado.
- `POST /posts` rechaza peticiones sin sesion.
- Rutas privadas redirigen si no hay usuario autenticado.

### Tests de repositorios

Validan el acceso a base de datos.

Ejemplos:

- Crear usuario.
- Buscar usuario por email.
- Crear publicacion.
- Listar feed.
- Insertar like.
- Eliminar follow.

Estos tests pueden ejecutarse contra una base de datos de test o usando SQLite si la arquitectura lo permite.

### Pruebas de integracion

Validan que varias capas funcionan juntas.

Ejemplos:

- Registro completo: handler, servicio, repositorio y base de datos.
- Crear publicacion y verla en el feed.
- Dar like y comprobar contador.
- Seguir usuario y comprobar que aparecen sus publicaciones.

### Pruebas manuales

Validan flujos completos desde la interfaz.

Ejemplos:

- Registro.
- Login.
- Logout.
- Crear publicacion.
- Comentar.
- Dar like.
- Buscar usuario.
- Editar perfil.

### Pruebas de UI/UX

Validan que la interfaz es clara y usable.

Ejemplos:

- Formularios comprensibles.
- Mensajes de error visibles.
- Botones principales accesibles.
- Feed legible.
- Interacciones htmx sin recargas innecesarias.
- Diseno responsive.

### Pruebas de seguridad basica

Validan aspectos minimos de seguridad.

Ejemplos:

- Contrasenas almacenadas con hash.
- Rutas privadas protegidas.
- Un usuario no puede editar publicaciones ajenas.
- Un usuario no puede borrar comentarios ajenos.
- Un usuario no puede seguirse a si mismo.
- No se permiten likes duplicados.
- Las sesiones se invalidan al hacer logout.

## 5. Herramientas

### Go testing

Se utilizara el paquete estandar de Go:

```text
go test ./...
```

### httptest

Se utilizara para probar handlers HTTP sin levantar un servidor real.

### go vet

Se utilizara para detectar problemas comunes:

```text
go vet ./...
```

### Integracion continua

Se configurara un workflow para ejecutar pruebas automaticamente.

Comprobaciones recomendadas:

```text
go test ./...
go vet ./...
go build ./cmd/web
```

### IA para QA

Antigravity y opencode se utilizaran para:

- Generar casos de prueba.
- Revisar requisitos.
- Detectar errores potenciales.
- Proponer tests.
- Revisar permisos y seguridad.

## 6. Estrategia por modulo

### Autenticacion

Riesgo:

- Alto.

Pruebas necesarias:

- Registro con datos validos.
- Registro con email duplicado.
- Registro con username duplicado.
- Login correcto.
- Login con contrasena incorrecta.
- Logout.
- Acceso a ruta privada sin sesion.
- Acceso a ruta privada con sesion.
- Verificacion de hash de contrasena.

### Perfil

Riesgo:

- Medio.

Pruebas necesarias:

- Ver perfil propio.
- Ver perfil de otro usuario.
- Editar perfil propio.
- Impedir editar perfil ajeno.
- Validar campos de perfil.

### Publicaciones

Riesgo:

- Alto.

Pruebas necesarias:

- Crear publicacion valida.
- Rechazar publicacion vacia.
- Rechazar publicacion demasiado larga.
- Editar publicacion propia.
- Impedir editar publicacion ajena.
- Eliminar publicacion propia.
- Impedir eliminar publicacion ajena.

### Feed

Riesgo:

- Medio-alto.

Pruebas necesarias:

- Mostrar publicaciones ordenadas por fecha descendente.
- Mostrar publicaciones propias.
- Mostrar publicaciones de usuarios seguidos.
- No mostrar contenido incorrecto si el feed es personalizado.
- Comprobar estado vacio cuando no hay publicaciones.

### Likes

Riesgo:

- Medio-alto.

Pruebas necesarias:

- Dar like a una publicacion.
- Quitar like.
- Evitar likes duplicados.
- Actualizar contador.
- Validar respuesta htmx.

### Comentarios

Riesgo:

- Medio.

Pruebas necesarias:

- Crear comentario.
- Rechazar comentario vacio.
- Mostrar comentario en su publicacion.
- Eliminar comentario propio.
- Impedir eliminar comentario ajeno.
- Validar actualizacion parcial con htmx.

### Seguimiento

Riesgo:

- Medio-alto.

Pruebas necesarias:

- Seguir usuario.
- Dejar de seguir usuario.
- Evitar seguirse a si mismo.
- Evitar follows duplicados.
- Comprobar impacto en el feed.

### Busqueda

Riesgo:

- Medio.

Pruebas necesarias:

- Buscar usuario por coincidencia parcial.
- Mostrar resultados.
- Mostrar estado sin resultados.
- Validar respuesta htmx.

## 7. Checklist de pruebas manuales

### Autenticacion

```text
[ ] Registro con datos validos
[ ] Registro con email ya usado
[ ] Registro con username ya usado
[ ] Login correcto
[ ] Login incorrecto
[ ] Logout
[ ] Acceso a feed sin login
[ ] Acceso a feed con login
```

### Publicaciones

```text
[ ] Crear publicacion
[ ] Crear publicacion vacia
[ ] Editar publicacion propia
[ ] Intentar editar publicacion de otro usuario
[ ] Eliminar publicacion propia
[ ] Intentar eliminar publicacion de otro usuario
[ ] Ver publicacion en feed
```

### Interacciones

```text
[ ] Dar like
[ ] Quitar like
[ ] Comprobar contador de likes
[ ] Crear comentario
[ ] Crear comentario vacio
[ ] Eliminar comentario propio
[ ] Seguir usuario
[ ] Dejar de seguir usuario
[ ] Intentar seguirse a si mismo
```

### UI/UX

```text
[ ] Formularios con labels visibles
[ ] Mensajes de error comprensibles
[ ] Botones principales claros
[ ] Feed legible
[ ] Perfil legible
[ ] Interacciones htmx sin recarga completa
[ ] Vista movil aceptable
[ ] Vista escritorio aceptable
```

## 8. Criterios de aceptacion generales

Una funcionalidad se considera aceptada cuando:

- Cumple los criterios definidos en el PRD.
- El codigo compila.
- No rompe funcionalidades existentes.
- Tiene validacion en servidor si recibe datos de usuario.
- Protege permisos si afecta a contenido de usuarios.
- Ha sido revisada por el equipo.
- Ha sido probada manualmente o mediante tests.
- Se ha documentado el uso de IA si ha intervenido de forma relevante.

## 9. Revision de codigo

La revision de codigo tendra en cuenta:

- Claridad de nombres.
- Separacion de responsabilidades.
- Ausencia de duplicacion innecesaria.
- Handlers simples.
- Servicios con logica de negocio.
- Repositorios centrados en acceso a datos.
- Gestion adecuada de errores.
- Validaciones suficientes.
- No introducir dependencias innecesarias.
- Cumplimiento de la arquitectura definida.

## 10. QA con inteligencia artificial

La IA se utilizara como apoyo, no como sustituto de la revision humana.

### Uso de Antigravity

Antigravity se utilizara para:

- Revisar si una funcionalidad cumple los requisitos.
- Detectar riesgos de arquitectura.
- Proponer casos de prueba.
- Revisar la calidad general del proyecto.
- Ayudar a redactar informes de QA.

### Uso de opencode

opencode se utilizara para:

- Generar tests concretos.
- Ejecutar pruebas.
- Corregir errores.
- Refactorizar codigo.
- Revisar handlers, servicios y repositorios.

## 11. Prompts recomendados para QA

### Revision funcional

```text
Actua como agente QA del proyecto. Revisa esta funcionalidad de una red social desarrollada con Go, htmx y templ. Comprueba si cumple los criterios de aceptacion, detecta posibles errores y propone pruebas manuales y automaticas.
```

### Revision de seguridad

```text
Actua como auditor de seguridad basica. Revisa esta implementacion de autenticacion y permisos en Go. Busca problemas relacionados con sesiones, rutas privadas, hash de contrasenas, validacion de entrada y acceso no autorizado.
```

### Generacion de tests

```text
Actua como desarrollador QA. Genera tests en Go para esta capa de servicio. Los tests deben cubrir casos correctos, errores de validacion y errores de permisos.
```

### Revision htmx

```text
Actua como revisor frontend. Comprueba si esta interaccion htmx actualiza solo el fragmento necesario, si gestiona errores correctamente y si la experiencia de usuario es clara.
```

## 12. Registro de pruebas

Las pruebas importantes se documentaran en:

```text
docs/ia/revisiones/
```

Formato recomendado:

```text
Fecha:
Funcionalidad:
Tipo de prueba:
Herramienta:
Resultado:
Errores encontrados:
Correcciones aplicadas:
Estado final:
```

## 13. Trazabilidad entre requisitos y pruebas

| Requisito | Tipo de prueba | Prioridad |
| --- | --- | --- |
| Registro de usuarios | Unitario, handler, manual | Alta |
| Login/logout | Unitario, handler, manual | Alta |
| Rutas privadas | Handler, manual, seguridad | Alta |
| Perfil | Handler, manual | Media |
| Crear publicaciones | Servicio, handler, manual | Alta |
| Editar/eliminar publicaciones | Servicio, handler, seguridad | Alta |
| Feed | Servicio, integracion, manual | Media-alta |
| Likes | Servicio, handler, htmx, manual | Media-alta |
| Comentarios | Servicio, handler, htmx, manual | Media |
| Follows | Servicio, handler, manual | Media-alta |
| Busqueda | Handler, htmx, manual | Media |

## 14. Gestion de errores encontrados

Cuando se detecte un error:

1. Se describe el problema.
2. Se indica como reproducirlo.
3. Se identifica la funcionalidad afectada.
4. Se clasifica su gravedad.
5. Se corrige.
6. Se repite la prueba.
7. Se documenta el resultado.

Niveles de gravedad:

- Critico: impide usar la aplicacion o compromete seguridad.
- Alto: rompe una funcionalidad principal.
- Medio: afecta a una funcionalidad secundaria o produce comportamiento incorrecto.
- Bajo: detalle visual, texto o mejora menor.

## 15. Pruebas antes de entrega

Antes de entregar el proyecto se ejecutara una revision final.

Checklist final:

```text
[ ] La aplicacion arranca correctamente
[ ] La base de datos se inicializa correctamente
[ ] Registro funciona
[ ] Login funciona
[ ] Logout funciona
[ ] Feed funciona
[ ] Crear publicacion funciona
[ ] Editar publicacion funciona
[ ] Eliminar publicacion funciona
[ ] Likes funcionan
[ ] Comentarios funcionan
[ ] Follows funcionan
[ ] Busqueda funciona
[ ] Las rutas privadas estan protegidas
[ ] Los tests automaticos pasan
[ ] La interfaz es usable
[ ] README actualizado
[ ] Documentacion de IA actualizada
[ ] Memoria final preparada
```

## 16. Evidencias para la memoria

Para demostrar el trabajo de testing y QA se incluiran evidencias como:

- Capturas de tests ejecutados.
- Capturas de la aplicacion funcionando.
- Tabla de pruebas manuales.
- Prompts usados para QA.
- Resumen de errores encontrados y corregidos.
- Explicacion de la integracion continua.

## 17. Relacion con CI/CD

El proceso de integracion continua debe actuar como primera barrera de calidad.

Cada subida al repositorio deberia validar:

- Compilacion.
- Tests.
- Revision estatica basica.
- Generacion correcta de templates.

Si alguna comprobacion falla, la tarea no debe considerarse completada.

## 18. Conclusión

La estrategia de testing y QA combina pruebas automaticas, pruebas manuales, revision humana y apoyo de IA.

Este enfoque permite validar las funcionalidades principales de la red social y demostrar que el codigo generado con herramientas de IA ha sido revisado, probado y documentado de forma responsable.
