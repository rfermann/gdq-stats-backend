CREATE TABLE IF NOT EXISTS games
(
    id         UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    start_date TIMESTAMP(0) WITH TIME ZONE NOT NULL,
    end_date   TIMESTAMP(0) WITH TIME ZONE NOT NULL,
    duration   INTERVAL                    NOT NULL,
    name       TEXT                        NOT NULL,
    runners    TEXT                        NOT NULL,
    gdq_id     INT                         NOT NULL,
    event_id   UUID                        NOT NULL,
    FOREIGN KEY (event_id) REFERENCES events (id) ON DELETE CASCADE
);
