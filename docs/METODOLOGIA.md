# Metodologia de trabajo

## 1. Objetivo del documento

Este documento describe la metodologia de desarrollo que se utilizara durante el proyecto de red social con Go, htmx y templ.

La metodologia elegida debe permitir organizar el trabajo, aprovechar las herramientas de inteligencia artificial de forma estructurada y mantener la calidad del codigo y de la documentacion.

## 2. Metodologia elegida

La metodologia principal del proyecto sera una combinacion de:

- Extreme Programming.
- Kanban.

Esta combinacion encaja con el proyecto porque permite desarrollar por incrementos pequenos, validar cada funcionalidad de forma continua y adaptar el backlog segun las necesidades del equipo.

## 3. Justificacion de Extreme Programming

Extreme Programming, tambien conocida como XP, se centra en ciclos cortos de desarrollo, pruebas frecuentes, refactorizacion continua y entrega incremental de valor.

Se considera adecuada para este proyecto por los siguientes motivos:

- El proyecto tiene muchas funcionalidades pequenas que pueden desarrollarse de forma incremental.
- El uso de IA genera codigo rapidamente, por lo que es necesario revisar, probar y refactorizar de forma continua.
- Las historias de usuario permiten conectar requisitos funcionales con tareas concretas.
- Facilita mantener una version funcional de la aplicacion durante todo el desarrollo.
- Favorece la calidad mediante testing y revision constante.

Practicas de XP que se aplicaran:

- Historias de usuario.
- Desarrollo iterativo.
- Testing frecuente.
- Refactorizacion continua.
- Integracion continua.
- Codigo simple y mantenible.
- Revision del codigo generado por IA.

## 4. Justificacion de Kanban

Kanban se utilizara para visualizar el flujo de trabajo y controlar el estado de cada tarea.

Se considera adecuado porque:

- Permite saber en todo momento que tareas estan pendientes, en curso, en revision o completadas.
- Es flexible y no obliga a trabajar con sprints cerrados.
- Encaja bien con un proyecto academico donde pueden aparecer ajustes durante el desarrollo.
- Facilita priorizar el MVP antes de implementar funcionalidades extra.
- Ayuda a evitar empezar demasiadas tareas a la vez.

## 5. Tablero de trabajo

El tablero Kanban tendra los siguientes estados:

```text
Backlog
Pendiente
En desarrollo
En revision IA
En pruebas
Completado
```

### Backlog

Contiene todas las funcionalidades, mejoras, tareas tecnicas y tareas documentales identificadas.

### Pendiente

Tareas seleccionadas como proximas a realizar.

### En desarrollo

Tareas que se estan implementando activamente.

### En revision IA

Tareas ya implementadas que se revisan con ayuda de IA para detectar errores, mejorar estructura o comprobar requisitos.

### En pruebas

Tareas pendientes de validacion mediante tests automaticos o pruebas manuales.

### Completado

Tareas terminadas, revisadas, probadas y documentadas.

## 6. Flujo de trabajo por tarea

Cada tarea seguira este proceso:

1. Seleccionar una historia de usuario o tarea tecnica del backlog.
2. Revisar los criterios de aceptacion.
3. Pedir apoyo a IA si hace falta definir el enfoque tecnico.
4. Implementar la funcionalidad.
5. Revisar el codigo generado o modificado.
6. Ejecutar pruebas.
7. Corregir errores.
8. Documentar la interaccion con IA si ha sido relevante.
9. Marcar la tarea como completada.

## 7. Definition of Ready

Una tarea esta lista para empezar cuando cumple estas condiciones:

- Tiene una descripcion clara.
- Esta relacionada con una historia de usuario, requisito tecnico o entregable del proyecto.
- Tiene criterios de aceptacion.
- Se conocen los ficheros o modulos aproximados que puede afectar.
- Se sabe si requiere cambios en backend, frontend, base de datos, documentacion o despliegue.
- Si se usara IA, esta claro que agente intervendra.

## 8. Definition of Done

Una tarea se considera terminada cuando cumple estas condiciones:

- La funcionalidad esta implementada.
- El codigo compila.
- No rompe funcionalidades existentes.
- Los tests relevantes se han ejecutado.
- La interfaz se ha comprobado si la tarea afecta al frontend.
- Se han validado los criterios de aceptacion.
- El codigo mantiene la arquitectura acordada.
- Se ha documentado la decision o prompt de IA si corresponde.
- La tarea puede explicarse en la memoria del proyecto.

## 9. Backlog inicial

### Fase 1: Documentacion y planificacion

