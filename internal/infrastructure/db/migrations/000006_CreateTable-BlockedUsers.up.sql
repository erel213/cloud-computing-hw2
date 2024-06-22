CREATE TABLE IF NOT EXISTS blocked_users (
    id UUID PRIMARY KEY,
    user_block UUID NOT NULL,
    user_blocked UUID NOT NULL,

    CONSTRAINT fk_user_block
        FOREIGN KEY (user_block)
        REFERENCES users (user_id)
        ON DELETE CASCADE,

    CONSTRAINT fk_user_blocked
        FOREIGN KEY (user_blocked)
        REFERENCES users (user_id)
        ON DELETE CASCADE
);