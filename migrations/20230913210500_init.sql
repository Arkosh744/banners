-- +goose Up
-- +goose StatementBegin
CREATE TABLE slots
(
    id          serial PRIMARY KEY NOT NULL,
    description text               NOT NULL DEFAULT ''
);

CREATE TABLE banners
(
    id          serial PRIMARY KEY NOT NULL,
    description text               NOT NULL DEFAULT ''
);

CREATE TABLE social_groups
(
    id          serial PRIMARY KEY NOT NULL,
    description text               NOT NULL DEFAULT ''
);

INSERT INTO social_groups (description)
VALUES ('девушки 20-25'),
       ('дедушки 80+'),
       ('собачники 30-40');

CREATE TABLE banner_slot
(
    slot_id   bigint NOT NULL,
    banner_id bigint NOT NULL,
    PRIMARY KEY (slot_id, banner_id)
);

CREATE TABLE views
(
    id        serial PRIMARY KEY NOT NULL,
    slot_id   bigint             NOT NULL,
    banner_id bigint             NOT NULL,
    group_id  bigint             NOT NULL,
    date      timestamp          NOT NULL default NOW()
);

CREATE TABLE clicks
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
