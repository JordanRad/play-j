<script setup>
import Menubar from "primevue/menubar";
import { ref } from "vue";
const items = ref([
  {
    label: "Search",
    icon: "pi pi-fw pi-search",
    to: "search",
  },
  {
    label: "Home",
    icon: "pi pi-fw pi-home",
    to: "/",
  },
  {
    label: "Library",
    icon: "pi pi-fw pi-book",
    to: "library",
  },

  
//   {
//     label: "Settings",
//     to: "search",
//     items: [
//       {
//         label: "My profile",
//         icon: "pi pi-fw pi-user",
//       },
//       {
//         label: "Logout",
//         icon: "pi pi-fw pi-power-off",
//       },
//     ],
//   },
]);
</script>

<template>
  <div class="navigation">
    <Menubar class="bg-black-alpha-90" :model="items"> </Menubar>
  </div>
</template>
