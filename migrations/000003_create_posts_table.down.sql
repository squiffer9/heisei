CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    thread_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    author_ip INET NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    FOREIGN KEY (thread_id) REFERENCES threads(id) ON DELETE CASCADE
);

CREATE INDEX idx_posts_thread_id ON posts(thread_id);
CREATE INDEX idx_posts_created_at ON posts(created_at);
