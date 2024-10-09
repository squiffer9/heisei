CREATE TABLE threads (
    id SERIAL PRIMARY KEY,
    category_id INTEGER NOT NULL,
    title VARCHAR(200) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_post_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    post_count INTEGER NOT NULL DEFAULT 0,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);

CREATE INDEX idx_threads_category_id ON threads(category_id);
CREATE INDEX idx_threads_last_post_at ON threads(last_post_at);
