<template>
    <ElTimeline>
        <ElTimelineItem v-for="image in location" :key="image.id" :timestamp="`${image.locs} (${image.count}张)`"
            placement="top">
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

const location = shallowRef([])

onMounted(() => {
    post("/album/location", {}, result => {
        location.value = result.images
    })
})

function onImageClick(image) {
    router.push("/location/"+image.locs)
}
</script>

<style scoped>
.thumbnail {
    min-width: 100px;
    min-height: 100px;
}
</style>