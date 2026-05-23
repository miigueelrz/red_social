# PRD - Red Social con Go, htmx y templ

## 1. Resumen del producto

El proyecto consiste en desarrollar una red social web donde los usuarios puedan registrarse, iniciar sesion, crear publicaciones, interactuar con el contenido de otros usuarios mediante likes y comentarios, seguir perfiles y consultar un feed personalizado.

La aplicacion se desarrollara con Go en el backend, templ para la generacion de componentes HTML tipados y htmx para incorporar interactividad en el frontend sin necesidad de construir una SPA completa.

El proyecto se realizara apoyandose en herramientas de inteligencia artificial, principalmente Antigravity para tareas de alto nivel y opencode para tareas de implementacion, refactorizacion, pruebas y depuracion.

## 2. Objetivos

### Objetivo principal

Crear una red social funcional, mantenible y desplegable que demuestre el uso estructurado de IA durante todo el ciclo de desarrollo.

### Objetivos especificos

- Permitir que los usuarios creen una cuenta e inicien sesion.
- Permitir que los usuarios creen, editen y eliminen publicaciones.
- Mostrar un feed con publicaciones propias y de usuarios seguidos.
- Permitir interacciones sociales: likes, comentarios y seguimiento entre usuarios.
- Crear una interfaz clara, usable y responsive.
- Aplicar buenas practicas de ingenieria del software: Clean Code, SOLID, arquitectura por capas, testing y control de versiones.
- Documentar el uso de IA: prompts, agentes, decisiones, modelos elegidos, flujos de trabajo y resultados.
- Preparar la aplicacion para su despliegue.

## 3. Usuarios objetivo

### Usuario visitante

Persona que entra en la aplicacion sin haber iniciado sesion.

Necesidades:

- Ver la pantalla de bienvenida o acceso.
- Registrarse.
- Iniciar sesion.

### Usuario registrado

Persona con cuenta en la red social.

Necesidades:

- Gestionar su perfil.
- Crear publicaciones.
- Ver publicaciones de otros usuarios.
- Comentar.
- Dar like.
- Seguir o dejar de seguir usuarios.
- Buscar otros perfiles.

### Administrador

Usuario con permisos especiales para tareas de supervision basica.

Necesidades:

- Ver usuarios registrados.
- Revisar contenido.
- Eliminar publicaciones inapropiadas.

Esta figura se considera funcionalidad secundaria y solo se implementara si el tiempo del proyecto lo permite.

## 4. Alcance del MVP

El MVP debe incluir las funcionalidades minimas necesarias para considerar que la red social esta operativa.

### Funcionalidades incluidas

- Registro de usuarios.
- Inicio y cierre de sesion.
- Perfil de usuario.
- Edicion basica del perfil.
- Creacion de publicaciones.
- Edicion y eliminacion de publicaciones propias.
- Feed principal.
- Likes en publicaciones.
- Comentarios en publicaciones.
- Seguimiento entre usuarios.
- Busqueda de usuarios.
- Proteccion de rutas privadas.
- Validacion de formularios.
- Mensajes de error y confirmacion.

### Funcionalidades fuera del MVP

Estas funcionalidades podran desarrollarse si el MVP queda completado:

- Subida de imagenes.
- Notificaciones.
- Mensajes privados.
- Hashtags.
- Reportes de contenido.
- Panel de administracion completo.
- Recuperacion de contrasena por email.

## 5. Historias de usuario

### Autenticacion

- Como visitante, quiero registrarme para poder acceder a la red social.
- Como usuario registrado, quiero iniciar sesion para acceder a mis publicaciones y mi feed.
- Como usuario registrado, quiero cerrar sesion para proteger mi cuenta.

### Perfil

- Como usuario, quiero ver mi perfil para consultar mi informacion y mis publicaciones.
- Como usuario, quiero editar mi perfil para actualizar mi nombre, biografia o avatar.
- Como usuario, quiero visitar perfiles de otros usuarios para ver sus publicaciones.

### Publicaciones

