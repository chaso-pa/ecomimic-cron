-- CreateTable
CREATE TABLE `fx_reports` (
    `id` VARCHAR(191) NOT NULL,
    `symbol` VARCHAR(191) NOT NULL,
    `price` DOUBLE NOT NULL,
    `technical` TEXT NOT NULL,
    `fundamental` TEXT NOT NULL,
    `summary` TEXT NOT NULL,
    `confidence` INTEGER NOT NULL,
    `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` DATETIME(3) NOT NULL,

    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
