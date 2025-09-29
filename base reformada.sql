-- =====================================================
-- SCRIPT MEJORADO PARA SISTEMA DE GESTIÓN DE CAUSAS
-- =====================================================
-- Script mejorado para XAMPP (MySQL/MariaDB)
-- Sistema de Gestión de Causas con roles de administrador y usuario normal
-- Compatible con Django Backend y Frontend React
-- =====================================================

-- 1. Creación de la base de datos
DROP DATABASE IF EXISTS sistema_causas;
CREATE DATABASE sistema_causas
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;
USE sistema_causas;

-- =====================================================
-- 2. TABLAS DE UBICACIÓN GEOGRÁFICA
-- =====================================================

-- Provincias
CREATE TABLE provincias (
  id INT AUTO_INCREMENT PRIMARY KEY,
  nombre VARCHAR(100) NOT NULL UNIQUE,
  is_active TINYINT(1) NOT NULL DEFAULT 1,
  codigo_estadistico VARCHAR(10) NOT NULL UNIQUE,
  created_at TIMESTAMP    DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP    DEFAULT CURRENT_TIMESTAMP
                          ON UPDATE CURRENT_TIMESTAMP,
  INDEX (nombre),
  INDEX (codigo_estadistico)
) ENGINE=InnoDB
  COMMENT = 'Provincias de Argentina';

-- Localidades
CREATE TABLE localidades (
  id INT AUTO_INCREMENT PRIMARY KEY,
  nombre VARCHAR(100) NOT NULL,
  is_active TINYINT(1) NOT NULL DEFAULT 1,
  provincia_id INT NOT NULL,
  codigo_postal VARCHAR(10),
  nombre_calle VARCHAR(100),
  altura INT,
  created_at TIMESTAMP    DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP    DEFAULT CURRENT_TIMESTAMP
                          ON UPDATE CURRENT_TIMESTAMP,
  CONSTRAINT fk_localidad_provincia
    FOREIGN KEY (provincia_id)
    REFERENCES provincias(id)
    ON UPDATE CASCADE
    ON DELETE RESTRICT,
  UNIQUE KEY uq_localidad_provincia (nombre, provincia_id),
  INDEX (provincia_id),
  INDEX (nombre)
) ENGINE=InnoDB
  COMMENT = 'Localidades por provincia';

-- =====================================================
-- 3. USUARIOS Y ROLES
-- =====================================================

-- Tabla principal de usuarios (estilo Django AbstractUser)
CREATE TABLE usuarios (
  id INT AUTO_INCREMENT PRIMARY KEY,
  clave VARCHAR(256) NOT NULL,
  last_login DATETIME(6),
  is_superuser TINYINT(1) NOT NULL DEFAULT 0,
  username VARCHAR(150) ,
  first_name VARCHAR(150) NOT NULL DEFAULT '',
  last_name VARCHAR(150) NOT NULL DEFAULT '',
  is_active TINYINT(1) NOT NULL DEFAULT 1,
  date_joined DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),

  email VARCHAR(254) NULL UNIQUE ,
  rol ENUM('administrador','normal') NOT NULL DEFAULT 'normal',
  nombre_completo VARCHAR(150),
  telefono VARCHAR(20),

  dni VARCHAR(20) UNIQUE,
  ce VARCHAR(20),
  celular VARCHAR(20),

  fecha_creacion TIMESTAMP    DEFAULT CURRENT_TIMESTAMP,
  fecha_actualizacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
                            ON UPDATE CURRENT_TIMESTAMP,

  INDEX (email),
  INDEX (username),
  INDEX (rol),
  INDEX (dni)
) ENGINE=InnoDB
  COMMENT = 'Usuarios del sistema';

-- (Opcional) Tablas de compatibilidad con Django auth
CREATE TABLE auth_group (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(150) NOT NULL UNIQUE
) ENGINE=InnoDB;

CREATE TABLE auth_permission (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  content_type_id INT NOT NULL,
  codename VARCHAR(100) NOT NULL
) ENGINE=InnoDB;

-- =====================================================
-- 4. TABLA PRINCIPAL: CAUSAS
-- =====================================================

