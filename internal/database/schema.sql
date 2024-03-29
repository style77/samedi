CREATE TABLE IF NOT EXISTS blogs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,

    name TEXT NOT NULL UNIQUE,
    description TEXT NULL,
    title TEXT NOT NULL,
    author TEXT NOT NULL,
    logo TEXT NULL,
    github TEXT NULL,
    twitter TEXT NULL,
    linkedin TEXT NULL,
    language TEXT NOT NULL DEFAULT 'en'
);

CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL,

    title TEXT NOT NULL,
    content TEXT NOT NULL,
    blog_id INTEGER NOT NULL,
    FOREIGN KEY(blog_id) REFERENCES blogs(id)
);

CREATE INDEX IF NOT EXISTS idx_posts_created_at ON posts(created_at);
CREATE INDEX IF NOT EXISTS idx_posts_updated_at ON posts(updated_at);
CREATE INDEX IF NOT EXISTS idx_posts_title ON posts(title);
CREATE INDEX IF NOT EXISTS idx_posts_blog_id ON posts(blog_id);