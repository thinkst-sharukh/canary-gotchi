<script lang="ts" setup>
import {  defineProps, computed } from 'vue';

interface Props {
  modelValue: string,
  label?: string
  id?: string
  type?: string
  placeholder?: string
  bordered?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  id: `input-${Math.random().toString(36).substring(2, 9)}`,
  type: 'text',
  bordered: true
});

const emit = defineEmits(['update:modelValue']);
const internalValue = computed({
  get: () => props.modelValue,
  set: (value) => emit('update:modelValue', value),
});
</script>

<template>
  <div>
    <label :for="id" class="inline-block text-sm font-medium mb-2 " v-if="label">
      {{ label }}
    </label>
    <input
      :id="id"
      :type="type"
      v-model="internalValue"
      :placeholder="placeholder"
      class="py-3 px-4 block w-full  rounded-lg text-sm focus:border-blue-500 focus:ring-blue-500 disabled:opacity-50 disabled:pointer-events-none"
      :class="{'border-gray-200 border': bordered}"
      v-bind="$attrs"
    />
  </div>
</template>

<style scoped>

</style>
