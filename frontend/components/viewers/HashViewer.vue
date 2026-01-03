<script setup lang="ts">
import { ref, computed } from 'vue'

const props = defineProps<{
  data: Record<string, string>
}>()

const searchTerm = ref('')

const filteredData = computed(() => {
  if (!searchTerm.value) return props.data
  const term = searchTerm.value.toLowerCase()
  return Object.fromEntries(
    Object.entries(props.data).filter(([k, v]) => k.toLowerCase().includes(term) || v.toLowerCase().includes(term)),
  )
})
</script>

<template>
  <div class="space-y-4">
    <input
      v-model="searchTerm"
      type="text"
      placeholder="Search fields..."
      class="w-full px-3 py-2 border border-gray-300 rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
    />
    <div class="border border-gray-200 rounded-md overflow-hidden shadow-sm">
      <table class="min-w-full divide-y divide-gray-200">
        <thead class="bg-gray-50">
          <tr>
            <th
              scope="col"
              class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider w-1/3"
            >
              Field
            </th>
            <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
              Value
            </th>
          </tr>
        </thead>
        <tbody class="bg-white divide-y divide-gray-200">
          <tr v-for="(val, key) in filteredData" :key="key">
            <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900 font-mono">{{ key }}</td>
            <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500 font-mono">{{ val }}</td>
          </tr>
          <tr v-if="Object.keys(filteredData).length === 0">
            <td colspan="2" class="px-6 py-4 text-center text-sm text-gray-500">No fields match your search.</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>
