-- +goose Up
CREATE TABLE IF NOT EXISTS `cast` (
    `id` INT PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS `movie` (
    `id` INT PRIMARY KEY,
    `title` VARCHAR(255) NOT NULL,
    `release_year` INT NOT NULL,
    `quality` VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS `moviecast` (
    `id` INT PRIMARY KEY,
    `movie_id` INT NOT NULL,
    `cast_id` INT NOT NULL,
    FOREIGN KEY (`movie_id`) REFERENCES `movie`(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`cast_id`) REFERENCES `cast`(`id`) ON DELETE CASCADE
);

DELETE FROM
    `moviecast`;

DELETE FROM
    `movie`;

DELETE FROM
    `cast`;

-- Casts (Actors)
INSERT INTO
    `cast` (`id`, `name`)
VALUES
    (1, 'Robert Downey Jr'),
    (2, 'Scarlett Johansson'),
    (3, 'Chris Evans'),
    (4, 'Tom Holland'),
    (5, 'Leonardo DiCaprio'),
    (6, 'Brad Pitt'),
    (7, 'Morgan Freeman'),
    (8, 'Natalie Portman'),
    (9, 'Emma Stone'),
    (10, 'Jennifer Lawrence'),
    (11, 'Keanu Reeves'),
    (12, 'Christian Bale'),
    (13, 'Anne Hathaway'),
    (14, 'Ryan Gosling'),
    (15, 'Margot Robbie'),
    (16, 'Tom Hardy'),
    (17, 'Dwayne Johnson'),
    (18, 'Will Smith'),
    (19, 'Matt Damon'),
    (20, 'Angelina Jolie'),
    (21, 'Hugh Jackman'),
    (22, 'Samuel L Jackson'),
    (23, 'Gal Gadot'),
    (24, 'Jason Momoa'),
    (25, 'Chris Hemsworth'),
    (26, 'Mark Ruffalo'),
    (27, 'Ben Affleck'),
    (28, 'Zendaya'),
    (29, 'Millie Bobby Brown'),
    (30, 'Timothee Chalamet');

-- Movies
INSERT INTO
    `movie` (`id`, `title`, `release_year`, `quality`)
VALUES
    (1, 'The Last Mission', 2015, '720p'),
    (2, 'Shadow War', 2016, '1080p'),
    (3, 'Galaxy Strike', 2017, '4K'),
    (4, 'Hidden Truth', 2018, '1080p'),
    (5, 'Lost Empire', 2019, '4K'),
    (6, 'Silent Storm', 2020, '720p'),
    (7, 'Night Hunter', 2021, '1080p'),
    (8, 'Broken City', 2022, '4K'),
    (9, 'Final Horizon', 2023, '4K'),
    (10, 'Crimson Code', 2024, '1080p'),
    (11, 'Dark Protocol', 2014, '720p'),
    (12, 'Iron Justice', 2013, '1080p'),
    (13, 'Neon Future', 2022, '4K'),
    (14, 'Quantum Escape', 2021, '1080p'),
    (15, 'Golden Empire', 2019, '720p'),
    (16, 'Ocean of Fire', 2018, '4K'),
    (17, 'Frozen Silence', 2020, '1080p'),
    (18, 'Deep Space', 2023, '4K'),
    (19, 'Shadow City', 2017, '720p'),
    (20, 'Eternal Night', 2024, '4K');

-- MovieCast Relations
INSERT INTO
    `moviecast` (`id`, `movie_id`, `cast_id`)
VALUES
    (1, 1, 1),
    (2, 1, 2),
    (3, 1, 3),
    (4, 2, 4),
    (5, 2, 5),
    (6, 2, 6),
    (7, 3, 7),
    (8, 3, 8),
    (9, 3, 9),
    (10, 4, 10),
    (11, 4, 11),
    (12, 4, 12),
    (13, 5, 13),
    (14, 5, 14),
    (15, 5, 15),
    (16, 6, 16),
    (17, 6, 17),
    (18, 6, 18),
    (19, 7, 19),
    (20, 7, 20),
    (21, 7, 21),
    (22, 8, 22),
    (23, 8, 23),
    (24, 8, 24),
    (25, 9, 25),
    (26, 9, 26),
    (27, 9, 27),
    (28, 10, 28),
    (29, 10, 29),
    (30, 10, 30),
    (31, 11, 1),
    (32, 11, 5),
    (33, 12, 6),
    (34, 12, 7),
    (35, 13, 8),
    (36, 13, 9),
    (37, 14, 10),
    (38, 14, 11),
    (39, 15, 12),
    (40, 15, 13),
    (41, 16, 14),
    (42, 16, 15),
    (43, 17, 16),
    (44, 17, 17),
    (45, 18, 18),
    (46, 18, 19),
    (47, 19, 20),
    (48, 19, 21),
    (49, 20, 22),
    (50, 20, 23);

-- +goose Down
DROP TABLE IF EXISTS `moviecast`;

DROP TABLE IF EXISTS `movie`;

DROP TABLE IF EXISTS `cast`;