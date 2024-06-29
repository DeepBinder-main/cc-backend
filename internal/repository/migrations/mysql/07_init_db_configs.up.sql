CREATE TABLE IF NOT EXISTS file_stash_url (
    id INT AUTO_INCREMENT PRIMARY KEY,
    url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    single_row_enforcer INT NOT NULL DEFAULT 1,
    UNIQUE KEY single_row_enforcer_unique (single_row_enforcer)
);

CREATE TABLE IF NOT EXISTS rabbit_mq_config (
    conn_url VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    single_row_enforcer INT NOT NULL DEFAULT 1,
    UNIQUE KEY single_row_enforcer_unique (single_row_enforcer)
);