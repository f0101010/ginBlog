<template>
    <a-button type="primary" danger ghost @click="loginOut">退出</a-button>
</template>

<script>

export default {
    methods: {
        loginOut() {
            window.sessionStorage.clear("token")
            this.$router.push('/login')
        }
    }
}
</script>