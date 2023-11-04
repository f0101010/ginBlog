<template>
    <div>
        <div class="d-flex justify-center pa-3 text-h4 font-weihgt-bold">{{ artInfo.title }}</div>
        <div class="d-flex justify-center">
            <v-icon icon="mdi-calendar-month" />
            <span>{{ this.$moment(artInfo.CreatedAt).format('YYYY-MM-DD HH:MM') }}</span>
        </div>
        <v-divider class="pa-3 ma-3"></v-divider>
        <v-alert color="#BA68C8" class="ma-4" dark border="start" variant="outlined">{{ artInfo.desc }}</v-alert>
        <div class="content ma-5 pa-3 text-justify" v-html="artInfo.content"></div>
    </div>
</template>

<script>
export default {
    props: ['id'],
    data() {
        return {
            artInfo: {}
        }
    },
    created() {
        this.getArtInfo()
    },
    methods: {
        // 获取文章详情
        async getArtInfo() {
            const { data: res } = await this.$http.get(`article/info/${this.id}`)
            this.artInfo = res.data
        }
    }
}
</script>


<style scoped>
.content {
    p {
        width: 90%;
    }

    img {
        width: 90%;
    }

    span {
        width: 90%;
    }
}
</style>