<template>
    <v-card class="mx-auto" max-width="300">
        <v-img cover :src="profileInfo.img">
            <v-card-title>
                <v-col align="center">
                    <v-avatar size="130" color="grey">
                        <v-img cover :src="profileInfo.avatar"></v-img>
                    </v-avatar>
                    <div class="ma-4 text-white">{{ profileInfo.name }}</div>
                </v-col>
            </v-card-title>
            <v-divider></v-divider>
        </v-img>

        <v-col>
            <div class="ma-3 font-italic">About Me:</div>
            <div class="ma-3 text-grey font-italic">{{ profileInfo.desc }}</div>
        </v-col>

        <v-divider color="indigo"></v-divider>

        <v-list nav dense>
            <v-list-item>
                <template v-slot:prepend>
                    <v-icon color="#03A9F4" icon="mdi-qqchat" />
                </template>
                <v-list-item-title class="text-grey">{{ profileInfo.qq_chat }}</v-list-item-title>
            </v-list-item>

            <v-list-item>
                <template v-slot:prepend>
                    <v-icon color="#4CAF50" icon="mdi-wechat" />
                </template>
                <v-list-item-title class="text-grey">{{ profileInfo.wechat }}</v-list-item-title>
            </v-list-item>

            <v-list-item>
                <template v-slot:prepend>
                    <v-icon color="#F44336" icon="mdi-music-circle" />
                </template>
                <v-list-item-title class="text-grey">{{ profileInfo.music }}</v-list-item-title>
            </v-list-item>

            <v-list-item>
                <template v-slot:prepend>
                    <v-icon color="#2196F3" icon="mdi-television-classic" />
                </template>
                <v-list-item-title class="text-grey">{{ profileInfo.bili }}</v-list-item-title>
            </v-list-item>

            <v-list-item>
                <template v-slot:prepend>
                    <v-icon color="#000000" icon="mdi-email" />
                </template>
                <v-list-item-title class="text-grey">{{ profileInfo.email }}</v-list-item-title>
            </v-list-item>
        </v-list>
    </v-card>
</template>

<script>
export default {
    data() {
        return {
            profileInfo: {
                id: 1,
            },
        }
    },
    created() {
        this.getProfileInfo()
    },
    methods: {
        // 获取个人信息
        async getProfileInfo() {
            const { data: res } = await this.$http.get(`profile/${this.profileInfo.id}`)
            this.profileInfo = res.data
        },
    }
}
</script>

<style></style>