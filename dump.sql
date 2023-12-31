PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE `skills` (`id` integer,`created_at` datetime,`updated_at` datetime,`skill_name` text,PRIMARY KEY (`id`));
INSERT INTO skills VALUES(1,'2023-07-13 03:32:00.3824828+07:00','2023-07-13 03:32:00.3824828+07:00','a');
INSERT INTO skills VALUES(2,'2023-07-13 03:32:00.3824828+07:00','2023-07-13 03:32:00.3824828+07:00','b');
CREATE TABLE `profiles` (`id` integer,`created_at` datetime,`updated_at` datetime,`name` text,PRIMARY KEY (`id`));
INSERT INTO profiles VALUES(1,'2023-07-13 03:25:33.7645675+07:00','2023-07-13 03:25:33.7645675+07:00','board');
INSERT INTO profiles VALUES(2,'2023-07-13 03:25:33.7645675+07:00','2023-07-13 03:25:33.7645675+07:00','expert');
INSERT INTO profiles VALUES(3,'2023-07-13 03:25:33.7645675+07:00','2023-07-13 03:25:33.7645675+07:00','trainer');
INSERT INTO profiles VALUES(4,'2023-07-13 03:25:33.7645675+07:00','2023-07-13 03:25:33.7645675+07:00','competitor');
CREATE TABLE `users` (`id` integer,`created_at` datetime,`updated_at` datetime,`name` text,`email` text,`username` text,`password` text,`profile_id` integer,PRIMARY KEY (`id`),CONSTRAINT `fk_users_profile` FOREIGN KEY (`profile_id`) REFERENCES `profiles`(`id`));
INSERT INTO users VALUES(1,'2023-07-13 03:25:34.0244543+07:00','2023-07-13 03:25:34.0244543+07:00','root','root@root.com','root','$2a$10$G/YqnTTaBo61b2VF0mbJlOyARKnzofNAjmNii/5ZCTnGJwZHTIMgK',1);
INSERT INTO users VALUES(2,'2023-07-13 03:28:48.3102253+07:00','2023-07-13 03:28:48.3102253+07:00','name','mail@mail.com','abdul','$2a$10$TtFm9UpakUQPmQr6pXVc1.4JfUIXSNh1aqCGEd2RqCg/yiz0INoui',1);
INSERT INTO users VALUES(3,'2023-07-13 03:29:53.3934411+07:00','2023-07-13 03:29:53.3934411+07:00','name','mail@mail.com','rozak','$2a$10$WcGKcRZLEDcmuj5Mujdk7e0iSOk./1LutaqSHQ6T2cZI9X//RhpV.',2);
INSERT INTO users VALUES(4,'2023-07-13 03:32:00.531631+07:00','2023-07-13 03:32:00.531631+07:00','name','mail@mail.com','john','$2a$10$LZ0bjIcTxEqmEGpXkucQsOc/aqktuhtB9bJO4gvGLXcbSekAHPh32',2);
CREATE TABLE `user_skills` (`user_id` integer,`skill_id` integer,PRIMARY KEY (`user_id`,`skill_id`),CONSTRAINT `fk_user_skills_user` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`),CONSTRAINT `fk_user_skills_skill` FOREIGN KEY (`skill_id`) REFERENCES `skills`(`id`));
INSERT INTO user_skills VALUES(4,1);
INSERT INTO user_skills VALUES(4,2);
CREATE TABLE `activities` (`id` integer,`created_at` datetime,`updated_at` datetime,`skill_id` integer,`title` text,`description` text,`start_date` datetime,`end_date` datetime,PRIMARY KEY (`id`),CONSTRAINT `fk_activities_skill` FOREIGN KEY (`skill_id`) REFERENCES `skills`(`id`));
INSERT INTO activities VALUES(1,'2023-07-13 03:32:33.6770999+07:00','2023-07-13 03:32:33.6770999+07:00',1,'lorem','lorem ipsum','2023-01-04 00:00:00+00:00','2023-01-13 00:00:00+00:00');
CREATE TABLE `activity_users` (`activity_id` integer,`user_id` integer,PRIMARY KEY (`activity_id`,`user_id`),CONSTRAINT `fk_activity_users_activity` FOREIGN KEY (`activity_id`) REFERENCES `activities`(`id`),CONSTRAINT `fk_activity_users_user` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`));
CREATE UNIQUE INDEX `idx_skills_skill_name` ON `skills`(`skill_name`);
CREATE UNIQUE INDEX `idx_profiles_name` ON `profiles`(`name`);
CREATE UNIQUE INDEX `idx_users_username` ON `users`(`username`);
COMMIT;
