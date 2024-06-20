CREATE TABLE IF NOT EXISTS Message (
    message_id UUID PRIMARY KEY,
    from_user UUID NOT NULL,
    'to' UUID NOT NULL,
    'message' TEXT NOT NULL,
    to_group boolean NOT NULL,

    CONSTRAINT fk_from_user
        FOREIGN KEY (from_user)
        REFERENCES User(user_id)
        ON DELETE CASCADE,

    CREATE INDEX idx_to
        ON Message('to')
        ON DELETE CASCADE

    CREATE INDEX idx_to_group
        ON Message(to_group)
        ON DELETE CASCADE
);