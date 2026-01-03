<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  value: string
}>()

const formattedValue = computed(() => {
  try {
    return JSON.stringify(JSON.parse(props.value), null, 2)
  } catch (e) {
    return props.value
  }
})
</script>

<template>
  <div
    class="bg-white p-4 rounded-md border border-gray-200 font-mono text-sm overflow-x-auto whitespace-pre-wrap text-gray-800 shadow-sm"
  >
    {{ formattedValue }}
  </div>
</template>
