<template>
    <ImageList :title="locs" :images="images" />
</template>

<script setup>
import { onMounted, shallowRef } from 'vue';
import { useRoute } from 'vue-router';
import ImageList from './ImageList.vue';
import { post } from '../../request';

const route = useRoute()
const locs = route.params.locs
const a = locs.split("|")
const loc = a[a.length - 1]
const images = shallowRef([])

onMounted(() => {
    post("/album/loc", { loc, page: 0, limit: 1000 }, result => {
        images.value = result.images
    })
})
</script>