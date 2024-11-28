<template>
  <main class="h-full">
    <!-- Hero -->
    <div class="overflow-hidden size-full flex items-center">
      <div class="max-w-[85rem] mx-auto px-4 sm:px-6 lg:px-8 py-10 sm:py-24 w-full">
        <div class="text-center">
          <h1 class="text-4xl sm:text-6xl font-bold text-primary">
            Canary Gotchi
          </h1>

          <p class="mt-3 text-gray-600">
            Hows your Canary(A better tag line)
          </p>

        </div>
        <div class="mt-7 sm:mt-12 mx-auto border border-gray-200 rounded-md shadow-sm p-8">
          <div class="space-y-2">
            <h2 class="font-medium text-lg">
              Leaderboard
            </h2>
            <p class="text-sm">
              Top Performers
            </p>
          </div>

          <div class="mt-8 border border-gray-200 rounded-lg shadow-sm overflow-hidden">
            <table class="w-full">
              <thead>
                <tr class="divide-x divide-gray-200">
                  <th class="text-left bg-gray-100 font-medium px-4 py-2" v-for="head in headings" :key="head.key">{{
                    head.name }}
                  </th>
                </tr>
              </thead>

              <tbody class="divide-y divide-gray-200 border-t border-gray-200">
                <tr class="divide-x divide-gray-200" v-for="g in gotchis" :key="g.id">
                  <td class="p-4">
                    {{ g.name }}
                  </td>
                  <td class="px-4 py-2">
                    {{ g.level }}
                  </td>
                  <td class="px-4 py-2">
                    {{ g.verified ? 'Yes' : 'No' }}
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
    <!-- End Hero -->
  </main>
</template>

<script lang="ts" setup>
import { getAllGotchi } from '@/services/gotchi';
import type { IGotchi } from '@/types';
import { onMounted, ref } from 'vue';
import { toast } from 'vue3-toastify';

const headings = [
  {
    name: 'Name',
    key: 'name'
  },
  {
    name: 'Level',
    key: 'level'
  },
  {
    name: 'Verified',
    key: 'verified'
  }
]

const gotchis = ref<IGotchi[]>([])
const loading = ref(false)

onMounted(async () => {
  if (loading.value) {
    return;
  }

  let loadingToastId;
  try {
    loading.value = true
    toast.remove()
    loadingToastId = toast.loading("Please wait...")
    const response = await getAllGotchi()

    gotchis.value = response.data.data
    gotchis.value.sort((a, b) => b.level - a.level)

  } catch (error) {
    console.log(error)
  } finally {
    loading.value = false
    toast.remove(loadingToastId)
  }
})
</script>