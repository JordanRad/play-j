<script setup>
import Navigation from "../components/Navigation.vue";
import Card from "primevue/card";
import Button from "primevue/button";
import Badge from "primevue/badge";
import Menu from "primevue/menu";
import ProgressSpinner from "primevue/progressspinner";
import SplitButton from "primevue/splitbutton";
import Divider from "primevue/divider";
import PlaylistService from '../services/PlaylistService'
import { onMounted, ref } from "vue";
const isLoading = ref(true);
const playlists = ref(null)
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

// onMounted(()=>{
//   data = PlaylistService.getAllPlaylists(4)
//   console.log(data)
// })
</script>
<template>
  <div class="bg-black-alpha-90 h-screen">
    <navigation />
    <div class="flex flex-column m-3">
      <div class="p-component mb-5 flex w-full justify-content-between">
        <div class="text-3xl text-white">My Library</div>
        <Button class="shadow-5" icon="pi pi-plus" label="Add playlist" />
      </div>
      <div class="flex flex-col">
        <div class="flex w-full justify-content-center" v-if="!isLoading">
           <ProgressSpinner />
        </div>
         <div v-else class="w-12 md:w-4 shadow-5 border-1 border-white-alpha-10">
          <Card>
            <template #title>
              <div class="text-2xl">Winter 2022</div>
            </template>
            <template #content>
              <div class="flex justify-content-between align-items-center">
                <div class="text-lg text-500">14 Songs</div>
                <div>
                  <SplitButton :model="items" label="View" />
                </div>
              </div>
            </template>
          </Card>
         </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.p-card .p-card-footer {
  display: flex !important;
  width: 100px !important;
  justify-content: center !important;
}
</style>
