-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS slots
(
    id          serial PRIMARY KEY NOT NULL,
    description text               NOT NULL DEFAULT ''
);

CREATE TABLE IF NOT EXISTS banners
(
    id          serial PRIMARY KEY NOT NULL,
    description text               NOT NULL DEFAULT ''
);

CREATE TABLE IF NOT EXISTS social_groups
(
    id          serial PRIMARY KEY NOT NULL,
    description text               NOT NULL DEFAULT ''
);

INSERT INTO social_groups (description)
VALUES ('девушки 20-25'),
       ('дедушки 80+'),
       ('пожилые 30-40');

CREATE TABLE IF NOT EXISTS banner_slot
(
    slot_id   bigint NOT NULL,
    banner_id bigint NOT NULL,
    PRIMARY KEY (slot_id, banner_id)
);

CREATE TABLE IF NOT EXISTS views
(
    id        serial PRIMARY KEY NOT NULL,
    slot_id   bigint             NOT NULL,
    banner_id bigint             NOT NULL,
    group_id  bigint             NOT NULL,
    date      timestamp          NOT NULL default NOW()
);

CREATE TABLE IF NOT EXISTS clicks
(
    id        serial PRIMARY KEY NOT NULL,
    slot_id   bigint             NOT NULL,
    banner_id bigint             NOT NULL,
    group_id  bigint             NOT NULL,
    date      timestamp          NOT NULL default NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS clicks;
DROP TABLE IF EXISTS views;
DROP TABLE IF EXISTS banners;
DROP TABLE IF EXISTS banner_slot;
DROP TABLE IF EXISTS social_groups;
DROP TABLE IF EXISTS slots;
-- +goose StatementEnd
