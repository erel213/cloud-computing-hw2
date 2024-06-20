CREATE TABLE IF NOT EXISTS messages (
    message_id UUID PRIMARY KEY,
    from_user UUID NOT NULL,
    send_to UUID NOT NULL,
    message_content TEXT NOT NULL,
    to_group boolean NOT NULL,

    CONSTRAINT fk_from_user
        FOREIGN KEY (from_user)
        REFERENCES users(user_id)
        ON DELETE CASCADE
);

CREATE INDEX idx_to
    ON messages(send_to);

CREATE INDEX idx_to_group
    ON messages(to_group);