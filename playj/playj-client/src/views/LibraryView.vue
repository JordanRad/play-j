<script setup>
import Navigation from "../components/Navigation.vue";
import Card from "primevue/card";
import Button from "primevue/button";
import Badge from "primevue/badge";
import Menu from "primevue/menu";
import ProgressSpinner from "primevue/progressspinner";
</script>
<template>
  <div class="surface-100 h-screen">
    <navigation />
    <div class="flex flex-column m-3">
      <div class="p-component mb-5 flex w-full justify-content-between">
        <div class="text-3xl text-white">My Library</div>
        <Button icon="pi pi-plus" label="Create playlist" />
      </div>
      <div class="flex">
        <!-- <ProgressSpinner /> -->
        <Card class="w-12 md:w-4">
          <template #title> Playlist name </template>
          <template #content> 14 Songs </template>
          <template #footer>
            <div class="flex justify-content-between flex-wrap">
              <Button icon="pi pi-list" label="View" />
              <Button
                icon="pi pi-times"
                label="Delete"
                class="p-button-danger"
                style="margin-left: 0.5em"
              />
            </div>
          </template>
        </Card>
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
