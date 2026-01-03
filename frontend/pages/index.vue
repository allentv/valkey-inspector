<script setup lang="ts">
const { selectedKey } = useKeys()
</script>

<template>
  <KeyDetail :key-name="selectedKey" />
</template>
