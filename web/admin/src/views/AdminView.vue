<template>
    <a-layout class="container">
        <NavBar></NavBar>

        <a-layout>
            <a-layout-header class="headerBtn">
                <HeaderTop></HeaderTop>
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