-- CreateTable
CREATE TABLE `fred_indices` (
    `id` VARCHAR(191) NOT NULL,
    `crawl_status` BOOLEAN NOT NULL DEFAULT false,
    `country` VARCHAR(191) NOT NULL,
    `symbol` VARCHAR(191) NOT NULL,

    UNIQUE INDEX `fred_indices_country_symbol_key`(`country`, `symbol`),
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
