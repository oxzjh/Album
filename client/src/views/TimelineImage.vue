<template>
    <ImageList :title="`${year}年${month}月`" :images="images" />
</template>

<script setup>
import { onMounted, shallowRef } from 'vue';
import { useRoute } from 'vue-router';
import ImageList from './ImageList.vue';
import { post } from '../../request';

const route = useRoute()
const year = parseInt(route.params.year)
const month = parseInt(route.params.month)
const images = shallowRef([])


onMounted(() => {
    post("/album/date", { year, month, page: 0, limit: 1000 }, result => {
        images.value = result.images
    })
})
</script>