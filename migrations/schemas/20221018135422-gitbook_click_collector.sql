
-- +migrate Up
CREATE TABLE IF NOT EXISTS gitbook_click_collectors (
    id SERIAL NOT NULL PRIMARY KEY,
    command TEXT NOT NULL,
    action TEXT,
    count_clicks INT NOT NULL, 
    created_at timestamptz DEFAULT now()
);
-- +migrate Down
DROP TABLE IF EXISTS gitbook_click_collectors;