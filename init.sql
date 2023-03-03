CREATE USER darika WITH PASSWORD 'password';

CREATE TABLE IF NOT EXISTS mytracks (
                                         id bigserial CONSTRAINT mytracks_pk PRIMARY KEY,
                                         created_at TIMESTAMP,
                                         name VARCHAR,
                                         duration INTEGER
 );

