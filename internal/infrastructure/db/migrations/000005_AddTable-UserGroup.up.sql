CREATE TABLE IF NOT EXISTS user_group(
    user_group_id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    group_id UUID NOT NULL,

    CONSTRAINT fk_user
        FOREIGN KEY  (user_id)
        REFERENCES users(user_id)
        ON DELETE CASCADE,
    
    CONSTRAINT fk_group
    FOREIGN KEY (group_id)
    REFERENCES groups(group_id)
    ON DELETE CASCADE
)