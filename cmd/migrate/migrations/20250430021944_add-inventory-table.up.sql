CREATE TABLE IF NOT EXISTS inventory (
    `product_id` INT UNSIGNED NOT NULL,
    `quantity` INT UNSIGNED NOT NULL,

    PRIMARY KEY (`product_id`),
    FOREIGN KEY (`product_id`) REFERENCES products(`id`) ON DELETE CASCADE
);