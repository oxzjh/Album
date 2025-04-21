<template>
    <ElTimeline>
        <ElTimelineItem v-for="image in timeline" :key="image.id"
            :timestamp="`${image.year}/${image.month} (${image.count}张)`" placement="top">
            <ImageItem class="thumbnail" :src="getThumbnailURL(image.id)" @click="onImageClick(image)" />
        </ElTimelineItem>
    </ElTimeline>
</template>

<script setup>
import { ElTimeline, ElTimelineItem } from 'element-plus';
import { onMounted, shallowRef } from 'vue';
import { getThumbnailURL, post } from '../../request';
import router from '@/router';
import ImageItem from './ImageItem.vue';

const timeline = shallowRef([])

onMounted(() => {
    post("/album/timeline", {}, result => {
        for (let image of result.images) {
            const d = new Date(image.ts * 1000)
            image.year = d.getFullYear()
            image.month = d.getMonth() + 1
        }
        timeline.value = result.images
    })
})

function onImageClick(image) {
    router.push("/timeline/" + image.year + "/" + image.month)
}
</script>

<style scoped>
.thumbnail {
    min-width: 100px;
    min-height: 100px;
}
</style>