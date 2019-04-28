
-- +migrate Up
CREATE TABLE users (
   id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
   name varchar(191),
   email varchar(191),
   password varchar(191),
   created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
   updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
   deleted_at timestamp NULL DEFAULT NULL,
   PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- +migrate StatementBegin
INSERT INTO `dragon`.`users` (`id`, `name`, `email`, `password`, `created_at`, `updated_at`) VALUES ('1', 'Dhian', 'dhian.alyusi@gmail.com', '$2a$11$5pR0mp1/oV8H8dxJcsczDuer4c9D5P9YPQrzxrggYgQFvWpQlvvcm', '2019-04-20 02:49:07', '2019-04-20 02:49:07');
-- +migrate StatementEnd

-- +migrate Down
DROP TABLE users;