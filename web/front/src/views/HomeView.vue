<template>
  <v-app>
    <TopBar></TopBar>

    <v-main>
      <v-container>
        <v-row>
          <v-col cols="3">
            <NavBar></NavBar>
          </v-col>
          <v-col>
            <v-sheet max-width="55vw" min-height="80vh" rounded="lg">
              <router-view></router-view>
            </v-sheet>
          </v-col>
        </v-row>
      </v-container>
    </v-main>

    <FooterContent></FooterContent>
  </v-app>
</template>

<script>
import TopBar from '../components/TopBar.vue'
import FooterContent from '../components/FooterContent.vue'
import NavBar from '../components/NavBar.vue'

export default {
  components: {
    TopBar,
    FooterContent,
    NavBar,
  }
}
</script>

<style scoped>
.v-main {
  background-color: #EEEEEE;
}
</style>
