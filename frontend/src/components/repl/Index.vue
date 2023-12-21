<template>
  <div>
    <input v-model="input" />
    <button @click="sendToEngine">Send</button>

    <div>The response is : {{ response }}</div>
  </div>
</template>

<script>
import { Start } from "../../../wailsjs/go/main/App";

export default {
  data() {
    return {
      input: "",
      response: "",
    };
  },
  methods: {
    sendToEngine() {
      console.log("im there");

      Start(this.input)
        .then((resp) => {
          console.log("im here");
          this.response = resp;
          console.log(resp);
        })
        .catch(() => {
          console.log("im failed");
        });
      console.log("im done");
    },
  },
};
</script>
