--
-- Структура таблицы `language`
--

CREATE TABLE `language` (
  `id` char(36) NOT NULL,
  `date_entered` datetime DEFAULT NULL,
  `date_modified` datetime DEFAULT NULL,
  `deleted` tinyint(1) DEFAULT 0,
  `key` char(32) DEFAULT NULL,
  `title` char(32) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Дамп данных таблицы `language`
--

INSERT INTO `language` (`id`, `date_entered`, `date_modified`, `deleted`, `key`, `title`) VALUES
('6e22e7c2-d619-11ec-9d64-0242ac120002', '2022-05-17 22:42:22', '2022-05-17 22:42:22', 0, 'ru', 'Russian'),
('7b8fb69c-d619-11ec-9d64-0242ac120002', '2022-05-17 22:42:22', '2022-05-17 22:42:22', 0, 'en', 'English');

-- --------------------------------------------------------

--
-- Структура таблицы `user`
--

CREATE TABLE `user` (
  `id` char(36) NOT NULL,
  `id_module` char(36) DEFAULT NULL,
  `login` char(36) DEFAULT NULL,
  `password` char(36) DEFAULT NULL,
  `date_entered` datetime DEFAULT NULL,
  `date_modified` datetime DEFAULT NULL,
  `deleted` tinyint(1) DEFAULT 0,
  `created_by` char(36) DEFAULT NULL,
  `assigned_user_id` char(36) DEFAULT NULL,
  `modified_user_id` char(36) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Дамп данных таблицы `user`
--

INSERT INTO `user` (`id`, `id_module`, `login`, `password`, `date_entered`, `date_modified`, `deleted`, `created_by`, `assigned_user_id`, `modified_user_id`) VALUES
('a6da9bc8-d619-11ec-9d64-0242ac120002', 'a6da9bc8-d619-11ec-9d64-0242ac120002', 'alex', '81dc9bdb52d04dc20036dbd8313ed055', '2022-05-17 22:43:34', '2022-05-17 22:43:34', 0, NULL, NULL, NULL);

-- --------------------------------------------------------

--
-- Структура таблицы `user_attr`
--

CREATE TABLE `user_attr` (
  `id_user` char(36) NOT NULL,
  `id_language` char(36) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `description` text DEFAULT NULL,
  `phone` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Дамп данных таблицы `user_attr`
--

INSERT INTO `user_attr` (`id_user`, `id_language`, `name`, `description`, `phone`) VALUES
('a6da9bc8-d619-11ec-9d64-0242ac120002', '6e22e7c2-d619-11ec-9d64-0242ac120002', 'Alex', 'Alex description', '123456789'),
('a6da9bc8-d619-11ec-9d64-0242ac120002', '7b8fb69c-d619-11ec-9d64-0242ac120002', 'Jon', 'Jon description', '987654321');

--
-- Индексы сохранённых таблиц
--

--
-- Индексы таблицы `language`
--
ALTER TABLE `language`
  ADD PRIMARY KEY (`id`);

--
-- Индексы таблицы `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`);

--
-- Индексы таблицы `user_attr`
--
ALTER TABLE `user_attr`
  ADD PRIMARY KEY (`id_user`,`id_language`),
  ADD KEY `user_attr-language` (`id_language`);

--
-- Ограничения внешнего ключа сохраненных таблиц
--

--
-- Ограничения внешнего ключа таблицы `user_attr`
--
ALTER TABLE `user_attr`
  ADD CONSTRAINT `user_attr-language` FOREIGN KEY (`id_language`) REFERENCES `language` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `user_attr-user` FOREIGN KEY (`id_user`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;
