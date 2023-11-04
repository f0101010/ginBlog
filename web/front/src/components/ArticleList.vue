<template>
    <v-col>
        <v-card class="ma-3" v-for="item in artList" :key="item.id" link @click="$router.push(`article/detail/${item.ID}`)">
            <v-row no-gutters>
                <v-col class="d-fiex justify-center aligin-center mx-3" cols="1">
                    <v-img max-height="100" max-width="100" :src="item.img"></v-img>
                </v-col>
                <v-col>
                    <v-card-title class="my-2">
                        <v-chip color="pink" label class="text-white">{{ item.Category.name }}</v-chip>
                        {{ item.title }}
                    </v-card-title>
                    <v-card-subtitle v-text="item.desc"></v-card-subtitle>
                    <v-divider></v-divider>
                    <v-card-text>
                        <v-icon icon="mdi-calendar-month" />
                        <span>{{ this.$moment(item.CreatedAt).format('YYYY-MM-DD HH:MM') }}</span>
                    </v-card-text>
                </v-col>
            </v-row>
        </v-card>
        <v-pagination v-model="queryParam.pagenum" :length="Math.ceil(total / queryParam.pagesize)"
            @click="getArtList()"></v-pagination>
    </v-col>
</template>

<script>
export default {
    data() {
        return {
            artList: [],
            queryParam: {
                pagesize: 5,
                pagenum: 1
            },
            total: 0,
        }
    },
    created() {
        this.getArtList()
    },
    methods: {
        // 获取文章列表
        async getArtList() {
            const { data: res } = await this.$http.get('article', {
                params: {
                    pagesize: this.queryParam.pagesize,
                    pagenum: this.queryParam.pagenum
                },
            })
            this.artList = res.data
            this.total = res.total
        },

    }
}
</script>

<style></style>