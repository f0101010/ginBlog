<template>
    <a-layout class="container">
        <a-layout-sider>
            <div class="log">
                <span>My Blog</span>
            </div>
            <NavBar></NavBar>
        </a-layout-sider>
        <a-layout>
            <a-layout-header class="headerBtn">
                <a-button type="primary" danger ghost @click="loginOut">退出</a-button>
            </a-layout-header>
            <a-layout-content>
                <RouterView></RouterView>
            </a-layout-content>
            <a-layout-footer>
                <FooterBar></FooterBar>
            </a-layout-footer>
        </a-layout>
    </a-layout>
</template>


<script>
import NavBar from '../components/admin/NavBar.vue'
import FooterBar from '../components/admin/FooterBar.vue'

export default {
    components: {
        NavBar,
        FooterBar
    },
    methods: {
        loginOut() {
            window.sessionStorage.clear("token")
            this.$router.push('/login')
        }
    }
}

</script>


<style scoped>
.container {
    height: 100%;
}

.headerBtn {
    display: flex;
    justify-content: flex-end;
    align-items: center;
}

.log {
    height: 32px;
    margin: 16px;
    background-color: #fff;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 17px;
}
</style>