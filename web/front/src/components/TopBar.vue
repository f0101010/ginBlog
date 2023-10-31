<template>
    <div>
        <v-app-bar app color="#8B658B">
            <v-container class="py-0 fill-height">
                <v-avatar class="mx-12" size="48" color="grey"></v-avatar>
                <v-btn @click="$router.push('/')">首页</v-btn>
                <v-btn v-for="item in cateList" :key="item.id">{{ item.name }}</v-btn>
            </v-container>
            <v-spacer></v-spacer>
            <v-responsive max-width="260" color="white">
                <v-text-field dense flat hide-details solo-inverted rounded dark clearable></v-text-field>
            </v-responsive>
        </v-app-bar>
    </div>
</template>

<script>
export default {
    data() {
        return {
            cateList: []
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
        }
    }
}
</script>


<style scoped></style>