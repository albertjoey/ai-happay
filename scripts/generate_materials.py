#!/usr/bin/env python3
import random
from datetime import datetime

types = ["image_text", "short_video", "long_video", "short_drama", "drama", "novel"]
type_names = {
    "image_text": "图文",
    "short_video": "短视频", 
    "long_video": "长视频",
    "short_drama": "短剧",
    "drama": "漫剧",
    "novel": "小说"
}
categories = ["搞笑", "美食", "旅行", "音乐", "舞蹈", "游戏", "宠物", "健身", "科技", "时尚"]
counts = {
    "image_text": 50,
    "short_video": 50,
    "long_video": 30,
    "short_drama": 20,
    "drama": 30,
    "novel": 20
}

print("-- 生成6种物料类型的测试数据")
print(f"-- 生成时间: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}")
print()

for t in types:
    print(f"-- {type_names[t]} ({t})")
    for i in range(1, counts[t] + 1):
        title = f"{type_names[t]}作品{i:03d}"
        subtitle = f"精彩{type_names[t]}内容"
        category = random.choice(categories)
        author = f"创作者{random.randint(1, 100):03d}号"
        view_count = random.randint(10000, 500000)
        like_count = random.randint(1000, 50000)
        comment_count = random.randint(100, 5000)
        share_count = random.randint(100, 5000)
        collect_count = random.randint(1000, 20000)
        chapter_count = 0
        word_count = 0
        duration = 0
        
        if t == "novel":
            word_count = random.randint(100000, 500000)
            chapter_count = random.randint(50, 300)
        elif t in ["drama", "short_drama"]:
            chapter_count = random.randint(20, 100)
        elif t == "long_video":
            duration = random.randint(1800, 7200)
        elif t == "short_video":
            duration = random.randint(15, 75)
        
        sql = f"""INSERT INTO material (tenant_id, title, subtitle, type, cover_url, content_url, description, author, category, view_count, like_count, comment_count, share_count, collect_count, duration, word_count, chapter_count, status, sort, created_at, updated_at) VALUES (1, '{title}', '{subtitle}', '{t}', 'https://picsum.photos/seed/{t}{i}/400/600', '', '这是一部精彩的{type_names[t]}作品，内容丰富，值得一看！', '{author}', '{category}', {view_count}, {like_count}, {comment_count}, {share_count}, {collect_count}, {duration}, {word_count}, {chapter_count}, 1, {i}, NOW(), NOW());"""
        print(sql)
    print()