- Crear PRD.
- Definir agentes IA y flujo entre ellos.
- Definir metodologia.
- Definir arquitectura.
- Definir modelo de datos.
- Definir estrategia de testing y QA.
- Definir estrategia de despliegue.

### Fase 2: Base tecnica del proyecto

- Inicializar proyecto Go.
- Crear estructura de carpetas.
- Configurar templ.
- Configurar htmx.
- Configurar estilos base.
- Configurar base de datos.
- Crear migraciones iniciales.
- Crear sistema de configuracion por variables de entorno.

### Fase 3: Autenticacion

- Registro de usuarios.
- Login.
- Logout.
- Hash seguro de contrasenas.
- Sesiones.
- Middleware de autenticacion.
- Proteccion de rutas privadas.

### Fase 4: Perfil de usuario

- Vista de perfil.
- Edicion de perfil.
- Listado de publicaciones del usuario.
- Vista publica de otros perfiles.

### Fase 5: Publicaciones

- Crear publicacion.
- Mostrar publicaciones.
- Editar publicacion propia.
- Eliminar publicacion propia.
- Validaciones de contenido.

### Fase 6: Feed

- Feed principal.
- Orden por fecha.
- Mostrar publicaciones propias.
- Mostrar publicaciones de usuarios seguidos.
- Preparar paginacion si es necesario.

### Fase 7: Interacciones sociales

- Dar like.
- Quitar like.
- Mostrar contador de likes.
- Crear comentarios.
- Mostrar comentarios.
- Eliminar comentarios propios.

### Fase 8: Seguimiento y busqueda

- Seguir usuario.
- Dejar de seguir usuario.
- Evitar seguirse a uno mismo.
- Buscar usuarios.
- Mostrar resultados con htmx.

### Fase 9: Calidad

- Tests unitarios.
- Tests de handlers.
- Pruebas manuales.
- Revision de seguridad basica.
- Refactorizacion.
- Revision final con IA.

### Fase 10: Entrega

- Dockerfile.
- docker-compose.yml.
- CI/CD.
- Despliegue.
- README.
- Memoria final.
- Grupo.txt.
- ZIP del proyecto.

## 10. Priorizacion

La prioridad principal sera completar el MVP antes de anadir funcionalidades extra.

Prioridad alta:

- Registro.
- Login/logout.
- Perfil.
- Publicaciones.
- Feed.
- Likes.
- Comentarios.
- Follows.
- Busqueda.
- Documentacion de IA.

Prioridad media:

- Docker.
- CI/CD.
- Despliegue.
- Mejoras visuales.
- Tests adicionales.

Prioridad baja:

- Notificaciones.
- Mensajes privados.
- Subida de imagenes.
- Panel de administracion.
- Reportes.

## 11. Gestion del repositorio

Se utilizara Git como sistema de control de versiones.

Buenas practicas:

- Commits pequenos y descriptivos.
- Un commit por funcionalidad o cambio coherente.
- Evitar mezclar cambios de documentacion, backend y frontend sin relacion.
- Revisar cambios antes de confirmar.
- Mantener el repositorio en un estado ejecutable.

Formato recomendado de commits:

```text
docs: add project PRD
docs: define AI agents workflow
feat: add user registration
feat: add post creation
fix: validate empty comments
test: add auth service tests
chore: configure docker compose
```

## 12. Flujo de ramas

Para un proyecto academico, se propone un flujo sencillo:

```text
main
feature/auth
feature/posts
feature/feed
feature/social-interactions
docs/final-report
```

La rama `main` debe mantenerse estable. Las ramas de funcionalidad se usaran para desarrollar bloques concretos antes de integrarlos.

Si el equipo decide simplificar, tambien puede trabajar directamente en `main`, siempre que los commits sean frecuentes y el proyecto se mantenga funcional.

## 13. Integracion continua

La integracion continua se utilizara para validar automaticamente que el proyecto compila y que los tests pasan.

Comprobaciones recomendadas:

- Descargar dependencias.
- Generar templates si es necesario.
- Ejecutar `go test ./...`.
- Ejecutar `go vet ./...`.
- Construir la aplicacion.

Ejemplo de pipeline:

```text
checkout
setup-go
install-dependencies
generate-templ
go-test
go-vet
go-build
```

## 14. Testing dentro de la metodologia

El testing se integra en cada fase del desarrollo.

Tipos de pruebas:

- Tests unitarios para servicios.
- Tests de handlers HTTP.
- Tests de repositorios cuando sea viable.
- Pruebas manuales de flujos completos.
- Revision visual de la interfaz.

