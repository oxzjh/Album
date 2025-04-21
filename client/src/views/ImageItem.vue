<template>
    <ElImage :src="src" :preview-src-list="src_list" :initial-index="initial_index">
        <template #placeholder>
            <div class="image-bg">
                <ElIcon>
                    <Picture />
                </ElIcon>
            </div>
        </template>
        <template #error>
            <div class="image-bg">
                <ElIcon>
                    <Warning />
                </ElIcon>
            </div>
        </template>
    </ElImage>
</template>

<script setup>
import { Picture, Warning } from '@element-plus/icons-vue';
import { ElIcon, ElImage } from 'element-plus';

defineProps(["src", "src_list", "initial_index"])
</script>

<style scoped>
.image-bg {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 100%;
    height: 100%;
    min-width: 100px;
    min-height: 100px;
    background: var(--el-fill-color-light);
    color: var(--el-text-color-secondary);
    font-size: 30px;
}
</style>