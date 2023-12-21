<template>
  <div>
    <Markdown :source="source" />
  </div>
</template>

<script>
import Markdown from "vue3-markdown-it";
import MarkdownIt from "markdown-it";
import markdownFile from "@/components/HelloWorld.vue";

export default {
  components: {
    Markdown,
  },
  data() {
    return {
      source: "# heading  ukb",
    };
  },
  created() {
    this.loadMarkdownFile();
  },
  methods: {
    async loadMarkdownFile() {
      try {
        const response = await fetch("./assets/images/logo-universal.png");
        if (!response.ok) {
          this.source = "Failed to fetch the file";
          throw new Error("Failed to fetch the file");
        }
        this.source = await response.text();
        // this.markdownContent = md.render(markdownText);
      } catch (error) {
        console.error("Error loading markdown file:", error);
      }
      //   try {
      //     const markdownFile = await import.meta.globEager("./kama.md");
      //     const markdownContent = markdownFile["./kama.md"].default;
      //     // this.renderedMarkdown = md.render(markdownContent);
      //   } catch (error) {
      //     console.error("Error loading and rendering Markdown:", error);
      //   }
    },
  },
};
</script>