- Como usuario, quiero crear una publicacion para compartir contenido.
- Como usuario, quiero editar mis publicaciones para corregir errores.
- Como usuario, quiero eliminar mis publicaciones para retirar contenido que ya no quiero mostrar.
- Como usuario, quiero ver un feed para descubrir publicaciones recientes.

### Interacciones

- Como usuario, quiero dar like a publicaciones para mostrar que me gustan.
- Como usuario, quiero quitar mi like si cambio de opinion.
- Como usuario, quiero comentar publicaciones para participar en conversaciones.
- Como usuario, quiero eliminar mis comentarios si ya no quiero que aparezcan.

### Seguimiento

- Como usuario, quiero seguir a otros usuarios para ver sus publicaciones en mi feed.
- Como usuario, quiero dejar de seguir a usuarios para controlar el contenido que veo.

### Busqueda

- Como usuario, quiero buscar otros usuarios para encontrar perfiles concretos.

## 6. Requisitos funcionales

### RF-01 Registro de usuarios

El sistema debe permitir crear una cuenta introduciendo nombre de usuario, email y contrasena.

Criterios de aceptacion:

- El email debe ser unico.
- El nombre de usuario debe ser unico.
- La contrasena debe almacenarse cifrada.
- Si hay errores de validacion, el sistema debe mostrarlos de forma clara.

### RF-02 Inicio de sesion

El sistema debe permitir iniciar sesion con email y contrasena.

Criterios de aceptacion:

- Si las credenciales son correctas, se crea una sesion.
- Si las credenciales son incorrectas, se muestra un mensaje de error.
- Las rutas privadas no deben ser accesibles sin sesion.

### RF-03 Cierre de sesion

El sistema debe permitir cerrar la sesion activa.

Criterios de aceptacion:

- La sesion debe invalidarse.
- El usuario debe ser redirigido a la pantalla de login o inicio.

### RF-04 Perfil de usuario

El sistema debe mostrar la informacion publica de un usuario.

Criterios de aceptacion:

- Se debe mostrar nombre de usuario, biografia y publicaciones.
- El usuario debe poder editar solo su propio perfil.

### RF-05 Creacion de publicaciones

El sistema debe permitir crear publicaciones de texto.

Criterios de aceptacion:

- Una publicacion debe tener contenido obligatorio.
- El contenido no debe superar el limite definido.
- Al crear una publicacion, debe aparecer en el feed.

### RF-06 Edicion y eliminacion de publicaciones

El sistema debe permitir modificar o eliminar publicaciones propias.

Criterios de aceptacion:

- Solo el autor puede editar o eliminar su publicacion.
- Tras editar una publicacion, el cambio debe verse reflejado en el feed.
- Tras eliminar una publicacion, no debe aparecer en el feed.

### RF-07 Feed

El sistema debe mostrar un listado de publicaciones recientes.

Criterios de aceptacion:

- El feed debe mostrar publicaciones ordenadas por fecha descendente.
- El feed debe incluir publicaciones propias y de usuarios seguidos.
- Cada publicacion debe mostrar autor, fecha, contenido, likes y comentarios.

### RF-08 Likes

El sistema debe permitir dar y quitar likes en publicaciones.

Criterios de aceptacion:

- Un usuario solo puede dar un like por publicacion.
- El contador de likes debe actualizarse.
- La interaccion debe realizarse con htmx sin recargar toda la pagina.

### RF-09 Comentarios

El sistema debe permitir comentar publicaciones.

Criterios de aceptacion:

- El comentario no puede estar vacio.
- Los comentarios deben mostrarse asociados a la publicacion.
- La creacion del comentario debe poder actualizar la interfaz mediante htmx.

### RF-10 Seguimiento entre usuarios

El sistema debe permitir seguir y dejar de seguir usuarios.

Criterios de aceptacion:

- Un usuario no puede seguirse a si mismo.
- No se deben duplicar relaciones de seguimiento.
- El feed debe tener en cuenta las relaciones de seguimiento.

### RF-11 Busqueda de usuarios

El sistema debe permitir buscar usuarios por nombre de usuario.

Criterios de aceptacion:

- La busqueda debe devolver coincidencias parciales.
- Los resultados deben enlazar al perfil correspondiente.
- La busqueda podra actualizar resultados usando htmx.

