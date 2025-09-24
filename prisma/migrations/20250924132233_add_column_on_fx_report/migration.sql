-- AlterTable
ALTER TABLE `fx_reports` ADD COLUMN `incoming_indices` JSON NULL,
    MODIFY `updated_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3);
