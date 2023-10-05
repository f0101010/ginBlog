<template>
  <a-config-provider :locale="locale">
    <router-view />
  </a-config-provider>
</template>

<style>
html,
body,
#app {
  height: 100%;
  width: 100%;
  margin: 0;
  padding: 0;
}

.ant-card {
  margin: 10px;
  background-color: white;
}

.ant-table {
  margin-top: 15px;
}
</style>


<script>
import zhCN from 'ant-design-vue/es/locale/zh_CN';
import dayjs from 'dayjs';
import 'dayjs/locale/zh-cn';
dayjs.locale('zh-cn');

export default {
  data() {
    return {
      locale: zhCN,
    };
  },
};
</script>
