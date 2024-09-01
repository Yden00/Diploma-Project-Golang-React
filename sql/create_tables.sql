CREATE TABLE files (
    id SERIAL PRIMARY KEY,
    filename TEXT NOT NULL,
    filetype TEXT,
    filesize BIGINT,
    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
