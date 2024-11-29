<script setup lang="ts">
import Logo from './Logo.vue';
</script>

<template>
  <main class="h-full">
    <!-- Hero -->
    <div class="overflow-hidden size-full flex items-center">
      <div class="max-w-[85rem] mx-auto px-4 sm:px-6 lg:px-8 py-10 sm:py-24 w-full">
        <div class="text-center">

          <div class="flex justify-center">
            <Logo class="lg:w-[50%] md:w-[60%]" />
          </div>

          <div class="sm:mt-12 mx-auto max-w-2xl relative z-10">
            <slot />

            <div class="absolute top-0 right-0 -translate-y-24 translate-x-20 -z-10">
              <img src="../assets/game-bg-corner.png" alt="logo" class="size-40" />
            </div>
            <div class="absolute bottom-0 left-0 -translate-x-24 -translate-y-6 -z-10">
              <img src="../assets/heart-bg-corner.png" alt="logo" class="size-40" />
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- End Hero -->
  </main>
</template>