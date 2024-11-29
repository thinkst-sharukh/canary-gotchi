<template>
  <button :type="type" class="py-3 px-4 inline-flex items-center gap-x-2 text-sm font-medium rounded-lg border border-transparent  disabled:opacity-50 disabled:pointer-events-none justify-center"
  :class="{
    'bg-blue-600 text-white hover:bg-blue-700 focus:outline-none focus:bg-blue-700': variant === 'primary',
    'bg-[#38D47F] text-white focus:outline-none focus:bg-[#38D47F] hover:bg-[#38D47F]': variant === 'success',
  }"
  >
    <slot />
  </button>
</template>

<script lang="ts" setup>
interface Props { 
  type?: 'submit' | 'reset' | 'button',
  variant?: 'primary' | 'secondary' | 'success'
}

withDefaults(defineProps<Props>(), {
  type: 'button',
  variant: 'primary'
})

</script>