Cada funcionalidad importante debe tener al menos una forma de validacion, automatica o manual.

## 15. QA con IA

La IA se utilizara como apoyo para QA.

Usos previstos:

- Generar casos de prueba.
- Revisar criterios de aceptacion.
- Detectar errores de permisos.
- Buscar duplicacion o acoplamiento excesivo.
- Revisar posibles problemas de seguridad.
- Proponer mejoras de usabilidad.

Ejemplo de prompt QA:

```text
Actua como agente QA del proyecto. Revisa esta funcionalidad de likes en una red social hecha con Go, htmx y templ. Comprueba si cumple los criterios de aceptacion, si hay problemas de seguridad o permisos, y propone tests automaticos y manuales.
```

## 16. Uso de IA dentro de la metodologia

El uso de IA se integrara en el flujo de desarrollo, no como una actividad separada.

Para cada tarea relevante:

1. Se define el objetivo.
2. Se selecciona el agente adecuado.
3. Se redacta un prompt con contexto.
4. Se revisa la respuesta.
5. Se aplica el cambio si procede.
6. Se prueban los resultados.
7. Se documenta la interaccion.

## 17. Gestion de prompts

Los prompts importantes se guardaran en:

```text
docs/ia/prompts/
```

Cada prompt debe incluir:

- Fecha.
- Herramienta usada.
- Modelo usado.
- Agente asociado.
- Objetivo.
- Prompt.
- Resumen de respuesta.
- Decision tomada.
- Resultado.

## 18. Gestion de decisiones tecnicas

Las decisiones tecnicas importantes se registraran en:

```text
docs/ia/decisiones/
```

Ejemplos de decisiones a documentar:

- Eleccion de PostgreSQL frente a SQLite.
- Eleccion de arquitectura por capas.
- Uso de templ frente a plantillas HTML tradicionales.
- Uso de htmx frente a framework SPA.
- Sistema de autenticacion y sesiones.
- Estrategia de despliegue.

## 19. Revision y refactorizacion

La refactorizacion sera continua y se realizara despues de tener una funcionalidad funcionando.

Objetivos de la refactorizacion:

- Reducir duplicacion.
- Mejorar nombres.
- Separar responsabilidades.
- Simplificar funciones grandes.
- Mantener handlers ligeros.
- Mover logica de negocio a servicios.
- Mantener el acceso a datos en repositorios.

La IA puede proponer mejoras, pero el equipo decidira si se aceptan.

## 20. Criterios de calidad

El proyecto debe cumplir estos criterios:

- Codigo claro y mantenible.
- Arquitectura coherente.
- Baja duplicacion.
- Separacion de responsabilidades.
- Validaciones en servidor.
- Seguridad basica en autenticacion.
- Interfaz usable.
- Tests o pruebas documentadas.
- Despliegue o instrucciones claras de ejecucion.
- Documentacion completa del uso de IA.

## 21. Relacion con los agentes IA

La metodologia se coordina con los agentes definidos en `AGENTES_IA.md`.

| Fase | Agente principal | Resultado |
| --- | --- | --- |
| Planificacion | Producto | PRD e historias |
| Diseno tecnico | Arquitectura | Arquitectura y modelo de datos |
| Implementacion backend | Backend | Handlers, servicios y repositorios |
| Implementacion frontend | Frontend | Componentes templ e interacciones htmx |
| Revision | QA | Tests y checklist |
| Entorno y despliegue | DevOps | Docker, CI/CD y despliegue |
| Memoria | Documental | Documentacion final |

## 22. Indicadores de avance

Para medir el progreso se usaran los siguientes indicadores:

- Numero de historias completadas.
- Numero de requisitos del PRD implementados.
- Numero de pruebas realizadas.
- Estado del MVP.
- Estado de la documentacion.
- Estado del despliegue.
- Calidad de la memoria final.

## 23. Adaptacion durante el proyecto

La metodologia podra ajustarse si aparecen problemas de tiempo o complejidad.

Reglas de adaptacion:

- No anadir extras si el MVP no esta completo.
- Reducir alcance antes que entregar funcionalidades rotas.
- Priorizar funcionalidades demostrables en la exposicion.
- Mantener siempre la documentacion de IA actualizada.
- Conservar una version ejecutable del proyecto.

## 24. Conclusión

La combinacion de Extreme Programming y Kanban permite desarrollar la red social de forma incremental, controlada y adaptable.

Al integrarse con agentes de IA, esta metodologia ayuda a organizar el trabajo, mantener trazabilidad, revisar la calidad y demostrar que la inteligencia artificial se ha usado de forma razonada durante todo el proyecto.
