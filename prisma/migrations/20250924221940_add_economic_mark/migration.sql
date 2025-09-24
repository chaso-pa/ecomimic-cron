-- CreateTable
CREATE TABLE `economic_marks` (
    `id` VARCHAR(191) NOT NULL,
    `published_at` DATETIME(3) NOT NULL,
    `country` VARCHAR(191) NOT NULL,
    `title` VARCHAR(191) NOT NULL,
    `importance` INTEGER NOT NULL,
    `estimate` VARCHAR(191) NOT NULL,
    `result` VARCHAR(191) NOT NULL,
    `past_result` VARCHAR(191) NOT NULL,
    `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),

    UNIQUE INDEX `economic_marks_title_published_at_key`(`title`, `published_at`),
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
