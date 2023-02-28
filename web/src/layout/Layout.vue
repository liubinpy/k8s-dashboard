<template>
    <a-layout>
      <!-- 固钉 -->
      <a-affix>
           <!-- 头部 -->
          <a-layout-header></a-layout-header>
      </a-affix>

      <!-- 中部  68px是header的高度 -->
      <a-layout style="height:calc(100vh - 68px)">
        <!-- 侧边栏 -->
        <a-layout-sider width="240px"  v-model:collapsed="collapsed" collapsible></a-layout-sider>

        <!-- main -->
        <a-layout style="padding: 0 24px">
            <a-layout-content :style="{
                background: 'rgb(31, 30, 30)',
                margin: 0,
                minHeight: '280px',
                overflowY: 'auto'}">
                <router-view></router-view>
            </a-layout-content>
            <a-layout-footer style="text-align: center">
                &copy;2023 create by Bennie
            </a-layout-footer>
        </a-layout>
      </a-layout>
    </a-layout>
</template>

<script setup>
import { ref } from "vue";

const collapsed = ref(false)


</script>

<style  scoped>
.ant-layout-footer {
    padding: 5px 50px  !important;
    color: rgb(239, 239, 239);
}
</style>