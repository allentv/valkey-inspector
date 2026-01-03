<script setup lang="ts">
import { computed } from 'vue';

const props = defineProps<{
  name: string
  type: string
  ttl?: number
}>()

const typeColors: Record<string, string> = {
  string: 'bg-blue-100 text-blue-800 dark:bg-blue-900 dark:text-blue-200',
  hash: 'bg-green-100 text-green-800 dark:bg-green-900 dark:text-green-200',
  list: 'bg-orange-100 text-orange-800 dark:bg-orange-900 dark:text-orange-200',
  set: 'bg-purple-100 text-purple-800 dark:bg-purple-900 dark:text-purple-200',
  zset: 'bg-red-100 text-red-800 dark:bg-red-900 dark:text-red-200',
  stream: 'bg-gray-100 text-gray-800 dark:bg-gray-700 dark:text-gray-300',
}

const badgeClass = computed(() => {
  return typeColors[props.type.toLowerCase()] || 'bg-gray-100 text-gray-800'
})

const ttlClass = computed(() => {
  if (props.ttl === -1) return 'bg-green-500' // Persistent
  if (props.ttl !== undefined && props.ttl >= 0) return 'bg-yellow-500' // Volatile
  return 'bg-gray-300' // Unknown
})

const typeLabel = computed(() => props.type.toUpperCase().slice(0, 3))
</script>

<template>
  <div class="flex items-center justify-between p-3 border-b border-gray-200 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-700 cursor-pointer transition-colors group">
    <div class="flex items-center gap-2 overflow-hidden">
      <!-- Type Badge -->
      <span :class="['px-2 py-0.5 text-xs font-bold rounded min-w-[3rem] text-center', badgeClass]">
        {{ typeLabel }}
      </span>
      <!-- Key Name -->
      <span class="truncate font-mono text-sm text-gray-700 dark:text-gray-300 group-hover:text-gray-900 dark:group-hover:text-white" :title="name">{{ name }}</span>
    </div>
    <!-- TTL Indicator -->
    <div :class="['w-2 h-2 rounded-full flex-shrink-0', ttlClass]" :title="ttl === -1 ? 'Persistent' : 'Volatile'"></div>
  </div>
</template>