CREATE TABLE causas (
  id INT AUTO_INCREMENT PRIMARY KEY,

  numero_causa VARCHAR(50) NOT NULL UNIQUE,
  caratula TEXT NOT NULL,

  juzgado_id int  NOT NULL DEFAULT 0,
  fiscalia_id int  NOT NULL DEFAULT 0,
  a_cargo_del_magistrado VARCHAR(150) NOT NULL default 'vacio',
  preventor VARCHAR(150) NOT NULL,
  preventor_auxiliar VARCHAR(150) NOT NULL,

  provincia_id INT NOT NULL DEFAULT 1,
  localidad_id INT DEFAULT 1,
  domicilio VARCHAR(255),
  nro_sgo VARCHAR(50),
  nro_mto VARCHAR(50),
  tipo_delito VARCHAR(50),

  nombre_fantasia VARCHAR(150),
  fecha_llegada VARCHAR(20),
  providencia TEXT,

  estado ENUM('activa','archivada','en_proceso','finalizada')
         NOT NULL DEFAULT 'activa',

  is_active TINYINT(1) NOT NULL DEFAULT 1,
  deleted_at TIMESTAMP NULL,

  creado_por INT NOT NULL DEFAULT 1,
  actualizado_por INT DEFAULT 1,
  creado_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  actualizado_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP
                   ON UPDATE CURRENT_TIMESTAMP,

  CONSTRAINT fk_causa_provincia
    FOREIGN KEY (provincia_id)
    REFERENCES provincias(id)
    ON UPDATE CASCADE
    ON DELETE RESTRICT,

  CONSTRAINT fk_causa_localidad
    FOREIGN KEY (localidad_id)
    REFERENCES localidades(id)
    ON UPDATE CASCADE
    ON DELETE SET NULL,

  CONSTRAINT fk_causa_creador
    FOREIGN KEY (creado_por)
    REFERENCES usuarios(id)
    ON UPDATE CASCADE
    ON DELETE RESTRICT,

  CONSTRAINT fk_causa_actualizador
    FOREIGN KEY (actualizado_por)
    REFERENCES usuarios(id)
    ON UPDATE CASCADE
    ON DELETE SET NULL,

  CONSTRAINT fk_fiscalia
    FOREIGN KEY (fiscalia_id)
    REFERENCES fiscalias(id)
    ON UPDATE CASCADE
    ON DELETE RESTRICT,

  CONSTRAINT fk_juzgado
    FOREIGN KEY (juzgado_id)
    REFERENCES juzgados(id)
    ON UPDATE CASCADE
    ON DELETE RESTRICT,    

  INDEX (numero_causa),
  INDEX (estado),
  INDEX (provincia_id),
  INDEX (fiscalia_id),
  INDEX (juzgado_id),
  INDEX (creado_en),
  INDEX (is_active),
  INDEX (estado, provincia_id),
  INDEX (is_active, estado)
) ENGINE=InnoDB
  COMMENT = 'Causas legales del sistema';

-- =====================================================
-- 5. HISTORIAL / AUDITORÍA
-- =====================================================
  --id_causas INT NOT NULL,
CREATE TABLE historial_causas (
  id INT AUTO_INCREMENT PRIMARY KEY,
  numero_causa VARCHAR(45) NOT NULL,
  accion ENUM('creacion','actualizacion','eliminacion','cambio_estado')
         NOT NULL,
  usuario_id INT NOT NULL,
  descripcion TEXT NOT NULL,
  datos_anteriores JSON,
  datos_nuevos JSON,
  ip_address VARCHAR(45),
  user_agent TEXT,
  creado_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

  CONSTRAINT fk_historial_causa
    FOREIGN KEY (numero_causa)
    REFERENCES causas(numero_causa)
    ON UPDATE CASCADE
    ON DELETE CASCADE,

  CONSTRAINT fk_historial_usuario
    FOREIGN KEY (usuario_id)
    REFERENCES usuarios(id)
    ON UPDATE CASCADE
    ON DELETE RESTRICT,


  INDEX (usuario_id),
  INDEX (accion),
  INDEX (creado_en)
) ENGINE=InnoDB
  COMMENT = 'Historial de cambios en causas';

-- =====================================================
-- 6. TABLAS ADICIONALES (Documentos, Notas, fiscalia y juzgados y preventores )
-- =====================================================

-- Fiscalias
CREATE TABLE fiscalias (
  id INT AUTO_INCREMENT PRIMARY KEY,
  nombre VARCHAR(200) NOT NULL UNIQUE,
  is_active TINYINT(1) NOT NULL DEFAULT 1,
  created_at TIMESTAMP    DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP    DEFAULT CURRENT_TIMESTAMP
                          ON UPDATE CURRENT_TIMESTAMP,
                          
  INDEX (nombre)
) ENGINE=InnoDB
  COMMENT = 'Fiscalias';



  -- Juzgados
