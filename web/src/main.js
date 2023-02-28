import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

// 引入ant
import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/antd.dark.css'
import * as Icons from '@ant-design/icons-vue';
 


const app = createApp(App)
// 图标注册全局组件
for (const i in Icons) {
    app.component(i, Icons[i])
}
app.use(Antd).use(store).use(router).mount('#app')
