# Arquitectura del proyecto

## 1. Objetivo del documento

Este documento describe la arquitectura tecnica propuesta para la red social desarrollada con Go, htmx y templ.

Su objetivo es justificar las decisiones de diseno, explicar la separacion de responsabilidades y demostrar la aplicacion de buenas practicas de ingenieria del software como Clean Code, principios SOLID, patrones de diseno, seguridad, mantenibilidad y escalabilidad.

## 2. Vision general

La aplicacion sera una red social web tradicional renderizada desde servidor, con interacciones parciales mediante htmx.

El backend estara desarrollado en Go y se encargara de:

- Gestionar peticiones HTTP.
- Aplicar reglas de negocio.
- Validar datos.
- Gestionar sesiones.
- Acceder a la base de datos.
- Renderizar vistas mediante templ.

El frontend se construira con:

- templ para definir componentes HTML tipados desde Go.
- htmx para actualizar fragmentos de interfaz sin recargar toda la pagina.
- CSS para el diseno visual y responsive.

## 3. Estilo arquitectonico

Se utilizara una arquitectura por capas.

Capas principales:

- Capa de presentacion.
- Capa HTTP.
- Capa de servicios.
- Capa de repositorios.
- Capa de dominio/modelos.
- Capa de infraestructura.

Esta arquitectura permite separar responsabilidades y reducir el acoplamiento entre partes del sistema.

```text
Usuario
  |
  v
Navegador + htmx
  |
  v
Handlers HTTP
  |
  v
Servicios
  |
  v
Repositorios
  |
  v
Base de datos
```

## 4. Stack tecnologico

### Backend

- Go como lenguaje principal.
- Servidor HTTP con libreria estandar o router ligero.
- Middlewares para autenticacion, logging y seguridad.
- bcrypt o libreria equivalente para hash de contrasenas.

### Frontend

- templ para plantillas y componentes.
- htmx para interacciones dinamicas.
- CSS propio o framework ligero si se decide posteriormente.

### Base de datos

- PostgreSQL como base de datos recomendada.
- SQLite como alternativa para desarrollo local si se prioriza simplicidad.

La opcion recomendada para entrega es PostgreSQL, porque representa mejor un entorno real y facilita justificar escalabilidad y despliegue.

### Infraestructura

- Docker para empaquetar la aplicacion.
- Docker Compose para entorno local con base de datos.
- GitHub Actions o GitLab CI para integracion continua.
- Render, Railway, Fly.io o VPS para despliegue.

## 5. Estructura de carpetas propuesta

```text
cmd/
  web/
    main.go

internal/
  config/
  handlers/
  services/
  repositories/
  models/
  auth/
  middleware/
  templates/
  validation/

migrations/

static/
  css/
  js/
  img/

docs/
  PRD.md
  AGENTES_IA.md
  METODOLOGIA.md
  ARQUITECTURA.md
  ia/

tests/
```

## 6. Responsabilidad de cada carpeta

### `cmd/web`

Contiene el punto de entrada de la aplicacion.

Responsabilidades:

- Cargar configuracion.
- Conectar con la base de datos.
- Inicializar repositorios.
- Inicializar servicios.
- Inicializar handlers.
- Registrar rutas.
- Arrancar el servidor HTTP.

### `internal/config`

Gestiona la configuracion de la aplicacion.

Responsabilidades:

- Leer variables de entorno.
- Definir puerto del servidor.
- Definir cadena de conexion a base de datos.
- Configurar entorno de ejecucion.

### `internal/handlers`

Contiene los handlers HTTP.

Responsabilidades:

- Recibir peticiones.
- Leer parametros y formularios.
- Ejecutar validaciones simples.
- Llamar a servicios.
- Decidir que vista o fragmento devolver.

Los handlers no deben contener logica de negocio compleja.

### `internal/services`

Contiene la logica de negocio.

Responsabilidades:

- Aplicar reglas del dominio.
- Coordinar operaciones entre repositorios.
- Comprobar permisos.
- Evitar duplicados.
- Gestionar casos de uso.

Ejemplos:

- Registrar usuario.
- Crear publicacion.
- Dar like.
- Seguir usuario.
- Construir feed.

### `internal/repositories`

Contiene el acceso a datos.

Responsabilidades:

- Ejecutar consultas SQL.
- Crear, leer, actualizar y eliminar entidades.
- Encapsular detalles de la base de datos.

