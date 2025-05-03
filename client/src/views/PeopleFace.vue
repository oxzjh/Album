<template>
    <ElPageHeader class="header" :icon="ArrowLeft" title="返回" @back="onBackClick">
        <template #extra>
            <ElSwitch class="switch" v-model="thumbnailMode" active-text="缩略图模式" inactive-text="头像模式" />
        </template>
    </ElPageHeader>
    <ElDivider />
    <div v-if="thumbnailMode">
        <ImageList :images="faces" :preview-list="previewList" />
    </div>
    <div v-else>
        <div class="avatar-container">
            <ElAvatar class="avatar" v-for="face, index in faces" :key="face.id + face.fid" shape="circle" :size="100"
                :src="getFaceURL(face.id, face.fid)" @click="onAvatarClick(index)" />
        </div>
        <ElImageViewer v-if="showImageViewer" :url-list="previewList" :initial-index="initIndex"
            @close="showImageViewer = false" />
    </div>
</template>

<script setup>
import { onMounted, ref, shallowRef } from 'vue';
import { useRoute } from 'vue-router';
import { getFaceURL, getImageURL, post } from '../../request';
import { ElAvatar, ElDivider, ElImage, ElImageViewer, ElPageHeader, ElSwitch } from 'element-plus';
import { ArrowLeft } from '@element-plus/icons-vue';
import router from '@/router';
import ImageList from './ImageList.vue';

const route = useRoute()
const cid = parseInt(route.params.cid)

const thumbnailMode = ref(false)
const faces = shallowRef([])
const previewList = ref([])
const initIndex = ref(0)
const showImageViewer = ref(false)

onMounted(() => {
    post("/album/people/face", { cid }, result => {
        for (let face of result.faces) {
            previewList.value.push(getImageURL(face.path))
        }
        faces.value = result.faces
    })
})

function onBackClick() {
    router.back()
}

function onAvatarClick(index) {
    initIndex.value = index
    showImageViewer.value = true
}
</script>

<style scoped>
.header {
    margin-top: 20px;
}

.switch {
    margin-right: 10px;
    --el-switch-on-color: #13ce66;
    --el-switch-off-color: #ff4949;
}

.avatar-container {
    display: flex;
    flex-wrap: wrap;
}

.avatar {
    margin: 10px;
}
</style>