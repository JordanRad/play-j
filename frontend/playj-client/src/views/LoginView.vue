<script setup>
import Card from "primevue/card";
import Button from "primevue/button";
import InputText from "primevue/inputtext";
import AccountService from "@/services/AccountService.js";
import { ref } from "vue";
import { useRouter } from "vue-router";
const router = useRouter();
const email = ref("");
const password = ref("");
const loginError = ref("");
async function onSubmit() {
  console.log(email.value, password.value);
  let loginResponse = await AccountService.login(email.value, password.value);
  if (loginResponse != null) {
    localStorage.setItem("user", JSON.stringify(loginResponse));
    router.push({ name: "home" });
  }

  loginError.value = "Invalid credentials";
}
</script>
<template>
  <div class="bg-black-alpha-90 h-screen">
    <div class="flex flex-column p-5">
      <div
        class="p-component my-5 flex flex-column w-full justify-content-center align-items-center"
      >
        <div class="text-3xl text-white">Login</div>
      </div>
      <div
        class="mt-2 h-20rem flex flex-column justify-content-center align-items-center"
      >
        <InputText
          autocomplete="off"
          class="md:w-6 lg:w-3 w-12 my-4"
          type="text"
          v-model="email"
          placeholder="Email"
        />

        <InputText
          class="md:w-6 lg:w-3 w-12 mb-4"
          type="password"
          v-model="password"
          placeholder="Password"
        />

        <Button @click="onSubmit" class="md:w-6 lg:w-3 w-12" label="Login" />
      </div>
    </div>
  </div>
</template>
