<template>
    <v-card class="mx-auto" max-width="300">
        <v-img cover src="../assets/navbg1.png">
            <v-card-title>
                <v-col align="center">
                    <v-avatar size="130" color="grey">
                        <v-img cover src="../assets/userphoto.png"></v-img>
                    </v-avatar>
                    <div class="ma-4 text-white">weject</div>
                </v-col>
            </v-card-title>
            <v-divider></v-divider>
        </v-img>

        <v-col>
            <div class="ma-3">About Me:</div>
            <div class="ma-3">12321321312</div>
        </v-col>

        <v-divider color="indigo"></v-divider>

        <v-list nav dense>
            <v-list-item>
                <template v-slot:prepend>
                    <v-icon icon="mdi-qqchat" />
                </template>
                <v-list-item-title class="text-grey">1121212</v-list-item-title>
            </v-list-item>

            <v-list-item>
                <template v-slot:prepend>
                    <v-icon icon="mdi-wechat" />
                </template>
                <v-list-item-title class="text-grey">1121212</v-list-item-title>
            </v-list-item>

            <v-list-item>
                <template v-slot:prepend>
                    <v-icon icon="mdi-sina-weibo" />
                </template>
                <v-list-item-title class="text-grey">1121212</v-list-item-title>
            </v-list-item>

            <v-list-item>
                <template v-slot:prepend>
                    <v-icon icon="mdi-facebook" />
                </template>
                <v-list-item-title class="text-grey">1121212</v-list-item-title>
            </v-list-item>

            <v-list-item>
                <template v-slot:prepend>
                    <v-icon icon="mdi-email" />
                </template>
                <v-list-item-title class="text-grey">1121212</v-list-item-title>
            </v-list-item>
        </v-list>
    </v-card>
</template>

<script>
</script>

<style></style>