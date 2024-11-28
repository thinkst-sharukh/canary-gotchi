<template>
  <Layout>
    <form @submit.prevent="handleSubmit">
      <div class="flex gap-x-3 p-8 bg-white border rounded-lg shadow-lg shadow-gray-100">
        <div class="w-full space-y-4">
          <Input v-model="name" placeholder="Enter gotchi name" required min="3" max="32"
            :disabled="loading || isExistingGotchi" />
          <Input v-model="domainHash" placeholder="Enter your Domain Hash" required :disabled="loading" />
          <Input v-model="authToken" placeholder="Enter your Auth Token" required :disabled="loading" />
        </div>
      </div>

      <div class="mt-4 flex flex-col items-center justify-center">
        <Button type="submit" variant="success" :disabled="loading || (!name || !domainHash || !authToken)"
          class="w-[50%] text-lg">
          <template v-if="!isExistingGotchi">
            {{ loading ? 'Enrolling...' : 'Enroll Gotchi' }}
          </template>

          <template v-else>
            {{ loading ? 'Updating...' : 'Update Gotchi' }}
          </template>
        </Button>

        <a href="https://docs.canary.tools/guide/getting-started.html" target="_blank" class="text-sm mt-2 underline">
          Documentation</a>
      </div>
    </form>
  </Layout>
</template>

<script setup lang="ts">
import Button from '@/components/Button.vue';
import Input from '@/components/Input.vue';
import Layout from '@/components/Layout.vue';
import { names } from '@/router';
import { getGotchi, updateGotchi, verifyAuthKey } from '@/services/gotchi';
import { useGotchiStore } from '@/stores/gotchi';
import { isAxiosError } from 'axios';
import { onMounted, ref, } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { toast } from 'vue3-toastify';

// TODO:: remove this
// const authToken = ref('eb0126fc8ee5be257e64713d9280e335cc6be4ea8dc9');
// const domainHash = ref('86c2ae5a');
const name = ref('');
const authToken = ref('');
const domainHash = ref('');
const loading = ref(false)
const route = useRoute();
const router = useRouter();
const id = route.query.id; // 2a3c4560-cf11-4cda-bcf5-f4b1e01e7a03
const { setGotchi } = useGotchiStore()
const isExistingGotchi = ref(false)

onMounted(() => {
  if (!id) {
    // Redirect to another route if `id` is not set
    router.push({ name: names.home });
  }

  getGotchiData()
})

const getGotchiData = async () => {
  if (loading.value) {
    return;
  }

  let loadingToastId;
  try {
    toast.remove()
    loadingToastId = toast.loading("Please wait...")
    loading.value = true
    const response = await getGotchi(id as string)

    const gotchi = response.data.data

    if (gotchi.id) {
      isExistingGotchi.value = true
      name.value = gotchi.name
      authToken.value = gotchi.auth_token
      domainHash.value = gotchi.hash

      setGotchi(response.data.data)
    } else {
      isExistingGotchi.value = false
    }
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
  } catch (_) {
    isExistingGotchi.value = false
  } finally {
    loading.value = false
    toast.remove(loadingToastId)
  }
}

const verifyAuth = async () => {
  if (loading.value) {
    return;
  }

  if (!id) {
    alert("Invalid ID")
    return;
  }

  let loadingToastId;

  try {
    toast.remove()
    loadingToastId = toast.loading("Enrolling...")
    loading.value = true
    const response = await verifyAuthKey({
      hash: domainHash.value,
      id: id as string,
      name: name.value,
      token: authToken.value
    })

    setGotchi(response.data.data)
    router.push({ name: names.sequence })
  } catch (error) {
    if (isAxiosError(error)) {
      toast.error(error?.response?.data.error)
    } else {
      toast.error("Failed to create enrollment")
    }
  } finally {
    loading.value = false
    toast.remove(loadingToastId)
  }
}

const updateGotchiHandler = async () => {
  if (loading.value) {
    return;
  }

  if (!id) {
    alert("Invalid ID")
    return;
  }

  let loadingToastId;

  try {
    toast.remove()
    loadingToastId = toast.loading("Updating gotchi...")
    loading.value = true
    const response = await updateGotchi({
      hash: domainHash.value,
      id: id as string,
      token: authToken.value
    })

    setGotchi(response.data.data)
    router.push({ name: names.sequence })
  } catch (error) {
    if (isAxiosError(error)) {
      toast.error(error?.response?.data.error)
    } else {
      toast.error("Failed to update gotchi")
    }
  } finally {
    loading.value = false
    toast.remove(loadingToastId)
  }
}

const handleSubmit = () => {
  if (isExistingGotchi.value) {
    updateGotchiHandler();
  } else {
    verifyAuth();
  }
}

</script>