Los repositorios no deben decidir reglas de negocio. Solo deben acceder y persistir datos.

### `internal/models`

Contiene las entidades principales del dominio.

Ejemplos:

- User.
- Post.
- Comment.
- Like.
- Follow.
- Session.

### `internal/auth`

Contiene funcionalidades relacionadas con autenticacion.

Responsabilidades:

- Hash de contrasenas.
- Verificacion de contrasenas.
- Creacion de tokens de sesion.
- Gestion de cookies de sesion.

### `internal/middleware`

Contiene middleware reutilizable.

Ejemplos:

- Autenticacion.
- Logging.
- Recuperacion de panics.
- Proteccion CSRF si se implementa.

### `internal/templates`

Contiene componentes y vistas creados con templ.

Ejemplos:

- Layout base.
- Formulario de login.
- Formulario de registro.
- Feed.
- Tarjeta de publicacion.
- Perfil de usuario.
- Fragmentos htmx.

### `internal/validation`

Contiene validaciones reutilizables.

Ejemplos:

- Validar email.
- Validar longitud de contrasena.
- Validar contenido de publicaciones.
- Validar nombre de usuario.

### `migrations`

Contiene scripts SQL para crear y modificar la base de datos.

### `static`

Contiene recursos estaticos servidos al navegador.

Ejemplos:

- CSS.
- JavaScript minimo si fuera necesario.
- Imagenes.

## 7. Modelo de dominio inicial

### User

Representa a un usuario registrado.

Campos principales:

- id
- username
- email
- password_hash
- bio
- avatar_url
- created_at
- updated_at

### Session

Representa una sesion activa.

Campos principales:

- id
- user_id
- token_hash
- expires_at
- created_at

### Post

Representa una publicacion.

Campos principales:

- id
- user_id
- content
- created_at
- updated_at

### Comment

Representa un comentario en una publicacion.

Campos principales:

- id
- post_id
- user_id
- content
- created_at
- updated_at

### Like

Representa la relacion entre un usuario y una publicacion que le gusta.

Campos principales:

- id
- post_id
- user_id
- created_at

### Follow

Representa la relacion de seguimiento entre dos usuarios.

Campos principales:

- id
- follower_id
- followed_id
- created_at

## 8. Modelo relacional propuesto

```text
users
  id PK
  username UNIQUE
  email UNIQUE
  password_hash
  bio
  avatar_url
  created_at
  updated_at

sessions
  id PK
  user_id FK -> users.id
  token_hash
  expires_at
  created_at

posts
  id PK
  user_id FK -> users.id
  content
  created_at
  updated_at

comments
  id PK
  post_id FK -> posts.id
  user_id FK -> users.id
  content
  created_at
  updated_at

likes
  id PK
  post_id FK -> posts.id
  user_id FK -> users.id
  created_at
  UNIQUE(post_id, user_id)

follows
  id PK
  follower_id FK -> users.id
  followed_id FK -> users.id
  created_at
  UNIQUE(follower_id, followed_id)
```

## 9. Patrones de diseno utilizados

### Repository Pattern

Los repositorios encapsulan el acceso a la base de datos.

Ventajas:

- Evita SQL disperso por toda la aplicacion.
- Facilita pruebas.
- Reduce acoplamiento entre logica de negocio y almacenamiento.
- Permite cambiar detalles de persistencia con menor impacto.

### Service Layer

Los servicios contienen la logica de negocio.

Ventajas:

- Mantiene handlers simples.
- Centraliza reglas de negocio.
- Facilita tests unitarios.
- Evita duplicacion de logica.

### Dependency Injection

Las dependencias se pasaran explicitamente al crear handlers y servicios.

Ventajas:

- Mejora testabilidad.
- Hace visibles las dependencias.
- Evita variables globales innecesarias.

### Middleware Pattern

Los middlewares se usaran para preocupaciones transversales.

Ejemplos:

- Autenticacion.
- Logging.
- Proteccion de rutas privadas.
- Recuperacion de errores.

### Component Pattern

templ permitira dividir la interfaz en componentes reutilizables.

Ejemplos:

- PostCard.
- CommentList.
- LikeButton.
- FollowButton.
- ProfileHeader.

## 10. Aplicacion de principios SOLID

### Single Responsibility Principle

Cada modulo tendra una responsabilidad clara.

Ejemplos:

