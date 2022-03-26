<script setup>
import Navigation from "../components/Navigation.vue";
import Card from "primevue/card";
import Button from 'primevue/button'
import Badge from 'primevue/badge'
import Menu from 'primevue/menu'
import ProgressSpinner from 'primevue/progressspinner';
</script>
<template>
<div class="surface-100 h-screen"> 
  <navigation />
  <div class="flex flex-column m-3">
    <div class="p-component mb-5 flex w-full justify-content-between">
        <div class="text-3xl text-white">My Library</div>
        <Button icon="pi pi-plus" class="p-button-rounded p-button-lg" />
    </div>
    <div class="flex">
        <!-- <ProgressSpinner /> -->
      <Card>
        <template #title> Playlist name </template>
        <template #content>
          14 Songs
        </template>
        <template #footer>
          <Button icon="pi pi-list" label="View" />
          <Button
            icon="pi pi-times"
            label="Delete"
            class="p-button-secondary"
            style="margin-left: 0.5em"
          />
         
        </template>
      </Card>
    </div>
  </div>
  </div>
</template>
