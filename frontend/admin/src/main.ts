import { createApp } from 'vue';
import { createPinia } from 'pinia';
import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/reset.css';
import VxeTable from 'vxe-table';
import 'vxe-table/lib/style.css';
import App from './App.vue';
import router from './router';
import './assets/main.css';

const app = createApp(App);

app.use(createPinia());
app.use(router);
app.use(Antd);
app.use(VxeTable);

app.mount('#app');
