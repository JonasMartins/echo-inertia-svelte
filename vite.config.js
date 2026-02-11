// vite.config.js
import { svelte } from "@sveltejs/vite-plugin-svelte";
import tailwindcss from "@tailwindcss/vite";
import path from "path";
import { fileURLToPath } from "url";
import { defineConfig } from "vite";

// Get the directory name of the current file (project root)
const __dirname = path.dirname(fileURLToPath(import.meta.url));

export default defineConfig({
  // Set the root to the sub-directory where your frontend files live
  root: path.resolve(__dirname, "project/src/services/web"),
  plugins: [tailwindcss(), svelte()],
  build: {
    manifest: true,
    // outDir is relative to 'root' above
    outDir: path.resolve(__dirname, "project/src/services/web/public/build"),
    // emptyOutDir is needed because outDir is now outside the 'root'
    emptyOutDir: true,
    rollupOptions: {
      // Input must be relative to the 'root' or absolute
      input: path.resolve(
        __dirname,
        "project/src/services/web/resources/js/app.js",
      ),
    },
    server: {
      // Ensure the dev server is accessible
      origin: "http://localhost:5173",
    },
  },
});
