<template>
  <w-app>
    <router-view></router-view>
  </w-app>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'App',
})
</script>

<style lang="scss">
* {
  font-family: arial;
}
</style>
