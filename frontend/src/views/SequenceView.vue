<template>
  <Layout>
    <div class="gap-x-3 p-8 bg-white border rounded-lg shadow-lg shadow-gray-100 flex flex-col space-y-4">
      <h2 class="md:text-xl font-bold">Here is your sequence</h2>
      <div class="grid md:flex gap-3 items-center justify-center">
        <span v-for="(sq, i) in sequence?.sequence.split(',')" :key="i"
          class="uppercase border p-4 rounded-lg flex items-center gap-x-2 font-medium">
          {{ sq }}
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
            stroke="currentColor" class="size-5" :class="{
              'rotate-180': sq === 'up',
              'rotate-90': sq === 'left',
              '-rotate-90': sq === 'right',
            }">
            <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 13.5 12 21m0 0-7.5-7.5M12 21V3" />
          </svg>
        </span>
      </div>

      <form @submit.prevent="verifySequence">
        <Input placeholder="Enter your sequence" v-model="enteredSequence" />
      </form>
    </div>

    <div class="mt-4">
      Sequence valid for:

      <div class="text-4xl font-semibold">
        <p v-if="remainingSeconds > 0">
          <span v-if="min > 0">{{ min }}m</span>
          <span v-if="sec > 0" :class="{
            'text-yellow-500': sec < 30 && sec > 10 && min <= 1,
            'text-red-500 animate-pulse': sec <= 10 && min <= 1,
          }">{{ sec }}s</span>
        </p>
        <div v-else class="space-y-4">
          <p class="text-red-500">
            Expired
          </p>
        </div>
      </div>
    </div>
  </Layout>
</template>

<script lang="ts" setup>
import Input from '@/components/Input.vue';
import Layout from '@/components/Layout.vue';
import { names } from '@/router';
import { regenerateSequence, validateSequence } from '@/services/sequence';
import { useGotchiStore } from '@/stores/gotchi';
import { isAxiosError } from 'axios';
import { computed, nextTick, onBeforeUnmount, onMounted, ref, } from 'vue';
import { useRouter } from 'vue-router';
import { toast } from 'vue3-toastify';

const { gotchi } = useGotchiStore()
const remainingSeconds = ref(0)
const timer = ref()
const router = useRouter()
const enteredSequence = ref('')
const sequence = ref(gotchi?.sequence)

const min = computed(() => Math.floor(remainingSeconds.value / 60))
const sec = computed(() => remainingSeconds.value % 60)

const calculateRemainingSeconds = () => {
  const currentTime = new Date(); // Current time in local timezone
  const targetTime = new Date(sequence.value?.expires || ''); // Parse the future UTC ISO string to a Date object

  // Ensure the target time is in UTC and current time is compared in the same format
  const diffInMs = targetTime.getTime() - currentTime.getTime(); // Difference in milliseconds

  if (diffInMs > 0) {
    remainingSeconds.value = Math.floor(diffInMs / 1000); // Convert to seconds
  } else {
    remainingSeconds.value = 0; // If the current time is already past the target date, set the timer to 0
  }
}

const startTimer = () => {
  calculateRemainingSeconds();

  nextTick(() => {
    if (remainingSeconds.value > 0) {
      timer.value = setInterval(() => {
        remainingSeconds.value -= 1;
        if (remainingSeconds.value <= 0) {
          clearInterval(timer.value); // Stop the timer when expired
          regenerate()
        }
      }, 1000);
    }
  })
}

const regenerate = async () => {
  if (!gotchi?.id) {
    return;
  }

  let loadingToastId;
  try {
    toast.remove()
    loadingToastId = toast.loading("Regenerating sequence...")
    const response = await regenerateSequence({ id: gotchi.id })

    sequence.value = response.data.data;
    toast.success("Sequence regenerated")
    startTimer()
  } catch (error) {
    if (isAxiosError(error)) {
      toast.error(error?.response?.data.error)
    } else {
      toast.error("Failed to regenerate sequence")
    }
  } finally {
    toast.remove(loadingToastId)
  }
}

const verifySequence = async () => {
  if (!gotchi?.id) {
    return;
  }

  let loadingToastId;
  try {
    toast.remove()
    loadingToastId = toast.loading("Verifying sequence...")
    await validateSequence({
      id: gotchi.id,
      sequence: enteredSequence.value
    })

    router.push({ name: names.verifiedSequence })
  } catch (error) {
    if (isAxiosError(error)) {
      toast.error(error?.response?.data.error)
    } else {
      toast.error("Failed to verify sequence")
    }
  } finally {
    toast.remove(loadingToastId)
  }
}

onMounted(() => {
  if (gotchi?.id) {
    startTimer()
  } else {
    router.push({ name: names.home })
  }
})

onBeforeUnmount(() => {
  if (timer.value) {
    clearInterval(timer.value); // Cleanup on component destroy
  }
})

</script>