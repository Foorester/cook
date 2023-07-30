--UP
CREATE TABLE ingredients (
                             id TEXT PRIMARY KEY,
                             name TEXT NOT NULL,
                             description TEXT NOT NULL,
                             recipe_id TEXT NOT NULL,
                             qty TEXT NOT NULL,
                             unit TEXT NOT NULL,
                             created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                             updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

--DOWN
DROP TABLE ingredients;
