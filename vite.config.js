// vite.config.js
import { svelte } from "@sveltejs/vite-plugin-svelte";

import tailwindcss from "@tailwindcss/vite";
import path from "path";
import { fileURLToPath } from "url";
import { defineConfig } from "vite";

const __dirname = path.dirname(fileURLToPath(import.meta.url));

export default defineConfig({
  root: path.resolve(__dirname, "project/src/services/web"),
  plugins: [tailwindcss(), svelte()],
  server: { port: 3000 },
  resolve: {
    alias: {
      $lib: path.resolve(__dirname, "project/src/services/web/resources/js/lib"),
    }
  },
  build: {
    manifest: true,
    outDir: path.resolve(__dirname, "project/src/services/web/public/build"),
    emptyOutDir: true,
    rollupOptions: {
      input: path.resolve(__dirname, "project/src/services/web/resources/js/app.js")
    }
  }
});
