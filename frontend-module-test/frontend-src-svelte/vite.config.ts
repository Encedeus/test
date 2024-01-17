import { defineConfig } from "vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [svelte()],
  build: {
    manifest: "manifest.json",
    rollupOptions: {
      input: {
        main: "src/main.ts",
      },
    },
  },
})