## 7. Requisitos no funcionales

### Rendimiento

- Las paginas principales deben cargar rapidamente.
- Las consultas del feed deben estar paginadas o preparadas para paginacion.
- Las operaciones de likes y comentarios deben actualizar solo fragmentos de interfaz cuando sea posible.

### Seguridad

- Las contrasenas deben almacenarse usando hash seguro.
- Las rutas privadas deben estar protegidas por middleware de autenticacion.
- El sistema debe validar datos en servidor.
- Se debe evitar la exposicion de informacion sensible.
- Se debe contemplar proteccion CSRF en formularios.

### Mantenibilidad

- El codigo debe organizarse por responsabilidades.
- La logica de negocio no debe estar mezclada con la capa de presentacion.
- El acceso a datos debe centralizarse en repositorios.
- Los componentes templ deben ser reutilizables.

### Usabilidad

- La interfaz debe ser clara y consistente.
- Los formularios deben mostrar errores comprensibles.
- Las acciones principales deben estar visibles.
- La aplicacion debe adaptarse a escritorio y movil.

### Escalabilidad

- La arquitectura debe permitir anadir nuevas funcionalidades sin reescribir el nucleo.
- El modelo de datos debe permitir incorporar notificaciones, mensajes privados o multimedia en futuras iteraciones.

### Testabilidad

- La logica de negocio debe poder probarse con tests unitarios.
- Los handlers HTTP deben poder probarse de forma aislada.
- Las funcionalidades principales deben contar con pruebas manuales documentadas.

## 8. Stack tecnologico

### Backend

- Go como lenguaje principal.
- Servidor HTTP usando la libreria estandar o router ligero.
- Arquitectura por capas: handlers, services, repositories y models.

### Frontend

- templ para componentes HTML.
- htmx para interacciones parciales.
- CSS propio o framework ligero si se considera necesario.

### Base de datos

- PostgreSQL como opcion recomendada para el proyecto final.
- SQLite como alternativa para desarrollo local si se busca simplicidad.

### Herramientas de desarrollo

- Git para control de versiones.
- GitHub, GitLab o Bitbucket para repositorio remoto.
- Docker y Docker Compose para entorno de ejecucion.
- GitHub Actions o GitLab CI para integracion continua.

### Herramientas de IA

- Antigravity para planificacion, diseno, arquitectura, revision global y documentacion.
- opencode para implementacion, refactorizacion, generacion de tests y depuracion.

## 9. Arquitectura propuesta

La aplicacion seguira una arquitectura por capas:

- Capa de presentacion: templates creados con templ y actualizaciones parciales con htmx.
- Capa HTTP: handlers de Go encargados de recibir peticiones, validar entrada basica y devolver respuestas.
- Capa de servicios: logica de negocio.
- Capa de repositorios: acceso a base de datos.
- Capa de modelos: entidades principales del dominio.
- Middleware: autenticacion, sesiones, logging y seguridad.

Estructura inicial propuesta:

```text
cmd/web/
internal/handlers/
internal/services/
internal/repositories/
internal/models/
internal/auth/
internal/middleware/
internal/templates/
migrations/
docs/
docs/ia/
```

## 10. Modelo de datos inicial

### users

- id
- username
- email
- password_hash
- bio
- avatar_url
- created_at
- updated_at

### sessions

- id
- user_id
- token_hash
- expires_at
- created_at

### posts

- id
- user_id
- content
- created_at
- updated_at

### comments

- id
- post_id
- user_id
- content
- created_at
- updated_at

### likes

- id
- post_id
- user_id
- created_at

### follows

- id
- follower_id
- followed_id
- created_at

## 11. Metodologia de trabajo

Se utilizara una combinacion de Extreme Programming y Kanban.

Extreme Programming aporta:

- Historias de usuario pequenas.
- Desarrollo incremental.
- Testing frecuente.
- Refactorizacion continua.
- Revision constante del codigo.

Kanban aporta:

- Visualizacion del estado de las tareas.
- Priorizacion sencilla.
- Flujo continuo de trabajo.
- Adaptacion a cambios durante el desarrollo.

