<script setup>
import Navigation from "../components/Navigation.vue";
import Button from "primevue/button";
import PlayerService from "@/services/PlayerService";
import { onMounted, ref } from "vue";

const file = ref(null);
const isPlaying = ref(false);
// Audio
var context;
var audioBuffer;
function initAudio() {
  if (!window.AudioContext) {
    if (!window.webkitAudioContext) {
      alert("your browser does not support audio play");
    }
    window.AudioContext = window.webkitAudioContext;
  }

  context = new AudioContext();
  console.log("context int");
}

async function play() {
  isPlaying.value = true;
  if (file !== null) {
    context = new AudioContext();
    let utfEncoder = new TextEncoder();
    const arr = utfEncoder.encode(file.value);
    console.log(arr);

    let audio;

    try {
      audio = await context.decodeAudioData(arr.buffer);
    } catch (error) {
      console.log(error);
    }
    
    console.log(audio);
    const source = context.createBufferSource();
    source.buffer = audio;
    source.connect(context.destination);
    source.start(0);
  }
}
function pause() {
  isPlaying.value = false;
}
onMounted(async () => {
  let musicFileResponse = await PlayerService.getTrackById(1);
  if (musicFileResponse !== null) {
    file.value = musicFileResponse.file;
    //console.log(file.value)
  }
});
</script>

<template>
  <div class="bg-black-alpha-90 h-screen">
    <navigation />
    <div class="flex m-4">
      <div class="p-component text-2xl text-white">Discover</div>
      <div>
        <Button
          v-if="!isPlaying"
          @click="play"
          class="shadow-5"
          icon="pi pi-play"
        />
        <Button v-else @click="pause" class="shadow-5" icon="pi pi-pause" />
      </div>
    </div>
  </div>
</template>
