CREATE TABLE IF NOT EXISTS EVENT_TYPES (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  name text NOT NULL,
  description text NOT NULL
);

CREATE TABLE IF NOT EXISTS EVENTS (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  year INTEGER NOT NULL,
  start_date TIMESTAMP(0) WITH TIME ZONE,
  end_date TIMESTAMP(0) WITH TIME ZONE,
  active_event bool NOT NULL DEFAULT false,
  viewers INTEGER NOT NULL DEFAULT 0,
  donations FLOAT NOT NULL DEFAULT 0,
  donors INTEGER NOT NULL DEFAULT 0,
  games_completed INTEGER NOT NULL DEFAULT 0,
  twitch_chats INTEGER NOT NULL DEFAULT 0,
  tweets INTEGER NOT NULL DEFAULT 0,
  schedule_id INTEGER NOT NULL DEFAULT 0,
  event_type_id uuid NOT NULL,
  FOREIGN KEY (event_type_id) REFERENCES EVENT_TYPES (id) ON DELETE CASCADE
);

INSERT INTO
  EVENT_TYPES (id, "name", description)
VALUES
  (
    '82ae1c3d-1a1a-48df-87bb-95733abedd5d',
    'FrostFatales',
    'Frost Fatales'
  );

INSERT INTO
  EVENT_TYPES (id, "name", description)
VALUES
  (
    'aef59e6a-88e2-40dd-9815-96312f0fd8ef',
    'SGDQ',
    'Summer Games Done Quick'
  );

INSERT INTO
  EVENT_TYPES (id, "name", description)
VALUES
  (
    'cce68b8d-024d-4e76-9a51-f199b089c9dc',
    'AGDQ',
    'Awesome Games Done Quick'
  );

INSERT INTO
  EVENTS (
    year,
    start_date,
    end_date,
    active_event,
    schedule_id,
    event_type_id
  )
VALUES
  (
    2016,
    '2016-07-03 19:30:00.000',
    '2016-07-10 12:00:00.000',
    false,
    18,
    'aef59e6a-88e2-40dd-9815-96312f0fd8ef'
  ),
  (
    2017,
    '2017-01-08 17:30:00.000',
    '2017-01-15 19:30:00.000',
    false,
    19,
    'cce68b8d-024d-4e76-9a51-f199b089c9dc'
  ),
  (
    2017,
    '2017-07-02 18:30:00.000',
    '2017-07-09 13:30:00.000',
    false,
    20,
    'aef59e6a-88e2-40dd-9815-96312f0fd8ef'
  ),
  (
    2018,
    '2018-01-07 17:30:00.000',
    '2018-01-14 11:30:00.000',
    false,
    22,
    'cce68b8d-024d-4e76-9a51-f199b089c9dc'
  ),
  (
    2018,
    '2018-06-24 19:30:00.000',
    '2018-07-01 11:30:00.000',
    false,
    23,
    'aef59e6a-88e2-40dd-9815-96312f0fd8ef'
  ),
  (
    2019,
    '2019-01-06 17:30:00.000',
    '2019-01-13 08:00:00.000',
    false,
    25,
    'cce68b8d-024d-4e76-9a51-f199b089c9dc'
  ),
  (
    2019,
    '2019-06-23 19:30:00.000',
    '2019-06-30 10:30:00.000',
    false,
    26,
    'aef59e6a-88e2-40dd-9815-96312f0fd8ef'
  ),
  (
    2020,
    '2020-01-05 17:30:00.000',
    '2020-01-12 09:30:00.000',
    false,
    28,
    'cce68b8d-024d-4e76-9a51-f199b089c9dc'
  ),
  (
    2020,
    '2020-08-16 18:30:00.000',
    '2020-08-23 11:00:00.000',
    false,
    30,
    'aef59e6a-88e2-40dd-9815-96312f0fd8ef'
  ),
  (
    2021,
    '2021-01-03 17:30:00.000',
    '2021-01-10 11:30:00.000',
    false,
    34,
    'cce68b8d-024d-4e76-9a51-f199b089c9dc'
  ),
  (
    2021,
    '2021-07-04 18:30:00.000',
    '2021-07-01 11:30:00.000',
    false,
    35,
    'aef59e6a-88e2-40dd-9815-96312f0fd8ef'
  ),
  (
    2022,
    '2022-01-09 17:30:00.000',
    '2022-01-16 08:30:00.000',
    false,
    37,
    'cce68b8d-024d-4e76-9a51-f199b089c9dc'
  ),
  (
    2022,
    '2022-06-26 19:30:00.000',
    '2022-07-03 11:30:00.000',
    false,
    39,
    'aef59e6a-88e2-40dd-9815-96312f0fd8ef'
  ),
  (
    2023,
    '2023-01-08 17:30:00.000',
    '2023-01-15 10:00:00.000',
    false,
    41,
    'cce68b8d-024d-4e76-9a51-f199b089c9dc'
  ),
  (
    2023,
    '2023-02-26 18:30:00.000',
    '2023-03-05 08:00:00.000',
    TRUE,
    42,
    '82ae1c3d-1a1a-48df-87bb-95733abedd5d'
  );
