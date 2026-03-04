-- 插入测试用户
INSERT INTO `user` (`tenant_id`, `username`, `email`, `phone`, `password`, `nickname`, `avatar`, `gender`, `bio`, `status`, `role`, `follow_count`, `fans_count`, `like_count`) VALUES
(1, 'user001', 'user001@example.com', '13800138001', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iAt6Z5EH', '用户小明', 'https://picsum.photos/seed/user1/200/200', 1, '热爱生活，喜欢分享', 1, 0, 10, 100, 500),
(1, 'user002', 'user002@example.com', '13800138002', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iAt6Z5EH', '用户小红', 'https://picsum.photos/seed/user2/200/200', 2, '喜欢旅行和美食', 1, 0, 15, 200, 800),
(1, 'blogger001', 'blogger001@example.com', '13800138003', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iAt6Z5EH', '创作者阿杰', 'https://picsum.photos/seed/blogger1/200/200', 1, '专业视频创作者，分享精彩生活', 1, 1, 50, 5000, 20000),
(1, 'blogger002', 'blogger002@example.com', '13800138004', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iAt6Z5EH', '创作者小美', 'https://picsum.photos/seed/blogger2/200/200', 2, '美食博主，带你吃遍天下', 1, 1, 30, 3000, 15000),
(1, 'blogger003', 'blogger003@example.com', '13800138005', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iAt6Z5EH', '创作者大伟', 'https://picsum.photos/seed/blogger3/200/200', 1, '旅行达人，记录世界的美', 1, 1, 40, 4000, 18000);

-- 插入测试话题
INSERT INTO `topic` (`tenant_id`, `name`, `description`, `cover`, `status`, `sort`) VALUES
(1, '热门话题', '最热门的讨论话题', 'https://picsum.photos/seed/topic1/400/300', 1, 1),
(1, '搞笑日常', '分享生活中的搞笑瞬间', 'https://picsum.photos/seed/topic2/400/300', 1, 2),
(1, '美食分享', '美食爱好者的天堂', 'https://picsum.photos/seed/topic3/400/300', 1, 3),
(1, '旅行日记', '记录旅途中的美好', 'https://picsum.photos/seed/topic4/400/300', 1, 4),
(1, '生活技巧', '实用的生活小妙招', 'https://picsum.photos/seed/topic5/400/300', 1, 5);

-- 插入测试标签
INSERT INTO `tag` (`tenant_id`, `name`, `type`, `status`, `sort`) VALUES
(1, '搞笑', 0, 1, 1),
(1, '热门', 0, 1, 2),
(1, '美食', 0, 1, 3),
(1, '旅行', 0, 1, 4),
(1, '生活', 0, 1, 5),
(1, '科技', 0, 1, 6),
(1, '娱乐', 0, 1, 7),
(1, '音乐', 0, 1, 8),
(1, '游戏', 0, 1, 9),
(1, '运动', 0, 1, 10);

-- 插入测试内容
INSERT INTO `content` (`tenant_id`, `user_id`, `title`, `description`, `cover`, `type`, `status`, `view_count`, `like_count`, `comment_count`, `collect_count`, `share_count`, `publish_at`) VALUES
(1, 3, '精彩短视频：城市夜景', '记录城市夜晚的美丽瞬间，霓虹闪烁，车水马龙', 'https://picsum.photos/seed/content1/400/600', 2, 1, 12345, 5678, 234, 456, 123, NOW()),
(1, 3, '搞笑日常：猫咪的日常', '家里的小猫咪又搞怪了，笑死我了', 'https://picsum.photos/seed/content2/400/600', 2, 1, 23456, 8901, 345, 567, 234, NOW()),
(1, 4, '美食制作：家常红烧肉', '教你做最正宗的红烧肉，肥而不腻，入口即化', 'https://picsum.photos/seed/content3/400/600', 2, 1, 34567, 12345, 456, 678, 345, NOW()),
(1, 5, '旅行Vlog：云南之旅', '带你领略云南的美景，感受大自然的魅力', 'https://picsum.photos/seed/content4/400/600', 2, 1, 45678, 15678, 567, 789, 456, NOW()),
(1, 3, '热门短剧：都市爱情', '一段发生在都市的浪漫爱情故事', 'https://picsum.photos/seed/content5/400/600', 3, 1, 56789, 19012, 678, 890, 567, NOW()),
(1, 4, '漫剧推荐：热血少年', '超燃的漫剧，少年热血，永不言败', 'https://picsum.photos/seed/content6/400/600', 4, 1, 67890, 22345, 789, 901, 678, NOW()),
(1, 5, '小说连载：穿越时空', '一个现代人穿越到古代的奇幻故事', 'https://picsum.photos/seed/content7/400/600', 5, 1, 78901, 25678, 890, 1012, 789, NOW()),
(1, 3, '图文分享：生活小技巧', '10个超实用的生活小技巧，让生活更便利', 'https://picsum.photos/seed/content8/400/600', 6, 1, 89012, 29012, 901, 1123, 890, NOW()),
(1, 4, '长视频：纪录片《自然之美》', '探索大自然的奥秘，感受生命的奇迹', 'https://picsum.photos/seed/content9/400/600', 1, 1, 90123, 32345, 1012, 1234, 901, NOW()),
(1, 5, '搞笑视频：办公室趣事', '办公室里发生的搞笑事情，笑到肚子疼', 'https://picsum.photos/seed/content10/400/600', 2, 1, 101234, 35678, 1123, 1345, 1012, NOW());

-- 插入内容媒体资源
INSERT INTO `content_media` (`content_id`, `type`, `url`, `thumbnail`, `duration`, `width`, `height`, `size`, `sort`) VALUES
(1, 2, 'https://example.com/video1.mp4', 'https://picsum.photos/seed/content1/400/600', 60, 1920, 1080, 10485760, 1),
(2, 2, 'https://example.com/video2.mp4', 'https://picsum.photos/seed/content2/400/600', 45, 1920, 1080, 8388608, 1),
(3, 2, 'https://example.com/video3.mp4', 'https://picsum.photos/seed/content3/400/600', 180, 1920, 1080, 20971520, 1),
(4, 2, 'https://example.com/video4.mp4', 'https://picsum.photos/seed/content4/400/600', 300, 1920, 1080, 31457280, 1),
(5, 2, 'https://example.com/video5.mp4', 'https://picsum.photos/seed/content5/400/600', 600, 1920, 1080, 52428800, 1);

-- 插入内容话题关联
INSERT INTO `content_topic` (`content_id`, `topic_id`) VALUES
(1, 1), (2, 2), (3, 3), (4, 4), (5, 1),
(6, 1), (7, 1), (8, 5), (9, 1), (10, 2);

-- 插入内容标签关联
INSERT INTO `content_tag` (`content_id`, `tag_id`) VALUES
(1, 2), (1, 5), (2, 1), (2, 2), (3, 3),
(4, 4), (4, 5), (5, 2), (5, 7), (6, 2),
(7, 2), (8, 5), (9, 2), (10, 1), (10, 2);

-- 插入Banner
INSERT INTO `banner` (`tenant_id`, `channel_id`, `title`, `image`, `link_type`, `link_url`, `status`, `sort`) VALUES
(1, 1, '新年活动', 'https://picsum.photos/seed/banner1/800/400', 2, 'https://example.com/activity', 1, 1),
(1, 1, '热门推荐', 'https://picsum.photos/seed/banner2/800/400', 1, '', 1, 2),
(1, 2, '搞笑专区', 'https://picsum.photos/seed/banner3/800/400', 1, '', 1, 1);

-- 插入金刚位
INSERT INTO `diamond_position` (`tenant_id`, `channel_id`, `name`, `icon`, `link_type`, `status`, `sort`) VALUES
(1, 1, '热门', 'https://picsum.photos/seed/icon1/100/100', 1, 1, 1),
(1, 1, '搞笑', 'https://picsum.photos/seed/icon2/100/100', 1, 1, 2),
(1, 1, '美食', 'https://picsum.photos/seed/icon3/100/100', 1, 1, 3),
(1, 1, '旅行', 'https://picsum.photos/seed/icon4/100/100', 1, 1, 4),
(1, 1, '生活', 'https://picsum.photos/seed/icon5/100/100', 1, 1, 5);

-- 插入关注关系
INSERT INTO `follow` (`follower_id`, `following_id`) VALUES
(1, 3), (1, 4), (1, 5), (2, 3), (2, 4);

-- 插入互动数据
INSERT INTO `interaction` (`user_id`, `target_id`, `type`, `target_type`) VALUES
(1, 1, 1, 1), (1, 2, 1, 1), (1, 3, 2, 1), (2, 1, 1, 1), (2, 4, 2, 1);

-- 插入评论
INSERT INTO `comment` (`content_id`, `user_id`, `content`, `like_count`, `status`) VALUES
(1, 1, '太美了！', 10, 1),
(1, 2, '拍得真好', 8, 1),
(2, 1, '笑死我了哈哈哈', 15, 1),
(3, 2, '看起来好好吃', 12, 1),
(4, 1, '好想去云南', 20, 1);

-- 插入观看历史
INSERT INTO `view_history` (`user_id`, `content_id`, `duration`, `progress`) VALUES
(1, 1, 60, 100),
(1, 2, 45, 100),
(1, 3, 180, 100),
(2, 1, 60, 100),
(2, 4, 300, 100);

-- 插入消息通知
INSERT INTO `message` (`user_id`, `from_user_id`, `type`, `title`, `content`, `target_id`, `target_type`, `is_read`) VALUES
(3, 1, 1, '收到新的点赞', '用户小明赞了你的内容', 1, 1, 0),
(3, 2, 1, '收到新的点赞', '用户小红赞了你的内容', 1, 1, 0),
(3, 1, 2, '收到新的评论', '用户小明评论了你的内容', 1, 1, 0),
(4, 1, 4, '新增粉丝', '用户小明关注了你', 3, 0, 0),
(5, 2, 4, '新增粉丝', '用户小红关注了你', 5, 0, 0);
