
CREATE TABLE cats (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    years_experience INT NOT NULL CHECK (years_experience >= 0),
    breed TEXT NOT NULL,
    salary BIGINT NOT NULL CHECK (salary >= 0),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);


CREATE TABLE missions (
    id SERIAL PRIMARY KEY,
    cat_id INT REFERENCES cats(id) ON DELETE SET NULL,
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    completed_at TIMESTAMP NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE targets (
    id SERIAL PRIMARY KEY,
    mission_id INT NOT NULL REFERENCES missions(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    country TEXT NOT NULL,
    notes TEXT,
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    completed_at TIMESTAMP NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
