<template>
    <div v-if="title">
        <ElPageHeader class="header" :icon="ArrowLeft" title="返回" :content="title" @back="onBackClick" />
        <ElDivider />
    </div>
    <div class="image-container">
        <ImageItem v-for="image, index in images" class="image-item" :key="image.id" :src="getThumbnailURL(image.id)"
            :src_list="previewList || srcList" :initial_index="initIndex" @mousedown="onImageClick(index)" />
    </div>
</template>

<script setup>
import { ElDivider, ElPageHeader } from 'element-plus';
import { getImageURL, getThumbnailURL } from '../../request';
import ImageItem from './ImageItem.vue';
import { ArrowLeft } from '@element-plus/icons-vue';
import router from '@/router';
import { ref, shallowRef, watch } from 'vue';

const props = defineProps(["title", "images" ,"previewList"])
const srcList = shallowRef([])
const initIndex = ref(0)

watch(() => props.images, images => {
    const a = []
    for (let image of images) {
        a.push(getImageURL(image.path))
    }
    srcList.value = a
})

function onBackClick() {
    router.back()
}

function onImageClick(index) {
    initIndex.value = index
}
</script>

<style scoped>
.header {
    margin-top: 20px;
}

.image-container {
    display: flex;
    flex-wrap: wrap;
}

.image-item {
    margin-left: 10px;
    margin-bottom: 10px;
    height: 150px;
}
</style>