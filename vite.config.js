// vite.config.js
import { svelte } from "@sveltejs/vite-plugin-svelte";
import tailwindcss from "@tailwindcss/vite";
import { defineConfig } from "vite";

export default defineConfig({
  plugins: [tailwindcss(), svelte()],
  build: {
    manifest: true,
    outDir: "public/build",
    rollupOptions: {
      input: "resources/js/app.js",
    },
  },
});
