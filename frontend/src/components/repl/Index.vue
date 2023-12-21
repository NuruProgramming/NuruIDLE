<template>
  <div>
    <div v-for="entry in entries">
      <div>>>> {{ entry.input }}</div>
      <div>
        {{ entry.response }}
      </div>
      <div v-if="entry.error" class="p-2">
        <el-alert
          title="Kosa"
          :description="entry.error"
          type="error"
          :closable="false" />
      </div>
    </div>
    <el-input
      v-model="input"
      @keyup.enter="sendToEngine"
      placeholder="Andika msimbo"></el-input>
  </div>
</template>

<script>
import { Start } from "../../../wailsjs/go/main/App";

export default {
  data() {
    return {
      input: "",
      entries: [],
    };
  },
  methods: {
    sendToEngine() {
      Start(this.input).then((resp) => {
        if (resp.Errors) {
          this.createNewEntry(this.input, "", resp.Errors[0]);
        } else {
          this.createNewEntry(this.input, resp.Evaluated, "");
        }
        this.input = "";
      });
    },
    createNewEntry(item, respone, error) {
      this.entries.push({
        input: item,
        response: respone,
        error: error,
      });
    },
  },
};
</script>
