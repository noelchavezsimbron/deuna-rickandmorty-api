CREATE TABLE episode
(
    id         serial,
    name       varchar(100) NOT NULL,
    air_date   varchar(50),
    episode    varchar(50),
    characters text,
    url        varchar(200),
    created    timestamp
);