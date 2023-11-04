<template>
    <div>
        <v-app-bar app color="#8B658B">
            <v-avatar class="mx-12" size="48" color="grey">
                <v-img cover src="../assets/02.png"></v-img>
            </v-avatar>
            <v-container class="py-0 fill-height justify-center">
                <v-btn @click="$router.push('/')">首页</v-btn>
                <v-btn v-for="item in cateList" :key="item.id" @click="gotoCate(item.id)">{{ item.name }}</v-btn>
            </v-container>
            <v-spacer></v-spacer>
            <v-responsive max-width="260" color="white">
                <v-text-field dense flat hide-details solo-inverted rounded dark clearable v-model="searchName"
                    @change="searchTitle(searchName)"></v-text-field>
            </v-responsive>
        </v-app-bar>
    </div>
</template>

<script>
export default {
    data() {
        return {
            cateList: [],
            searchName: '',
        }
    },
    created() {
        this.GetCateList();
    },
    methods: {
        // 获取分类
        async GetCateList() {
            const { data: res } = await this.$http.get('category')
            this.cateList = res.data;
        },

        // 搜索文章标题
        searchTitle(title) {
            this.$router.push(`/search/${title}`)
        },

        // 跳转到分类
        gotoCate(cid) {
            this.$router.push(`/category/${cid}`).catch((err) => err)
            console.log(cid)
        }
    }
}
</script>


<style scoped></style>