CREATE TABLE experiments (
    experiment_id          INT UNSIGNED NOT NULL AUTO_INCREMENT,
    name        VARCHAR(128) NOT NULL,
    description    TEXT,
    created_at  DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    updated_at  DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),

    PRIMARY KEY (id)
);

CREATE TABLE runs (
    run_id             INT UNSIGNED NOT NULL AUTO_INCREMENT,
    experiment_id  INT UNSIGNED NOT NULL,
    name           VARCHAR(128),
    description    TEXT,
    status         TINYINT UNSIGNED NOT NULL DEFAULT 0,
    created_at     DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    ended_at       DATETIME(6),

    PRIMARY KEY (run_id),
    FOREIGN KEY (experiment_id)
        REFERENCES experiments(experiment_id)
        ON DELETE CASCADE
);

CREATE TABLE parameters (
    parameter_id      INT UNSIGNED NOT NULL AUTO_INCREMENT,
    run_id  INT UNSIGNED NOT NULL,
    name    VARCHAR(128) NOT NULL,
    value   TEXT NOT NULL,

    PRIMARY KEY (parameter_id),
    FOREIGN KEY (run_id)
        REFERENCES runs(run_id)
        ON DELETE CASCADE
);

CREATE TABLE metrics (
    metric_id      INT UNSIGNED NOT NULL AUTO_INCREMENT,
    run_id  INT UNSIGNED NOT NULL,
    name    VARCHAR(128) NOT NULL,

    PRIMARY KEY (metric_id),
    FOREIGN KEY (run_id)
        REFERENCES runs(run_id)
        ON DELETE CASCADE
);

CREATE TABLE plots (
    metric_id    INT UNSIGNED NOT NULL,
    step         BIGINT UNSIGNED NOT NULL,
    value        DOUBLE NOT NULL,

    PRIMARY KEY (metric_id, step) CLUSTERED,
    FOREIGN KEY (metric_id)
        REFERENCES metrics(metric_id)
        ON DELETE CASCADE
);