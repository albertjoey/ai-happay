-- 添加热度分数字段到content表
ALTER TABLE content ADD COLUMN hot_score DECIMAL(10,2) DEFAULT 0 COMMENT '热度分数' AFTER share_count;
ALTER TABLE content ADD INDEX idx_hot_score (hot_score);

-- 优化标签和话题索引
ALTER TABLE content_tag ADD INDEX idx_tag_id (tag_id);
ALTER TABLE content_topic ADD INDEX idx_topic_id (topic_id);

-- 添加扩展字段
ALTER TABLE content ADD COLUMN extra JSON COMMENT '扩展信息' AFTER hot_score;
