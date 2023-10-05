<template>
    <a-layout-sider breakpoint="lg">
        <div class="logo">
            <span>My Blog</span>
        </div>
        <a-menu v-model:openKeys="openKeys" v-model:selectedKeys="selectedKeys" mode="inline" :theme="theme" :items="items"
            @click="goTopage" />
    </a-layout-sider>
</template>


<script>

import { h, ref } from 'vue';
import {
    UserOutlined,
    FileOutlined,
    TagsOutlined,
    FormOutlined,
    FrownOutlined,
    SnippetsOutlined,
} from '@ant-design/icons-vue';

export default {
    data() {
        return {
            selectedKeys: ref([]),
            openKeys: ref([]),
            theme: ref('dark'),
            items: ref([
                {
                    key: 'index',
                    icon: () => h(FrownOutlined),
                    label: '仪表盘',
                    title: '仪表盘',
                },

                {
                    key: 'userlist',
                    icon: () => h(UserOutlined),
                    label: '用户列表',
                    title: '用户列表',
                },
                {
                    key: 'catelist',
                    icon: () => h(TagsOutlined),
                    label: '分类列表',
                    title: '分类列表',
                },
                {
                    key: 'sub1',
                    icon: () => h(FileOutlined),
                    label: '文章管理',
                    title: '文章管理',
                    children: [
                        {
                            key: 'addart',
                            icon: () => h(FormOutlined),
                            label: '写文章',
                            title: '写文章',
                        },
                        {
                            key: 'artlist',
                            icon: () => h(SnippetsOutlined),
                            label: '文章列表',
                            title: '文章列表',
                        },
                    ],
                },
            ]),
        };
    },
    methods: {
        goTopage({ key }) {
            this.$router.push('/admin' + `/${key}`);
        },
    }

}
</script>

<style scoped>
.logo {
    height: 32px;
    margin: 16px;
    background-color: #fff;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 17px;
}
</style>
