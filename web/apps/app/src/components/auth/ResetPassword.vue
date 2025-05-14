<script setup lang="ts">
withDefaults(defineProps<{
  email?:string
}>(), {
  email: '123'
}) 
</script>

<template>
<div>{{email}}</div>
</template>

<style scoped>

</style>