Estados propuestos para las tareas:

```text
Backlog
Pendiente
En desarrollo
En revision IA
En pruebas
Completado
```

## 12. Uso de IA en el desarrollo

El uso de IA sera documentado durante todo el proyecto.

### Agentes propuestos

- Agente de producto: definicion del PRD, historias de usuario y criterios de aceptacion.
- Agente de arquitectura: estructura del proyecto, patrones, base de datos y decisiones tecnicas.
- Agente backend: implementacion de handlers, servicios, repositorios y autenticacion.
- Agente frontend: componentes templ, interacciones htmx y diseno de interfaz.
- Agente QA: tests, casos de prueba, deteccion de errores y validacion de requisitos.
- Agente DevOps: Docker, CI/CD, variables de entorno y despliegue.
- Agente documental: memoria, registro de prompts y explicacion de decisiones.

### Flujo de trabajo con IA

1. Definir una historia de usuario.
2. Pedir a IA el diseno tecnico de la solucion.
3. Generar o modificar codigo con IA.
4. Revisar el codigo generado.
5. Ejecutar pruebas.
6. Pedir a IA una revision de calidad.
7. Documentar prompt, resultado y decision final.

## 13. Testing y QA

### Pruebas automaticas

- Tests unitarios para servicios.
- Tests de handlers HTTP.
- Tests de repositorios cuando sea viable.
- Validacion de reglas de negocio: likes unicos, permisos de edicion, follows no duplicados.

### Pruebas manuales

- Registro de usuario.
- Login y logout.
- Crear publicacion.
- Editar publicacion.
- Eliminar publicacion.
- Dar y quitar like.
- Crear comentario.
- Seguir y dejar de seguir usuario.
- Buscar usuario.
- Acceso a rutas protegidas sin sesion.

### QA con IA

La IA se utilizara para:

- Revisar posibles bugs.
- Detectar codigo duplicado.
- Proponer mejoras de arquitectura.
- Generar casos de prueba.
- Revisar accesibilidad y usabilidad.

## 14. Criterios de exito

El proyecto se considerara exitoso si:

- La aplicacion permite completar los flujos principales de una red social.
- El codigo esta organizado y es mantenible.
- La interfaz es clara y funcional.
- Existen pruebas automaticas o manuales documentadas.
- La aplicacion puede ejecutarse localmente de forma sencilla.
- Se documenta el uso de IA de forma completa.
- La memoria final explica decisiones, agentes, prompts, arquitectura, testing y despliegue.

## 15. Riesgos

### Riesgo: exceso de funcionalidades

Mitigacion:

- Priorizar el MVP antes de implementar extras.

### Riesgo: dependencia excesiva de IA sin comprender el codigo

Mitigacion:

- Revisar cada cambio generado.
- Documentar decisiones.
- Explicar la arquitectura en la memoria.

### Riesgo: errores de seguridad en autenticacion

Mitigacion:

- Usar hashing seguro para contrasenas.
- Proteger rutas privadas.
- Validar datos en servidor.

### Riesgo: falta de tiempo para despliegue

Mitigacion:

- Preparar Docker desde fases tempranas.
- Mantener una version ejecutable durante todo el desarrollo.

## 16. Roadmap inicial

### Fase 1: Planificacion

- PRD.
- Historias de usuario.
- Arquitectura.
- Modelo de datos.
- Flujo de agentes IA.

### Fase 2: Base tecnica

- Inicializar proyecto Go.
- Configurar templ.
- Configurar htmx.
- Configurar base de datos.
- Crear migraciones iniciales.

### Fase 3: Autenticacion

- Registro.
- Login.
- Logout.
- Middleware de rutas privadas.

### Fase 4: Funcionalidades sociales

- Perfil.
- Publicaciones.
- Feed.
- Likes.
- Comentarios.
- Follows.
- Busqueda.

### Fase 5: Calidad

- Tests.
- Revision con IA.
- Refactorizacion.
- Pruebas manuales.

### Fase 6: Entrega

- Docker.
- CI/CD.
- Despliegue.
- Memoria final.
- ZIP del codigo.
- Grupo.txt.
