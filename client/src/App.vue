<template>
    <div class="container">
        <ElMenu class="menu" router default-active="setting">
            <ElMenuItem index="/search">
                <ElIcon>
                    <Search />
                </ElIcon>
                <span>搜索</span>
            </ElMenuItem>
            <ElMenuItem index="/people">
                <ElIcon>
                    <User />
                </ElIcon>
                <span>人物</span>
            </ElMenuItem>
            <ElMenuItem index="/location">
                <ElIcon>
                    <Location />
                </ElIcon>
                <span>地点</span>
            </ElMenuItem>
            <ElMenuItem index="/timeline">
                <ElIcon>
                    <Calendar />
                </ElIcon>
                <span>时间线</span>
            </ElMenuItem>
            <ElMenuItem index="/setting">
                <ElIcon>
                    <Setting />
                </ElIcon>
                <span>设置</span>
            </ElMenuItem>
        </ElMenu>
        <ElScrollbar class="content">
            <RouterView />
        </ElScrollbar>
    </div>
</template>

<script setup>
import { Calendar, Location, Search, Setting, User } from '@element-plus/icons-vue';
import { ElIcon, ElMenu, ElMenuItem, ElScrollbar } from 'element-plus';
</script>

<style scoped>
.container {
    display: flex;
}

.menu {
    width: 200px;
    height: 100vh;
    overflow: hidden;
    background-color: #f8f8ff;
}

.el-menu-item.is-active {
    background-color: rgb(219, 234, 254) !important;
}

.content {
    width: 100%;
    height: 100vh;
}
</style>