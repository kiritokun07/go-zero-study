CREATE TABLE `shorturl`
(
    `shorten` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'shorten key',
    `url`     varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'original url',
    PRIMARY KEY (`shorten`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='shorturl';