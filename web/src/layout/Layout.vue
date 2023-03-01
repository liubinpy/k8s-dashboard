<template>
    <a-layout>
      <!-- 固钉 -->
      <a-affix>
           <!-- 头部 -->
          <a-layout-header>
            <!-- 平台信息 -->
            <div style="float: left;">
                <img style="height: 40px" :src="kubeLog">
                <span style="font-size: 16px; padding: 0 10px 0 10px; font-weight: bold;">k8sDashboard</span>
            </div>
            <!-- 集群信息 -->
            <a-menu style="float: left; width: 250px;" v-model:selectedKeys="selectedKeys1" theme="dark" mode="horizontal">
              <a-menu-item v-for="item in clusterList" :key="item">{{ item }}</a-menu-item>
            </a-menu>
            <!-- 用户信息 -->
            <div style="float: right;">
              <img :src="avator" style="height: 40px; border-radius: 50%; margin-right: 10px">
              <a-dropdown :overlayStyle="{paddingTop: '20px'}">
                <a >Admin <down-outlined /></a>
                <template #overlay>
                    <a-menu>
                      <a-menu-item><a @click="logout">退出登陆</a></a-menu-item>
                      <a-menu-item><a >修改密码</a></a-menu-item>
                    </a-menu>
                </template>
              </a-dropdown>
            </div>

          </a-layout-header>
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
import { ref } from 'vue';
import kubeLog from '@/assets/k8s-metrics.png'
import avator from '@/assets/avator.png'
import { useRouter } from 'vue-router';

const collapsed = ref(false)
const selectedKeys1 = ref([])
const clusterList = ref(["TEST1", "TEST2"])

const router = useRouter()

const logout = () => {
    // 移除用户信息
    localStorage.removeItem('username')
    localStorage.removeItem('token')
    router.push('/login')
}
</script>

<style  scoped>
.ant-layout-footer {
    padding: 5px 50px  !important;
    color: rgb(239, 239, 239);
}

.ant-layout-header {
    padding: 0 30px !important;
}
</style>