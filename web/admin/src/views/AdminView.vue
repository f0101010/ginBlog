<template>
    <a-layout class="container" has-sider>
        <NavBar></NavBar>
        <a-layout :style="{ marginLeft: '200px' }">
            <a-layout-header class="headerBtn">
                <HeaderTop></HeaderTop>
            </a-layout-header>
            <a-layout-content :style="{ margin: '24px 16px 0', overflow: 'initial' }">
                <RouterView :key="$route.path"></RouterView>
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
import HeaderTop from '../components/admin/HeaderTop.vue'

export default {
    components: {
        NavBar,
        FooterBar,
        HeaderTop
    },

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
</style>