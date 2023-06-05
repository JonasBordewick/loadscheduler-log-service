CREATE TABLE IF NOT EXISTS logs (
    id SERIAL PRIMARY KEY,
    applicant TEXT,
    log_time TIMESTAMPTZ,
    log_message TEXT
)