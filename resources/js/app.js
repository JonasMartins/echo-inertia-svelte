import { createInertiaApp } from "@inertiajs/svelte";
import { mount } from "svelte"; // Import mount from svelte
import "../css/app.css";

createInertiaApp({
  resolve: (name) => {
    const pages = import.meta.glob("./Pages/**/*.svelte", { eager: true });
    return pages[`./Pages/${name}.svelte`];
  },
  setup({ el, App, props }) {
    // Svelte 5 uses the mount function instead of 'new App'
    mount(App, {
      target: el,
      props,
    });
  },
});
