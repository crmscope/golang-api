SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


CREATE TABLE `language` (
  `id` int(32) NOT NULL,
  `date_entered` datetime DEFAULT NULL,
  `date_modified` datetime DEFAULT NULL,
  `deleted` tinyint(1) DEFAULT 0,
  `key` char(32) DEFAULT NULL,
  `title` char(32) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


INSERT INTO `language` (`id`, `date_entered`, `date_modified`, `deleted`, `key`, `title`) VALUES
(1, '2022-03-06 19:09:10', '2022-03-06 19:09:10', 0, 'ru', 'Русский'),
(2, '2022-03-06 19:09:10', '2022-03-06 19:09:10', 0, 'en', 'Английский');


CREATE TABLE `module` (
  `id` int(32) NOT NULL,
  `date_entered` datetime DEFAULT NULL,
  `date_modified` datetime DEFAULT NULL,
  `deleted` tinyint(1) DEFAULT 0,
  `key` char(32) DEFAULT NULL,
  `language_default_id` int(32) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


INSERT INTO `module` (`id`, `date_entered`, `date_modified`, `deleted`, `key`, `language_default_id`) VALUES
(1, '2022-03-06 19:08:42', '2022-03-06 19:08:42', 0, 'user', 2);


CREATE TABLE `module_description` (
  `id_language` int(32) NOT NULL,
  `id_module` int(32) NOT NULL,
  `title` char(32) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


INSERT INTO `module_description` (`id_language`, `id_module`, `title`, `description`) VALUES
(1, 1, 'Пользователи', NULL),
(2, 1, 'user', NULL);


CREATE TABLE `sys_users` (
  `id` int(32) NOT NULL,
  `id_module` int(32) DEFAULT NULL,
  `login` char(36) DEFAULT NULL,
  `password` char(36) DEFAULT NULL,
  `date_entered` datetime DEFAULT NULL,
  `date_modified` datetime DEFAULT NULL,
  `deleted` tinyint(1) DEFAULT 0,
  `created_by` char(36) DEFAULT NULL,
  `assigned_user_id` int(32) DEFAULT NULL,
  `modified_user_id` int(32) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


INSERT INTO `sys_users` (`id`, `id_module`, `login`, `password`, `date_entered`, `date_modified`, `deleted`, `created_by`, `assigned_user_id`, `modified_user_id`) VALUES
(1, NULL, 'alex', '81dc9bdb52d04dc20036dbd8313ed055', '2022-03-07 21:51:02', '2022-03-07 21:51:02', 0, NULL, NULL, NULL),
(2, 1, 'user2', '5f1b22728264c8fc163d99dfe826b9e8', '2022-04-04 23:31:54', '2022-04-04 23:31:54', 0, '1', 1, 1);


CREATE TABLE `sys_users_attr` (
  `id_users` int(32) NOT NULL,
  `id_language` int(32) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `description` text DEFAULT NULL,
  `phone` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


INSERT INTO `sys_users_attr` (`id_users`, `id_language`, `name`, `description`, `phone`) VALUES
(1, 1, 'Alexender', 'Other user', '54353535');

ALTER TABLE `language`
  ADD PRIMARY KEY (`id`);


ALTER TABLE `module`
  ADD PRIMARY KEY (`id`);


ALTER TABLE `module_description`
  ADD PRIMARY KEY (`id_language`,`id_module`),
  ADD KEY `module_description_ibfk_1` (`id_module`);

ALTER TABLE `sys_users`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `sys_users_attr`
  ADD PRIMARY KEY (`id_users`),
  ADD KEY `sys_users_attr_langua_1` (`id_language`);

ALTER TABLE `language`
  MODIFY `id` int(32) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

ALTER TABLE `module`
  MODIFY `id` int(32) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

ALTER TABLE `sys_users`
  MODIFY `id` int(32) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

ALTER TABLE `module_description`
  ADD CONSTRAINT `module_description_ibfk_1` FOREIGN KEY (`id_module`) REFERENCES `module` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `module_description_language_ibfk_1` FOREIGN KEY (`id_language`) REFERENCES `language` (`id`) ON DELETE CASCADE;

ALTER TABLE `sys_users`
  ADD CONSTRAINT `sys_users_module_ibfk_1` FOREIGN KEY (`id_module`) REFERENCES `module` (`id`) ON DELETE CASCADE;

ALTER TABLE `sys_users_attr`
  ADD CONSTRAINT `sys_users_attr_ibfk_1` FOREIGN KEY (`id_users`) REFERENCES `sys_users` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `sys_users_attr_langua_1` FOREIGN KEY (`id_language`) REFERENCES `language` (`id`) ON DELETE CASCADE;
COMMIT;

