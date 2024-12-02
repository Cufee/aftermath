-- Create "moderation_requests" table
CREATE TABLE `moderation_requests` (`id` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `moderator_comment` text NULL, `context` text NULL, `reference_id` text NOT NULL, `action_reason` text NULL, `action_status` text NOT NULL, `data` json NOT NULL, `requestor_id` text NOT NULL, `moderator_id` text NULL, PRIMARY KEY (`id`), CONSTRAINT `moderation_requests_users_moderation_requests` FOREIGN KEY (`requestor_id`) REFERENCES `users` (`id`) ON DELETE CASCADE, CONSTRAINT `moderation_requests_users_moderation_actions` FOREIGN KEY (`moderator_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION);
-- Create index "moderationrequest_id" to table: "moderation_requests"
CREATE INDEX `moderationrequest_id` ON `moderation_requests` (`id`);
-- Create index "moderationrequest_reference_id" to table: "moderation_requests"
CREATE INDEX `moderationrequest_reference_id` ON `moderation_requests` (`reference_id`);
-- Create index "moderationrequest_requestor_id" to table: "moderation_requests"
CREATE INDEX `moderationrequest_requestor_id` ON `moderation_requests` (`requestor_id`);
-- Create index "moderationrequest_moderator_id" to table: "moderation_requests"
CREATE INDEX `moderationrequest_moderator_id` ON `moderation_requests` (`moderator_id`);
-- Create index "moderationrequest_requestor_id_reference_id" to table: "moderation_requests"
CREATE INDEX `moderationrequest_requestor_id_reference_id` ON `moderation_requests` (`requestor_id`, `reference_id`);
-- Create index "moderationrequest_requestor_id_reference_id_action_status" to table: "moderation_requests"
CREATE INDEX `moderationrequest_requestor_id_reference_id_action_status` ON `moderation_requests` (`requestor_id`, `reference_id`, `action_status`);
