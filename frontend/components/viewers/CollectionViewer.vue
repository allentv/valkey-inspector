<template>
  <div class="border border-gray-200 rounded-md overflow-hidden shadow-sm">
    <ul class="divide-y divide-gray-200 bg-white">
      <li v-for="(item, index) in data" :key="index" class="px-6 py-4 text-sm font-mono text-gray-700 hover:bg-gray-50">
        <span class="text-gray-400 mr-4 select-none w-8 inline-block text-right">{{ index + 1 }}</span>
        {{ item }}
      </li>
      <li v-if="data.length === 0" class="px-6 py-4 text-center text-sm text-gray-500">Empty collection.</li>
    </ul>
  </div>
</template>