- Un handler gestiona HTTP.
- Un servicio aplica reglas de negocio.
- Un repositorio accede a datos.
- Un componente templ renderiza UI.

### Open/Closed Principle

La estructura debe permitir anadir nuevas funcionalidades sin modificar excesivamente codigo existente.

Ejemplo:

- Se podran anadir notificaciones creando nuevos servicios, repositorios y componentes sin reescribir autenticacion o publicaciones.

### Liskov Substitution Principle

Las interfaces de repositorios y servicios deberan poder sustituirse por implementaciones alternativas, por ejemplo mocks en tests.

### Interface Segregation Principle

Las interfaces deben ser pequenas y especificas.

Ejemplo:

- `UserRepository` no debe incluir metodos de publicaciones.
- `PostRepository` no debe incluir metodos de sesiones.

### Dependency Inversion Principle

Las capas superiores dependeran de abstracciones y no directamente de detalles de infraestructura siempre que tenga sentido.

Ejemplo:

- Un servicio de usuarios puede depender de una interfaz `UserRepository`.

## 11. Clean Code

Buenas practicas a aplicar:

- Nombres claros y descriptivos.
- Funciones cortas.
- Evitar duplicacion.
- Mantener handlers ligeros.
- No mezclar SQL con vistas.
- No mezclar logica de negocio con templates.
- Gestionar errores de forma explicita.
- Evitar variables globales innecesarias.
- Mantener comentarios solo cuando aporten contexto real.
- Escribir codigo que pueda entenderse durante la exposicion.

## 12. Flujo de una peticion

Ejemplo: crear una publicacion.

```text
1. El usuario envia el formulario desde el navegador.
2. htmx o el navegador realiza una peticion HTTP.
3. El middleware comprueba que el usuario esta autenticado.
4. El handler lee y valida el contenido.
5. El handler llama al servicio de publicaciones.
6. El servicio comprueba reglas de negocio.
7. El servicio llama al repositorio.
8. El repositorio guarda la publicacion en la base de datos.
9. El servicio devuelve el resultado al handler.
10. El handler renderiza un componente templ.
11. htmx actualiza el fragmento correspondiente de la pagina.
```

## 13. Uso de htmx

htmx se utilizara para mejorar la experiencia de usuario sin construir una SPA.

Interacciones previstas:

- Crear publicaciones sin recargar toda la pagina.
- Dar y quitar like actualizando solo el boton o contador.
- Crear comentarios actualizando la lista de comentarios.
- Seguir o dejar de seguir actualizando el boton de follow.
- Buscar usuarios mostrando resultados parciales.

Ventajas:

- Menor complejidad frontend.
- Mejor integracion con Go y templ.
- Renderizado desde servidor.
- Menos JavaScript manual.
- Flujo mas sencillo de mantener.

## 14. Uso de templ

templ se utilizara para crear componentes HTML tipados.

Ventajas:

- Componentes reutilizables.
- Mejor integracion con Go.
- Comprobacion en compilacion.
- Separacion clara entre estructura de vista y logica del servidor.
- Facilidad para renderizar fragmentos htmx.

Componentes propuestos:

```text
Layout
Navbar
LoginForm
RegisterForm
Feed
PostCard
PostForm
CommentList
CommentForm
LikeButton
FollowButton
ProfileHeader
SearchResults
ErrorMessage
```

## 15. Seguridad

La seguridad se considera parte de la arquitectura.

Medidas previstas:

- Hash seguro de contrasenas.
- Cookies de sesion con atributos seguros cuando proceda.
- Validacion de entrada en servidor.
- Proteccion de rutas privadas.
- Comprobacion de permisos para editar o borrar contenido.
- Evitar que un usuario siga a si mismo.
- Evitar likes duplicados mediante restriccion unica.
- Evitar follows duplicados mediante restriccion unica.
- Escapado de contenido renderizado.
- Proteccion CSRF en formularios si se implementa.

## 16. Gestion de errores

El sistema debe gestionar errores de forma controlada.

Tipos de errores:

- Errores de validacion.
- Errores de autenticacion.
- Errores de permisos.
- Errores de base de datos.
- Errores inesperados del servidor.

Buenas practicas:

- No mostrar detalles internos al usuario.
- Registrar errores relevantes.
- Mostrar mensajes claros en formularios.
- Usar codigos HTTP adecuados.
- Mantener respuestas htmx coherentes.

