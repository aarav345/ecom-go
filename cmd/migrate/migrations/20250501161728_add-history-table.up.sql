CREATE TABLE IF NOT EXISTS history (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` INT UNSIGNED NOT NULL,
    `product_name` VARCHAR(255) NOT NULL,
    `product_image` VARCHAR(255),
    `quantity` INT NOT NULL,
    `price` DECIMAL(10, 2) NOT NULL,
    `status` ENUM('pending', 'completed', 'cancelled') NOT NULL DEFAULT 'pending',
    `order_date` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES users(`id`) ON DELETE CASCADE
);