<template>
  <div>
    <div v-for="entry in entries">
      <div>>>> {{ entry.input }}</div>
      <div>
        {{ entry.response }}
      </div>
    </div>
    <el-input v-model="input" @keyup.enter="sendToEngine"></el-input>
  </div>
</template>

<script>
import { Start } from "../../../wailsjs/go/main/App";

export default {
  data() {
    return {
      input: "",
      entries: [
        {
          input: "x + 2",
          response: " 2",
        },
      ],
    };
  },
  methods: {
    sendToEngine() {
      console.log("im there");

      Start(this.input)
        .then((resp, other) => {
          console.log("im here");
          this.createNewEntry(this.input, resp);
          console.log(resp);
          console.log("other", other);
        })
        .catch((error) => {
          console.log(error);

          this.createNewEntry(this.input, error);
          console.log("im failed");
        });
      console.log("im done");
    },
    createNewEntry(item, respone) {
      this.entries.push({
        input: item,
        response: respone,
      });
    },
  },
};
</script>
