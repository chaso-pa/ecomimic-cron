-- CreateTable
CREATE TABLE `articles` (
    `id` VARCHAR(191) NOT NULL,
    `url` VARCHAR(191) NOT NULL,
    `title` TEXT NOT NULL,
    `content` TEXT NOT NULL,
    `content_hash` VARCHAR(191) NOT NULL,
    `published` DATETIME(3) NOT NULL,
    `crawled_at` DATETIME(3) NOT NULL,
    `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` DATETIME(3) NOT NULL,

    UNIQUE INDEX `articles_content_hash_key`(`content_hash`),
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE `article_domains` (
    `id` VARCHAR(191) NOT NULL,
    `domain` VARCHAR(191) NOT NULL,
    `url` VARCHAR(191) NOT NULL,
    `article_url_base` VARCHAR(191) NOT NULL,
    `crawl_status` BOOLEAN NOT NULL,
    `container_selector` VARCHAR(191) NOT NULL,
    `article_link_selector` VARCHAR(191) NOT NULL,
    `article_container_selector` VARCHAR(191) NOT NULL,
    `title_selector` VARCHAR(191) NOT NULL,
    `content_selector` VARCHAR(191) NOT NULL,
    `published_at_selector` VARCHAR(191) NOT NULL,
    `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updated_at` DATETIME(3) NOT NULL,

    UNIQUE INDEX `article_domains_url_key`(`url`),
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