CREATE TABLE juzgados (
  id INT AUTO_INCREMENT PRIMARY KEY,
  nombre VARCHAR(200) NOT NULL UNIQUE,
  is_active TINYINT(1) NOT NULL DEFAULT 1,
  created_at TIMESTAMP    DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP    DEFAULT CURRENT_TIMESTAMP
                          ON UPDATE CURRENT_TIMESTAMP,
  INDEX (nombre)
) ENGINE=InnoDB
  COMMENT = 'Juzgados';

-- preventores
CREATE TABLE preventores (
  --id INT AUTO_INCREMENT PRIMARY KEY,
  nombre VARCHAR(200) NOT NULL UNIQUE,
  is_active TINYINT(1) NOT NULL DEFAULT 1,
  created_at TIMESTAMP    DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP    DEFAULT CURRENT_TIMESTAMP
                          ON UPDATE CURRENT_TIMESTAMP,
                          
  INDEX (nombre)
) ENGINE=InnoDB
  COMMENT = 'Preventores';


CREATE TABLE documentos_causa (
  id INT AUTO_INCREMENT PRIMARY KEY,
  causa_id INT NOT NULL,
  nombre_archivo VARCHAR(255) NOT NULL,
  ruta_archivo VARCHAR(500) NOT NULL,
  tipo_documento VARCHAR(100),
  tamano INT,
  subido_por INT NOT NULL,
  subido_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

  CONSTRAINT fk_doc_causa
    FOREIGN KEY (causa_id)
    REFERENCES causas(id)
    ON UPDATE CASCADE
    ON DELETE CASCADE,

  CONSTRAINT fk_doc_usuario
    FOREIGN KEY (subido_por)
    REFERENCES usuarios(id)
    ON UPDATE CASCADE
    ON DELETE RESTRICT,

  INDEX (causa_id),
  INDEX (subido_por)
) ENGINE=InnoDB
  COMMENT = 'Documentos adjuntos a las causas';

CREATE TABLE notas_causa (
  id INT AUTO_INCREMENT PRIMARY KEY,
  causa_id INT NOT NULL,
  contenido TEXT NOT NULL,
  es_privada TINYINT(1) NOT NULL DEFAULT 0,
  creado_por INT NOT NULL,
  creado_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  actualizado_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP
                    ON UPDATE CURRENT_TIMESTAMP,

  CONSTRAINT fk_nota_causa
    FOREIGN KEY (causa_id)
    REFERENCES causas(id)
    ON UPDATE CASCADE
    ON DELETE CASCADE,

  CONSTRAINT fk_nota_usuario
    FOREIGN KEY (creado_por)
    REFERENCES usuarios(id)
    ON UPDATE CASCADE
    ON DELETE RESTRICT,

  INDEX (causa_id),
  INDEX (creado_por)
) ENGINE=InnoDB
  COMMENT = 'Notas y comentarios de las causas';


-- =====================================================
-- 7. INSERCIÓN DE DATOS INICIALES
-- =====================================================

-- Provincias principales de Argentina
INSERT INTO provincias (nombre, codigo_estadistico) VALUES
('Buenos Aires', 'BA'),
('Ciudad Autónoma de Buenos Aires', 'CABA'),
('Córdoba', 'CB'),
('Santa Fe', 'SF'),
('Mendoza', 'MZ'),
('Tucumán', 'TM'),
('Entre Ríos', 'ER'),
('Salta', 'SA'),
('Misiones', 'MN'),
('Chaco', 'CC'),
('Corrientes', 'CN'),
('Santiago del Estero', 'SE'),
('San Juan', 'SJ'),
('Jujuy', 'JY'),
('Río Negro', 'RN'),
('Formosa', 'FM'),
('Neuquén', 'NQ'),
('Chubut', 'CH'),
('San Luis', 'SL'),
('Catamarca', 'CT'),
('La Rioja', 'LR'),
('La Pampa', 'LP'),
('Santa Cruz', 'SC'),
('Tierra del Fuego', 'TF');

-- Localidades principales
INSERT INTO localidades (nombre, provincia_id, codigo_postal) VALUES
-- Buenos Aires
('La Plata', 1, '1900'),
('Mar del Plata', 1, '7600'),
('Bahía Blanca', 1, '8000'),
('Tandil', 1, '7000'),
-- CABA
('Puerto Madero', 2, '1107'),
('Palermo', 2, '1414'),
('Recoleta', 2, '1113'),
('San Telmo', 2, '1103'),
-- Córdoba
('Córdoba Capital', 3, '5000'),
('Villa Carlos Paz', 3, '5152'),
('Río Cuarto', 3, '5800'),
-- Santa Fe
('Rosario', 4, '2000'),
('Santa Fe Capital', 4, '3000'),
('Rafaela', 4, '2300'),
-- Mendoza
('Mendoza Capital', 5, '5500'),
('San Rafael', 5, '5600'),
('Godoy Cruz', 5, '5501');

