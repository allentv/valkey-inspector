<template>
  <div class="border border-gray-200 dark:border-gray-700 rounded-md overflow-hidden shadow-sm">
    <ul class="divide-y divide-gray-200 dark:divide-gray-700 bg-white dark:bg-gray-800">
      <li v-for="(item, index) in data" :key="index" class="px-6 py-4 text-sm font-mono text-gray-700 dark:text-gray-300 hover:bg-gray-50 dark:hover:bg-gray-700">
        <span class="text-gray-400 dark:text-gray-500 mr-4 select-none w-8 inline-block text-right">{{ index + 1 }}</span>
        <span class="break-all">{{ item }}</span>
      </li>
      <li v-if="data.length === 0" class="px-6 py-4 text-center text-sm text-gray-500 dark:text-gray-400">Empty collection.</li>
    </ul>
  </div>
</template>
