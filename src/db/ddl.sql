CREATE TABLE experiments (
    id          INT UNSIGNED NOT NULL AUTO_INCREMENT,
    name        VARCHAR(128) NOT NULL,
    created_at  DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),

    PRIMARY KEY (id)
);

CREATE TABLE runs (
    id             INT UNSIGNED NOT NULL AUTO_INCREMENT,
    experiment_id  INT UNSIGNED NOT NULL,
    name           VARCHAR(128),
    status         TINYINT UNSIGNED NOT NULL DEFAULT 0,
    created_at     DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    ended_at       DATETIME(6),

    PRIMARY KEY (id),
    FOREIGN KEY (experiment_id)
        REFERENCES experiments(id)
        ON DELETE CASCADE
);

CREATE TABLE parameters (
    id      INT UNSIGNED NOT NULL AUTO_INCREMENT,
    run_id  INT UNSIGNED NOT NULL,
    name    VARCHAR(128) NOT NULL,
    value   TEXT NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (run_id)
        REFERENCES runs(id)
        ON DELETE CASCADE
);

CREATE TABLE metrics (
    id      INT UNSIGNED NOT NULL AUTO_INCREMENT,
    run_id  INT UNSIGNED NOT NULL,
    name    VARCHAR(128) NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (run_id)
        REFERENCES runs(id)
        ON DELETE CASCADE
);

CREATE TABLE plots (
    metric_id    INT UNSIGNED NOT NULL,
    step         BIGINT UNSIGNED NOT NULL,
    value        DOUBLE NOT NULL,

    PRIMARY KEY (metric_id, step) CLUSTERED,
    FOREIGN KEY (metric_id)
        REFERENCES metrics(id)
        ON DELETE CASCADE
);

CREATE TABLE artifacts (
    id            BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    run_id        INT UNSIGNED NOT NULL,
    path          VARCHAR(1024) NOT NULL,
    object_uri    VARCHAR(2048) NOT NULL,
    size_bytes    BIGINT UNSIGNED,
    content_type  VARCHAR(255),
    checksum      BINARY(32),
    created_at    DATETIME(6) NOT NULL,

    PRIMARY KEY (id),
    UNIQUE KEY uq_artifact_path (run_id, path),
    FOREIGN KEY (run_id)
        REFERENCES runs(id)
        ON DELETE CASCADE
);