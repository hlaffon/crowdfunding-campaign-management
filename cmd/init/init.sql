DROP TABLE IF EXISTS contribution;
DROP TABLE IF EXISTS visit;
DROP TABLE IF EXISTS project;

CREATE TABLE IF NOT EXISTS project
(
    project_id INT NOT NULL DEFAULT 0 PRIMARY KEY,
    project_name VARCHAR NOT NULL DEFAULT '',
    slug VARCHAR NOT NULL DEFAULT '',
    currency CHAR(3) NOT NULL DEFAULT '',
    project_type VARCHAR(10) NOT NULL DEFAULT '',
    goal INT NOT NULL DEFAULT 0,
    start_date TIMESTAMP,
    end_date TIMESTAMP WITH TIME ZONE,
    end_date_extra_time TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS contribution
(
    id INT NOT NULL DEFAULT 0 PRIMARY KEY,
    amount INT NOT NULL DEFAULT 0,
    user_id INT NOT NULL DEFAULT 0,
    project_id INT NOT NULL DEFAULT 0,
    created_date TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (project_id) REFERENCES project (project_id)
);

CREATE TABLE IF NOT EXISTS visit
(
    id SERIAL PRIMARY KEY,
    project_id INT NOT NULL DEFAULT 0,
    visit_date TIMESTAMP,
    views INT NOT NULL DEFAULT 0,
    visitors INT NOT NULL DEFAULT 0,
    FOREIGN KEY (project_id) REFERENCES project (project_id)
);