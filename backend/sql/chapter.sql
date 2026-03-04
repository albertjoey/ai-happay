-- 章节表
CREATE TABLE IF NOT EXISTS `material_chapter` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `tenant_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '租户ID',
  `material_id` bigint unsigned NOT NULL COMMENT '物料ID',
  `title` varchar(255) NOT NULL COMMENT '章节标题',
  `content` longtext COMMENT '章节内容',
  `word_count` int unsigned DEFAULT 0 COMMENT '字数',
  `sort` int unsigned DEFAULT 0 COMMENT '排序',
  `is_free` tinyint unsigned DEFAULT 0 COMMENT '是否免费 0否 1是',
  `price` int unsigned DEFAULT 0 COMMENT '价格(积分)',
  `status` tinyint unsigned DEFAULT 1 COMMENT '状态 0禁用 1启用',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_material_id` (`material_id`),
  KEY `idx_tenant_id` (`tenant_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='物料章节表';

-- 为小说添加测试章节
INSERT INTO `material_chapter` (`tenant_id`, `material_id`, `title`, `content`, `word_count`, `sort`, `is_free`, `status`) VALUES
(1, 4, '第一章 初入江湖', '李云飞站在山巅，望着远处的云海，心中充满了对未来的憧憬。

他从小生活在这个偏僻的小村庄里，每天除了练武就是帮父亲干农活。虽然生活清苦，但他从未抱怨过。

"云飞，该回家吃饭了！"母亲的声音从山下传来。

"来了！"李云飞应了一声，最后看了一眼远方的云海，转身向山下跑去。

他不知道的是，这一眼，竟是他平静生活的终点。第二天，一队黑衣人闯入村庄，打破了这里的宁静...', 1500, 1, 1, 1),
(1, 4, '第二章 命运转折', '黑衣人首领冷冷地看着李云飞："交出《天机诀》，饶你不死。"

李云飞茫然不知所措："什么天机诀？我不知道你们在说什么。"

"敬酒不吃吃罚酒！"黑衣人首领一挥手，手下立刻冲了上来。

李云飞虽然年纪尚小，但自幼跟随父亲习武，身手不凡。他奋力抵抗，但终究寡不敌众。

就在他即将被制服的时候，父亲突然冲了出来，一掌击退了几个黑衣人。

"快走！带着你母亲走！"父亲大喊道，"记住，永远不要回来！"

李云飞含泪带着母亲逃离了村庄，从此开始了他的江湖之路...', 1800, 2, 1, 1),
(1, 4, '第三章 拜师学艺', '逃亡的路上，李云飞和母亲遇到了一位白发苍苍的老者。

老者看着李云飞，眼中闪过一丝惊讶："好苗子，真是好苗子！"

"前辈，您在说什么？"李云飞疑惑地问道。

"小子，你可愿意拜我为师？"老者问道。

李云飞想起父亲的嘱托，知道自己必须变强才能保护母亲，于是毫不犹豫地跪了下来："弟子愿意！"

老者哈哈大笑："好！从今天起，你就是我的关门弟子了。我乃青云子，这套《青云剑法》就传给你了。"

从此，李云飞开始了艰苦的修炼生涯...', 1600, 3, 1, 1),
(1, 4, '第四章 初露锋芒', '三年后，李云飞已经长成了一个英俊的少年。

这一天，青云子把他叫到跟前："云飞，你的剑法已经小有所成，是时候下山历练了。"

"师父，我..."李云飞有些不舍。

"去吧，江湖才是最好的老师。记住，无论遇到什么困难，都要坚持自己的本心。"

李云飞拜别师父，踏上了江湖之路。他不知道的是，一场惊天阴谋正在等着他...', 1400, 4, 0, 1),
(1, 4, '第五章 江湖险恶', '李云飞来到一座繁华的城市，这里人来人往，热闹非凡。

他走进一家酒楼，点了一壶酒，正准备休息一下，却听到隔壁桌的谈话。

"听说了吗？李家庄被灭门了，据说是因为一本叫《天机诀》的秘籍。"

李云飞心中一震，那不是他的家乡吗？他强忍着悲痛，继续听着。

"据说李家有个儿子逃走了，现在各方势力都在找他呢。"

李云飞握紧了拳头，他知道，自己的身份已经暴露，必须更加小心...', 1700, 5, 0, 1);

-- 为漫画/短剧添加测试章节
INSERT INTO `material_chapter` (`tenant_id`, `material_id`, `title`, `content`, `word_count`, `sort`, `is_free`, `status`) VALUES
(1, 5, '第1话 命运的开始', 'https://picsum.photos/seed/drama1-1/800/1200,https://picsum.photos/seed/drama1-2/800/1200,https://picsum.photos/seed/drama1-3/800/1200', 0, 1, 1, 1),
(1, 5, '第2话 意外的相遇', 'https://picsum.photos/seed/drama2-1/800/1200,https://picsum.photos/seed/drama2-2/800/1200,https://picsum.photos/seed/drama2-3/800/1200', 0, 2, 1, 1),
(1, 5, '第3话 心动的瞬间', 'https://picsum.photos/seed/drama3-1/800/1200,https://picsum.photos/seed/drama3-2/800/1200,https://picsum.photos/seed/drama3-3/800/1200', 0, 3, 0, 1),
(1, 5, '第4话 误会与解释', 'https://picsum.photos/seed/drama4-1/800/1200,https://picsum.photos/seed/drama4-2/800/1200,https://picsum.photos/seed/drama4-3/800/1200', 0, 4, 0, 1),
(1, 5, '第5话 真相大白', 'https://picsum.photos/seed/drama5-1/800/1200,https://picsum.photos/seed/drama5-2/800/1200,https://picsum.photos/seed/drama5-3/800/1200', 0, 5, 0, 1);
