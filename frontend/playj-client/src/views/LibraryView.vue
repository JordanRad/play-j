<script setup>
import Navigation from "../components/Navigation.vue";
import Card from "primevue/card";
import Button from "primevue/button";
import ProgressSpinner from "primevue/progressspinner";
import SplitButton from "primevue/splitbutton";
import Dialog from "primevue/dialog";
import PlaylistCard from "../components/PlaylistCard.vue";
import PlaylistService from "@/services/PlaylistService";
import { onMounted, ref } from "vue";
const isLoading = ref(true);
const playlists = ref(null);
const displayAddPlaylistDialog = ref(false);
const newPlaylistName = ref("");

function openAddPlaylistDialog() {
  displayAddPlaylistDialog.value = !displayAddPlaylistDialog.value;
}
async function addPlaylist() {
  console.log(newPlaylistName.value);
  let createPlaylistResponse = await PlaylistService.createAccountPlaylist(newPlaylistName.value);
  
  console.log(createPlaylistResponse);
  
  if (createPlaylistResponse !== null) {
    window.location.reload();
  }
}
onMounted(async () => {
  let playlistsResponse = await PlaylistService.getAccountPlaylistCollection();
  if (playlistsResponse !== null) {
    isLoading.value = false;
    playlists.value = playlistsResponse.resources;
  }
});
</script>
<template>
  <div class="bg-black-alpha-90 h-screen">
    <navigation />
    <div class="flex flex-column m-3">
      <div class="p-component mb-5 flex w-full justify-content-between">
        <div class="text-3xl text-white">My Library</div>
        <Button
          @click="openAddPlaylistDialog"
          class="shadow-5"
          icon="pi pi-plus"
          label="Add playlist"
        />
      </div>
      <div class="flex">
        <div class="flex w-full justify-content-center" v-if="isLoading">
          <ProgressSpinner />
        </div>
        <div v-else class="flex h-24rem w-full flex-wrap">
          <Dialog
            class="shadow-5 border-1 border-white-alpha-10"
            header="Add new playlist"
            modal
            v-model:visible="displayAddPlaylistDialog"
          >
            <div>
              <InputText
                autocomplete="off"
                type="text"
                v-model="newPlaylistName"
                placeholder="Playlist name"
              />
            </div>
            <template #footer>
              <Button @click="addPlaylist" label="Add" autofocus />
            </template>
          </Dialog>
          <div
            class="h-4rem md:w-5 w-12 shadow-5 border-1 border-white-alpha-10 mr-auto"
            v-for="(playlist, index) in playlists"
            :key="index"
          >
            <playlist-card
              :name="playlist.name"
              :track-i-ds="playlist.trackIDs"
            />
          </div>
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
