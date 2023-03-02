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
        <a-layout-sider width="240px"  v-model:collapsed="collapsed" collapsible>
             <a-menu
             :selectedKeys="selectedKeys2"
             :openKeys="openKeys"
             @openChange="onOpenChange"
             mode="inline"
             :style="{height: '100%',boderRight: 0}"
             >
              <!-- routers的信息就是router/index.js下面的路由信息 -->
              <template v-for="menu in routers" :key="menu">
                <!-- 处理无子路由的情况 -->
                 <a-menu-item
                   v-if="menu.children && menu.children.length === 1"
                   :index="menu.children[0].path"
                   :key="menu.children[0].path"
                   @click="routeChange('item', menu.children[0].path)"
                 >
                  <template #icon>
                    <component :is="menu.children[0].icon"></component>
                  </template>
                  <span>{{ menu.children[0].name }}</span>
                 </a-menu-item>
                 <!-- 处理有子路由的情况 -->
                 <a-sub-menu
                  v-else-if="menu.children && menu.children.length > 1"
                  :index="menu.path"
                  :key="menu.path"
                 >
                   <template #icon>
                      <component :is="menu.icon"></component>
                    </template>
                    <template #title>
                      <span>
                          <span :class="[collapsed ? 'is-collapse':'']">{{ menu.name }}</span>
                      </span>
                    </template>
                    <!-- 子路由 -->
                    <a-menu-item 
                        v-for="child in menu.children" 
                        :key="child.path" 
                        :index="child.path" 
                        @click="routeChange('sub', child.path)">
                        <span>{{ child.name }}</span>
                    </a-menu-item>
                 </a-sub-menu>
              </template>
             </a-menu>
        </a-layout-sider>

        <!-- main -->
        <a-layout style="padding: 0 24px">
          <!-- 面包屑 -->
            <a-breadcrumb style="margin: 12px 0">
              <a-breadcrumb-item>工作台</a-breadcrumb-item>
              <template v-for="(matched, index) in router.currentRoute.value.matched" :key="index">
                <a-breadcrumb-item>{{ matched.name }}</a-breadcrumb-item>
              </template>
            </a-breadcrumb>

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
import { ref, onMounted } from 'vue';
import kubeLog from '@/assets/k8s-metrics.png'
import avator from '@/assets/avator.png'
import { useRouter } from 'vue-router';

const collapsed = ref(false)
const selectedKeys1 = ref([])
const clusterList = ref(["TEST1", "TEST2"])
// 获取router
const router = useRouter()
// 路由信息
const routers = ref([])
const selectedKeys2 = ref([])
const openKeys = ref([])

// 切换路由
const routeChange = (type, path) => {
    // 如果是单独的栏目
    if (type != 'sub'){
        openKeys.value = []
    }

    // 选中当前path对应
    selectedKeys2.value = [path]

    // 页面跳转
    if (router.currentRoute.value.path !== path){
        router.push(path)
    }
}

// 从浏览器地址直接打开后的选择
function DirectOpen(value) {
    selectedKeys2.value = [value[1].path]
    openKeys.value = [value[0].path]
}

const onOpenChange = (value) => {
    const latestOpenKey = value.find(key => openKeys.value.indexOf(key) === -1)
    openKeys.value = latestOpenKey ? [latestOpenKey]:[] 
}

const logout = () => {
    // 移除用户信息
    localStorage.removeItem('username')
    localStorage.removeItem('token')
    router.push('/login')
}

onMounted(() => {
    routers.value = router.options.routes

    DirectOpen(router.currentRoute.value.matched)
})
</script>

<style  scoped>
.ant-layout-footer {
    padding: 5px 50px  !important;
    color: rgb(239, 239, 239);
}

.ant-layout-header {
    padding: 0 30px !important;
}

.is-collapse{
    display: none;
}
</style>