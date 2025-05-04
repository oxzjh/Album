<template>
    <ElFormItem v-for="_, index in paths" :key="index" label="图片目录：">
        <ElInput v-model="paths[index]" :disabled="status != 0" placeholder="请输入图片目录">
            <template #append>
                <span class="delete" @click="onDeletePathClick(index)">删除</span>
            </template>
        </ElInput>
    </ElFormItem>
    <ElButton type="success" @click="onAddPathClick" :disabled="status != 0">新增目录</ElButton>
    <ElButton type="primary" @click="onStartClick" :disabled="status != 0">开始扫描</ElButton>
    <ElButton type="danger" @click="onStopClick" :disabled="status == 0">停止扫描</ElButton>
    <ElProgress v-if="status != 0" :percentage="total > 0 ? processed * 100 / total : 100">
        {{ processed }} / {{ total }}
    </ElProgress>
</template>

<script setup>
import { ElButton, ElFormItem, ElInput, ElProgress } from 'element-plus';
import { onMounted, onUnmounted, ref } from 'vue';
import { post } from '../../request';

const paths = ref([])
const status = ref(0)
const processed = ref(0)
const total = ref(0)

let statusTimer = null

onMounted(() => {
    setTimeout(() => {
        post("/album/paths", {}, result => {
            paths.value = result.paths
        })
    }, 100)
    statusTimer = setInterval(getStatus, 1000);
    // getStatus()
})

onUnmounted(() => {
    clearInterval(statusTimer)
})

function getStatus() {
    post("/album/status", {}, result => {
        status.value = result.status
        processed.value = result.processed
        total.value = result.total
    }, true)
}

function onDeletePathClick(index) {
    if (status.value == 0) {
        paths.value.splice(index, 1)
    }
}

function onAddPathClick() {
    paths.value.push("")
}

function onStartClick() {
    status.value = 1
    post("/album/start", { paths: paths.value })
}

function onStopClick() {
    post("/album/stop", {}, () => {
        status.value = 0
    })
}
</script>

<style scoped>
.delete {
    color: red;
    cursor: pointer;
}
</style>