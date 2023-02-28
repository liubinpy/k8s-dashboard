import { createRouter, createWebHistory } from 'vue-router'

// 导入进度条
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'

// 导入layout布局
import Layout from '@/layout/Layout'

// 定义路由规则
const routes = [
  {
    path: '/',
    redirect: '/home' // 重定向
  },
  {
    path: "/home",
    // 引入布局组件
    component: Layout,
    children: [
      {
        path: "/home",
        name: "概览",
        icon: "fund-outlined",
        meta: { title: "概览", requireAuth: true },
        component: () => import('@/views/home/Home.vue'),
      }
    ]
  },
  {
    path: "/cluster",
    name: "集群",
    component: Layout,
    icon: "cloud-server-outlined",
    children: [
      {
        path: "/cluster/node",
        name: "Node",
        meta: { title: "Node", requireAuth: true },
        component: () => import('@/views/cluster/Node.vue'),
      },
      {
        path: "/cluster/namespace",
        name: "Namespace",
        meta: { title: "Namespace", requireAuth: true },
        component: () => import('@/views/cluster/Namespace.vue'),
      },
      {
        path: "/cluster/pv",
        name: "PV",
        meta: { title: "PV", requireAuth: true },
        component: () => import('@/views/cluster/PV.vue'),
      }
    ]
  },
  {
    path: "/workload",
    name: "工作负载",
    component: Layout,
    icon: "block-outlined",
    children: [
      {
        path: "/workload/pod",
        name: "Pod",
        meta: { title: "Pod", requireAuth: true },
        component: () => import('@/views/workload/Pod.vue'),
      },
      {
        path: "/workload/deployment",
        name: "Deployment",
        meta: { title: "Deployment", requireAuth: true },
        component: () => import('@/views/workload/Deployment.vue'),
      },
      {
        path: "/workload/daemonset",
        name: "DaemonSet",
        meta: { title: "DaemonSet", requireAuth: true },
        component: () => import('@/views/workload/DaemonSet.vue'),
      },
      {
        path: "/workload/statefulset",
        name: "StatefulSet",
        meta: { title: "StatefulSet", requireAuth: true },
        component: () => import('@/views/workload/StatefulSet.vue'),
      },
    ]
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

// 定义进度条
NProgress.inc(100)
// 进度条配置
NProgress.configure({
  easing: 'ease',
  speed: 600,
  showSpinner: false
})

// 前置路由守卫
router.beforeEach((to, from, next) => {
  // 启动进度条
  NProgress.start()

  // 设置title
  if (to.meta.title) {
    document.title = to.meta.title
  } else {
    document.title = 'k8s-dashboard' 
  }

  // 放行
  next()

})

router.afterEach(()=> {
  // 关闭进度条
  NProgress.done()
})



export default router
