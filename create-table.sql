CREATE TABLE
    `records` (
        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
        `name` VARBINARY(255) NOT NULL,
        `type` SMALLINT UNSIGNED NOT NULL,
        `value` VARBINARY(512) NOT NULL,
        `ttl` INT UNSIGNED NOT NULL,
        PRIMARY KEY (`id`) USING BTREE
    ) COLLATE = 'utf8mb4_0900_ai_ci';