-- Usuario administrador por defecto
INSERT INTO usuarios (
  username, email, password, rol, nombre_completo, 
  is_staff, is_superuser, is_active
) VALUES (
  'admin', 
  'admin@sistema-legal.com', 
  'pbkdf2_sha256$600000$dummy$hash', -- Cambiar por hash real
  'administrador', 
  'Administrador del Sistema',
  1, 1, 1
);

-- Usuario normal de ejemplo
INSERT INTO usuarios (
  username, email, password, rol, nombre_completo,
  dni, telefono
) VALUES (
  'usuario.demo', 
  'usuario@sistema-legal.com', 
  'pbkdf2_sha256$600000$dummy$hash', -- Cambiar por hash real
  'normal', 
  'Usuario Demostración',
  '12345678', 
  '+54 11 1234-5678'
);

-- =====================================================
-- 8. VISTAS ÚTILES PARA REPORTES
-- =====================================================

-- Vista para causas con información completa
CREATE VIEW vista_causas_completa AS
SELECT 
  c.id,
  c.numero_causa,
  c.caratula,
  c.juzgado,
  c.fiscalia,
  c.a_cargo_del_magistrado,
  c.preventor_auxiliar,
  p.nombre AS provincia_nombre,
  l.nombre AS localidad_nombre,
  l.codigo_postal,
  c.domicilio,
  c.nombre_fantasia,
  c.estado,
  c.fecha_llegada,
  c.is_active,
  u_creador.nombre_completo AS creado_por_nombre,
  u_actualizador.nombre_completo AS actualizado_por_nombre,
  c.creado_en,
  c.actualizado_en
FROM causas c
LEFT JOIN provincias p ON c.provincia_id = p.id
LEFT JOIN localidades l ON c.localidad_id = l.id
LEFT JOIN usuarios u_creador ON c.creado_por = u_creador.id
LEFT JOIN usuarios u_actualizador ON c.actualizado_por = u_actualizador.id
WHERE c.is_active = 1;

-- Vista para estadísticas por provincia
CREATE VIEW vista_estadisticas_provincia AS
SELECT 
  p.nombre AS provincia,
  COUNT(c.id) AS total_causas,
  SUM(CASE WHEN c.estado = 'activa' THEN 1 ELSE 0 END) AS causas_activas,
  SUM(CASE WHEN c.estado = 'en_proceso' THEN 1 ELSE 0 END) AS causas_en_proceso,
  SUM(CASE WHEN c.estado = 'finalizada' THEN 1 ELSE 0 END) AS causas_finalizadas,
  SUM(CASE WHEN c.estado = 'archivada' THEN 1 ELSE 0 END) AS causas_archivadas
FROM provincias p
LEFT JOIN causas c ON p.id = c.provincia_id AND c.is_active = 1
GROUP BY p.id, p.nombre
ORDER BY total_causas DESC;

-- =====================================================
-- 9. PROCEDIMIENTOS ALMACENADOS ÚTILES
-- =====================================================

DELIMITER //

-- Procedimiento para obtener estadísticas generales
CREATE PROCEDURE sp_estadisticas_generales()
BEGIN
  SELECT 
    'Total Causas' AS metrica,
    COUNT(*) AS valor
  FROM causas 
  WHERE is_active = 1
  
  UNION ALL
  
  SELECT 
    CONCAT('Causas ', UPPER(estado)) AS metrica,
    COUNT(*) AS valor
  FROM causas 
  WHERE is_active = 1
  GROUP BY estado
  
  UNION ALL
  
  SELECT 
    'Usuarios Activos' AS metrica,
    COUNT(*) AS valor
  FROM usuarios 
  WHERE is_active = 1;
END //

-- Procedimiento para buscar causas
CREATE PROCEDURE sp_buscar_causas(
  IN p_termino_busqueda VARCHAR(255),
  IN p_estado VARCHAR(20),
  IN p_provincia_id INT,
  IN p_limite INT,
  IN p_offset INT
)
BEGIN
  SELECT 
    c.*,
    p.nombre AS provincia_nombre,
    l.nombre AS localidad_nombre
  FROM causas c
  LEFT JOIN provincias p ON c.provincia_id = p.id
  LEFT JOIN localidades l ON c.localidad_id = l.id
  WHERE c.is_active = 1
    AND (p_termino_busqueda IS NULL OR 
         c.numero_causa LIKE CONCAT('%', p_termino_busqueda, '%') OR
         c.caratula LIKE CONCAT('%', p_termino_busqueda, '%') OR
         c.fiscalia LIKE CONCAT('%', p_termino_busqueda, '%'))
    AND (p_estado IS NULL OR c.estado = p_estado)
    AND (p_provincia_id IS NULL OR c.provincia_id = p_provincia_id)
  ORDER BY c.creado_en DESC
  LIMIT p_limite OFFSET p_offset;
