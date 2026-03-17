#!/usr/bin/env python3
"""
批量替换vxe-table为a-table的脚本
"""
import os
import re

# 需要替换的文件列表
files_to_replace = [
    '/Users/leo/Desktop/happytwo/frontend/admin/src/views/user/UserList.vue',
    '/Users/leo/Desktop/happytwo/frontend/admin/src/views/content/ContentList.vue',
    '/Users/leo/Desktop/happytwo/frontend/admin/src/views/topic/TopicList.vue',
    '/Users/leo/Desktop/happytwo/frontend/admin/src/views/channel/FeedConfigList.vue',
    '/Users/leo/Desktop/happytwo/frontend/admin/src/views/channel/DiamondList.vue',
    '/Users/leo/Desktop/happytwo/frontend/admin/src/views/channel/RecommendList.vue',
    '/Users/leo/Desktop/happytwo/frontend/admin/src/views/channel/ChannelList.vue',
    '/Users/leo/Desktop/happytwo/frontend/admin/src/views/system/AdminUserList.vue',
    '/Users/leo/Desktop/happytwo/frontend/admin/src/views/system/RoleList.vue',
    '/Users/leo/Desktop/happytwo/frontend/admin/src/views/tag/TagList.vue',
    '/Users/leo/Desktop/happytwo/frontend/admin/src/views/material/ChapterList.vue',
    '/Users/leo/Desktop/happytwo/frontend/admin/src/views/material/MaterialList.vue',
]

def replace_vxe_table(content):
    """替换vxe-table为a-table"""
    # 这个脚本只是标记,实际替换需要手动完成
    # 因为每个页面的列定义不同,需要逐个处理
    return content

def main():
    print("需要替换vxe-table的文件:")
    for i, file_path in enumerate(files_to_replace, 1):
        if os.path.exists(file_path):
            print(f"{i}. ✅ {file_path}")
        else:
            print(f"{i}. ❌ {file_path} (文件不存在)")
    
    print("\n注意: 由于每个页面的列定义不同,需要逐个手动替换")
    print("建议使用AdSlotList.vue作为参考模板")

if __name__ == '__main__':
    main()
