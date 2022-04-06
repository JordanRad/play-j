<script setup>
import Card from "primevue/card";
import SplitButton from "primevue/splitbutton";
import InputText from "primevue/inputtext";
import { ref, defineProps } from "vue";
const props = defineProps({
  name: String,
  trackIDs: Int32Array | null,
});

const items = ref([
  {
    label: "Rename playlist",
    icon: "pi pi-pencil",
  },
  {
    label: "Delete playlist",
    icon: "pi pi-trash",
  },
]);
</script>
<template>
  <Card>
    <template #title>
      <div class="text-2xl">{{ props.name }}</div>
    </template>
    <template #content>
      <div class="flex justify-content-between align-items-center">
        <div class="text-lg text-500">{{ props.trackIDs.length }} Songs</div>
        <div>
          <SplitButton :model="items" label="View" />
        </div>
      </div>
    </template>
  </Card>
</template>
