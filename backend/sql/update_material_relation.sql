-- 更新金刚位表,添加物料ID字段
ALTER TABLE diamond ADD COLUMN material_id bigint unsigned DEFAULT 0 COMMENT '关联物料ID' AFTER channel_id;

-- 更新推荐位表,修改content_ids为物料ID数组
-- content_ids字段已经存在,存储的是物料ID数组

-- 更新Feed流配置表,添加物料ID列表字段
ALTER TABLE feed_config ADD COLUMN material_ids json DEFAULT NULL COMMENT '物料ID列表' AFTER filter_rule;

-- 为金刚位更新测试数据,关联物料
UPDATE diamond SET material_id = 1 WHERE id = 1;
UPDATE diamond SET material_id = 2 WHERE id = 2;
UPDATE diamond SET material_id = 3 WHERE id = 3;
UPDATE diamond SET material_id = 4 WHERE id = 4;
UPDATE diamond SET material_id = 5 WHERE id = 5;
UPDATE diamond SET material_id = 6 WHERE id = 6;

-- 为推荐位更新测试数据,设置物料ID数组
UPDATE recommend SET content_ids = '[1, 2, 3, 4, 5, 6]' WHERE id = 1;
UPDATE recommend SET content_ids = '[7, 8, 9, 10, 11, 12]' WHERE id = 2;

-- 为Feed流配置更新测试数据,设置物料ID列表
UPDATE feed_config SET material_ids = '[1, 2, 3, 4, 5, 6, 7, 8, 9, 10]' WHERE id = 1;
UPDATE feed_config SET material_ids = '[11, 12, 13, 14, 15, 16, 17]' WHERE id = 2;

-- 查询验证
SELECT id, title, channel_id, material_id FROM diamond LIMIT 5;
SELECT id, title, channel_id, content_ids FROM recommend LIMIT 2;
SELECT id, title, channel_id, material_ids FROM feed_config LIMIT 2;
