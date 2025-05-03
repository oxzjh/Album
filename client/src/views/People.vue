<template>
    <div class="face-container">
        <div class="avatar-container" v-for="face in people" :key="face.id + face.fid">
            <ElAvatar shape="circle" :size="100" :src="getFaceURL(face.id, face.fid)" @click="onFaceClick(face)" />
            <span>{{ `共${face.count}张` }}</span>
        </div>
    </div>
</template>

<script setup>
import { ElAvatar } from 'element-plus';
import { onMounted, shallowRef } from 'vue';
import { getFaceURL, post } from '../../request';
import router from '@/router';

const people = shallowRef([])

onMounted(() => {
    post("/album/people", {}, result => {
        people.value = result.people
    })
})

function onFaceClick(face) {
    router.push("/people/" + face.cid)
}
</script>

<style scoped>
.face-container {
    display: flex;
    flex-wrap: wrap;
}

.avatar-container {
    margin: 20px;
    display: flex;
    flex-direction: column;
    align-items: center;
}
</style>