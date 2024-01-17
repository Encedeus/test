import { defineConfig } from "vite";

export default defineConfig({
    build: {
        manifest: "manifest.json",
        rollupOptions: {
            input: {
                main: "src/main.ts",
            },
        },
    },
});