## 17. Configuracion

La aplicacion se configurara mediante variables de entorno.

Variables previstas:

```text
APP_ENV
APP_PORT
DATABASE_URL
SESSION_SECRET
CSRF_SECRET
```

Ventajas:

- Separacion entre codigo y configuracion.
- Facilidad para despliegue.
- Mejor seguridad.
- Compatibilidad con Docker y plataformas cloud.

## 18. Testing y testabilidad

La arquitectura se disena para facilitar pruebas.

Estrategias:

- Servicios probables mediante repositorios mock.
- Handlers probables con `httptest`.
- Repositorios probables con base de datos de test.
- Validaciones probables de forma aislada.

Tipos de tests:

- Tests unitarios.
- Tests de integracion.
- Tests de handlers.
- Pruebas manuales documentadas.

## 19. Escalabilidad

Aunque el proyecto es academico, la arquitectura permite evolucionar.

Posibles ampliaciones:

- Notificaciones.
- Mensajes privados.
- Subida de imagenes.
- Moderacion.
- Panel de administracion.
- Paginacion avanzada.
- Busqueda avanzada.

La separacion por capas facilita anadir estos modulos sin reescribir el nucleo.

## 20. Rendimiento

Medidas previstas:

- Consultas SQL controladas.
- Indices en campos importantes.
- Paginacion del feed si el volumen crece.
- Actualizaciones parciales con htmx.
- Evitar cargar mas datos de los necesarios.

Indices recomendados:

```text
users(username)
users(email)
posts(user_id, created_at)
comments(post_id, created_at)
likes(post_id, user_id)
follows(follower_id, followed_id)
```

## 21. Accesibilidad y usabilidad

La interfaz debe ser clara y accesible.

Buenas practicas:

- Formularios con labels.
- Mensajes de error comprensibles.
- Botones con texto claro.
- Contraste suficiente.
- Diseno responsive.
- Navegacion sencilla.
- Estados vacios utiles.

## 22. Integracion continua

La arquitectura debe poder validarse automaticamente.

Pipeline recomendado:

```text
1. Descargar codigo.
2. Instalar Go.
3. Instalar dependencias.
4. Generar componentes templ.
5. Ejecutar go test ./...
6. Ejecutar go vet ./...
7. Construir binario.
```

## 23. Despliegue

La aplicacion debera poder desplegarse en un entorno externo o ejecutarse localmente con Docker.

Opciones:

- Render.
- Railway.
- Fly.io.
- VPS con Docker.
- Docker Compose local.

El despliegue debera documentar:

- Variables de entorno.
- Comandos de ejecucion.
- Migraciones.
- Conexion a base de datos.
- URL final si existe.

## 24. Decisiones arquitectonicas principales

### Decision 1: Go como backend

Go es adecuado por su rendimiento, simplicidad, tipado estatico y facilidad para crear servidores HTTP.

### Decision 2: templ para vistas

templ permite crear componentes HTML tipados y reutilizables, integrados con Go.

### Decision 3: htmx para interactividad

htmx permite crear una experiencia dinamica sin introducir la complejidad de una SPA.

### Decision 4: arquitectura por capas

La arquitectura por capas mejora mantenibilidad, testabilidad y separacion de responsabilidades.

### Decision 5: Repository Pattern

Este patron evita acoplar la logica de negocio a detalles concretos de SQL.

### Decision 6: PostgreSQL como base de datos recomendada

PostgreSQL aporta robustez, restricciones, relaciones e integridad de datos adecuadas para una red social.

## 25. Relacion con la documentacion del proyecto

Esta arquitectura se conecta con:

- `PRD.md`: define que debe hacer la aplicacion.
- `AGENTES_IA.md`: define que agentes apoyan cada parte.
- `METODOLOGIA.md`: define como se organiza el trabajo.
- `TESTING_QA.md`: definira como se valida la calidad.
- `DESPLIEGUE.md`: definira como se ejecuta y despliega.

## 26. Conclusión

La arquitectura propuesta permite construir una red social funcional, mantenible y escalable dentro del alcance del proyecto final.

La separacion por capas, el uso de patrones como Repository y Service Layer, la aplicacion de principios SOLID y el apoyo de templ y htmx permiten desarrollar una aplicacion clara, modular y adecuada para demostrar un proceso de desarrollo asistido por IA.
