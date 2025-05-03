<template>
    <div class="flex">
        <ElFormItem class="input" label="内容：">
            <ElInput v-model="text" placeholder="请输入内容，如一只白色的小狗" @input="onInputChange" />
        </ElFormItem>
        <ElFormItem class="top" label="数量：">
            <ElInputNumber class="number" v-model="top" controls-position="right" :min="1" :max="100" />
        </ElFormItem>
        <ElButton :icon="Search" type="primary" @click="onSearchClick" disabled>搜索</ElButton>
    </div>
    <ElUpload accept=".jpg,.jpeg,.png,.webp,.tif,.tiff" :auto-upload="false" drag :on-change="onUploadChange">
        <ElIcon class="el-icon--upload">
            <UploadFilled />
        </ElIcon>
        <div class="el-upload__text">
            拖拽图片到这里 <em>图搜图</em>
        </div>
    </ElUpload>
    <ImageList :images="images" />
</template>

<script setup>
import { Search, UploadFilled } from '@element-plus/icons-vue';
import { ElButton, ElFormItem, ElIcon, ElInput, ElInputNumber, ElMessage, ElUpload } from 'element-plus';
import { ref, shallowRef } from 'vue';
import ImageList from './ImageList.vue';
import { post } from '../../request';

const text = ref("")
const top = ref(10)
const images = shallowRef([])

function onInputChange() {
    if (text.value == "") {
        images.value = []
    } else {
        post("/album/search", { text: text.value }, result => {
            images.value = result.images
        }, true)
    }
}

function onSearchClick() {
    if (text.value == "") {
        ElMessage({ message: "请输入搜索内容", type: "error" })
        return
    }
    post("/album/search", { text: text.value }, onSearchResponse)
}

function onSearchResponse(result) {
    images.value = result.images
    if (result.images.length == 0) {
        ElMessage({ message: "未查询到匹配的图片", type: "warning" })
    }
}

function onUploadChange(file) {
    const reader = new FileReader()
    reader.readAsDataURL(file.raw)
    reader.onload = function () {
        var index = reader.result.indexOf(",")
        post("/album/search", { image: reader.result.substr(index + 1) }, onSearchResponse)
    }
}
</script>

<style scoped>
.flex {
    display: flex;
    margin-top: 10px;
}

.input {
    width: calc(100% - 242px);
}

.top {
    margin-left: 10px;
    margin-right: 10px;
}

.number {
    width: 80px;
}
</style>