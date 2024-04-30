CREATE TABLE edges  (
                        edge_id             SERIAL PRIMARY KEY,
                        source_node_id      INTEGER          NOT NULL,
                        destination_node_id INTEGER          NOT NULL,
                        weight              DOUBLE PRECISION NOT NULL DEFAULT 1
);
