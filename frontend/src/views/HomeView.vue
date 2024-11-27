<script setup lang="ts">
import Button from '@/components/Button.vue';
import Input from '@/components/Input.vue';
import axios from 'axios';
import { onMounted, ref } from 'vue';

// TODO:: remove this
// const authToken = ref('eb0126fc8ee5be257e64713d9280e335cc6be4ea8dc9');
// const domainHash = ref('86c2ae5a');
const authToken = ref('');
const domainHash = ref('');
const loading = ref(false)

const fetchPing = async () => {
  if (loading.value) {
    return;
  }

  try {
    loading.value = true
    // Send the GET request
    const response = await axios.post('/api/verify-api-key', {
      token: authToken.value,
      hash: domainHash.value
    })

    console.log(response.status);
    if (response.status === 200) {
      alert("Verified")
    } else {
      alert("Failed to verify")
    }
  } catch (error) {
    console.error('Error fetching data:', error);
    alert("Failed to verify")
  } finally {
    loading.value = false
  }
}

</script>

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

          <div class="mt-7 sm:mt-12 mx-auto max-w-xl relative">
            <!-- Form -->
            <form @submit.prevent="fetchPing" class="relative">
              <div class="relative z-10 flex gap-x-3 p-4 bg-white border rounded-lg shadow-lg shadow-gray-100">
                <div class="w-full space-y-4">
                  <Input v-model="domainHash" placeholder="Enter your Domain Hash" required />
                  <Input v-model="authToken" placeholder="Enter your Auth Token" required />
                </div>
              </div>

              <div class="mt-4 flex flex-col items-center justify-center">
                <Button type="submit" :disabled="loading || (!domainHash || !authToken)">
                  Submit
                </Button>

                <a href="https://docs.canary.tools/guide/getting-started.html" target="_blank" class="text-sm mt-2 underline">
                  Documentation</a>
              </div>

              <div class="hidden md:block absolute top-0 end-0 -translate-y-12 translate-x-20">
              <svg class="w-16 h-auto text-orange-500" width="121" height="135" viewBox="0 0 121 135" fill="none"
                xmlns="http://www.w3.org/2000/svg">
                <path d="M5 16.4754C11.7688 27.4499 21.2452 57.3224 5 89.0164" stroke="currentColor" stroke-width="10"
                  stroke-linecap="round" />
                <path d="M33.6761 112.104C44.6984 98.1239 74.2618 57.6776 83.4821 5" stroke="currentColor"
                  stroke-width="10" stroke-linecap="round" />
                <path d="M50.5525 130C68.2064 127.495 110.731 117.541 116 78.0874" stroke="currentColor"
                  stroke-width="10" stroke-linecap="round" />
              </svg>
            </div>
            <div class="hidden md:block absolute bottom-0 start-0 -translate-y-12 -translate-x-32">
              <svg class="w-40 h-auto text-cyan-500" width="347" height="188" viewBox="0 0 347 188" fill="none"
                xmlns="http://www.w3.org/2000/svg">
                <path
                  d="M4 82.4591C54.7956 92.8751 30.9771 162.782 68.2065 181.385C112.642 203.59 127.943 78.57 122.161 25.5053C120.504 2.2376 93.4028 -8.11128 89.7468 25.5053C85.8633 61.2125 130.186 199.678 180.982 146.248L214.898 107.02C224.322 95.4118 242.9 79.2851 258.6 107.02C274.299 134.754 299.315 125.589 309.861 117.539L343 93.4426"
                  stroke="currentColor" stroke-width="7" stroke-linecap="round" />
              </svg>
            </div>
            </form>
            <!-- End Form -->
          </div>
        </div>
      </div>
    </div>
    <!-- End Hero -->
  </main>
</template>
