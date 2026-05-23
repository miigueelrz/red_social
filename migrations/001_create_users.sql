-- =============================================================
-- Migración 001: Crear tabla de usuarios
-- =============================================================
-- Esta migración crea la tabla principal de usuarios de la red
-- social. Incluye todos los campos definidos en la arquitectura,
-- restricciones de unicidad, índices de rendimiento y triggers
-- para actualizar automáticamente updated_at.
-- =============================================================

-- Habilitar la extensión para UUIDs si se prefieren en el futuro
-- CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- -------------------------------------------------------------
-- Tabla: users
-- -------------------------------------------------------------
CREATE TABLE IF NOT EXISTS users (
    id            BIGSERIAL       PRIMARY KEY,
    username      VARCHAR(50)     NOT NULL UNIQUE,
    email         VARCHAR(255)    NOT NULL UNIQUE,
    password_hash VARCHAR(255)    NOT NULL,
    bio           TEXT            DEFAULT '',
    avatar_url    VARCHAR(500)    DEFAULT '',
    created_at    TIMESTAMPTZ     NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ     NOT NULL DEFAULT NOW()
);

-- -------------------------------------------------------------
-- Índices de rendimiento
-- Los campos username y email se buscan frecuentemente para
-- autenticación y verificación de unicidad.
-- -------------------------------------------------------------
CREATE INDEX IF NOT EXISTS idx_users_username ON users (username);
CREATE INDEX IF NOT EXISTS idx_users_email    ON users (email);

-- -------------------------------------------------------------
-- Función y trigger para actualizar updated_at automáticamente
-- -------------------------------------------------------------
CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_users_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION set_updated_at();

-- -------------------------------------------------------------
-- Comentarios de columna para documentación interna
-- -------------------------------------------------------------
COMMENT ON TABLE  users                IS 'Usuarios registrados en la red social';
COMMENT ON COLUMN users.id             IS 'Identificador único autoincremental';
COMMENT ON COLUMN users.username       IS 'Nombre de usuario único, máximo 50 caracteres';
COMMENT ON COLUMN users.email          IS 'Correo electrónico único del usuario';
COMMENT ON COLUMN users.password_hash  IS 'Hash bcrypt de la contraseña';
COMMENT ON COLUMN users.bio            IS 'Descripción breve del perfil del usuario';
COMMENT ON COLUMN users.avatar_url     IS 'URL de la imagen de avatar del usuario';
COMMENT ON COLUMN users.created_at     IS 'Fecha y hora de registro (UTC)';
COMMENT ON COLUMN users.updated_at     IS 'Fecha y hora de última modificación (UTC)';