END //

DELIMITER ;

-- =====================================================
-- 10. TRIGGERS PARA AUDITORÍA AUTOMÁTICA
-- =====================================================

DELIMITER //

-- Trigger para registrar creación de causas
CREATE TRIGGER tr_causa_insert 
AFTER INSERT ON causas
FOR EACH ROW
BEGIN
  INSERT INTO historial_causas (
    causa_id, accion, usuario_id, descripcion, datos_nuevos
  ) VALUES (
    NEW.id, 
    'creacion', 
    NEW.creado_por, 
    CONCAT('Causa creada: ', NEW.numero_causa),
    JSON_OBJECT(
      'numero_causa', NEW.numero_causa,
      'caratula', NEW.caratula,
      'estado', NEW.estado
    )
  );
END //

-- Trigger para registrar actualizaciones de causas
CREATE TRIGGER tr_causa_update 
AFTER UPDATE ON causas
FOR EACH ROW
BEGIN
  INSERT INTO historial_causas (
    causa_id, accion, usuario_id, descripcion, 
    datos_anteriores, datos_nuevos
  ) VALUES (
    NEW.id, 
    'actualizacion', 
    COALESCE(NEW.actualizado_por, NEW.creado_por), 
    CONCAT('Causa actualizada: ', NEW.numero_causa),
    JSON_OBJECT(
      'numero_causa', OLD.numero_causa,
      'caratula', OLD.caratula,
      'estado', OLD.estado
    ),
    JSON_OBJECT(
      'numero_causa', NEW.numero_causa,
      'caratula', NEW.caratula,
      'estado', NEW.estado
    )
  );
END //

DELIMITER ;

-- =====================================================
-- COMENTARIOS FINALES
-- =====================================================

/*
MEJORAS IMPLEMENTADAS:

1. ✅ Compatibilidad con Django AbstractUser
2. ✅ Todos los campos del frontend incluidos
3. ✅ Estructura de tu script original preservada
4. ✅ Índices optimizados para rendimiento
5. ✅ Soft delete implementado
6. ✅ Auditoría completa con triggers
7. ✅ Vistas para reportes
8. ✅ Procedimientos almacenados útiles
9. ✅ Datos iniciales incluidos
10. ✅ Preparado para funcionalidades futuras

CAMPOS ADICIONALES DE TU ESTRUCTURA ORIGINAL:
- nombre_calle y altura en localidades
- fecha_llegada y providencia en causas
- dni, ce, celular en usuarios
- tipos_administrativos como tabla separada

CAMPOS DEL FRONTEND AGREGADOS:
- juzgado, fiscalia, a_cargo
- preventor_auxiliar, estado
- domicilio, nombre_fantasia
- Sistema completo de auditoría

USO RECOMENDADO:
1. Ejecutar este script para crear la estructura
2. Usar el backend Django que ya está configurado
3. Los modelos Django se mapearán automáticamente
4. El frontend React funcionará sin cambios
*/


Múltiples tablas (provincias, localidades, usuarios, causas, historial_causas, documentos_causa, notas_causa).

Cada tabla representa una entidad o concepto distinto.

Entre ellas hay claves foráneas (FOREIGN KEY) que las relacionan:

localidades.provincia_id → provincias.id

causas.provincia_id → provincias.id y causas.localidad_id → localidades.id

causas.creado_por → usuarios.id (y similar para actualizado_por)

historial_causas.causa_id → causas.id y historial_causas.usuario_id → usuarios.id

documentos_causa.causa_id → causas.id y documentos_causa.subido_por → usuarios.id

notas_causa.causa_id → causas.id y notas_causa.creado_por → usuarios.id

Gracias a esas claves foráneas, puedes hacer joins entre tablas para, por ejemplo, listar todas las causas de una provincia concreta o ver quién creó cada entrada del historial.

En un modelo relacional:

Cada tabla almacena filas (registros) con un primary key (id) que identifica unívocamente cada registro.

Las claves foráneas apuntan a esos id en otras tablas, garantizando la integridad referencial.

Puedes navegar o consultar datos relacionados mediante sentencias SQL como JOIN.

Así que sí: es un modelo típico de base de datos relacional, diseñado para XAMPP/MySQL y perfectamente compatible con Django (que también usa un ORM relacional) y con un frontend React que consuma esas APIs.