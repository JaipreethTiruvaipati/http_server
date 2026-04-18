-- +goose Up
-- Add a non-null text column that forces existing users to fallback to 'unset'
ALTER TABLE users ADD COLUMN hashed_password TEXT NOT NULL DEFAULT 'unset';

-- +goose Down
-- Simple clean up if we ever wanted to revert this
ALTER TABLE users DROP COLUMN hashed_password;
