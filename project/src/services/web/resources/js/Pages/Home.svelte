<script>
  import { Button } from "$lib/components/ui/button/index.js";
  import ArrowUpIcon from "@lucide/svelte/icons/arrow-up";
  import Layout from "../Layouts/Layout.svelte";
  let { message } = $props();
</script>

<Layout>
  <div class="rounded-lg bg-white p-6 shadow-md">
    <h1 class="text-2xl font-bold mb-2">Home Page</h1>
    <p class="text-gray-600">{message}</p>
    <div class="flex flex-wrap items-center gap-2 md:flex-row">
      <Button variant="secondary">Button</Button>
      <Button variant="outline" size="icon" aria-label="Submit">
        <ArrowUpIcon />
      </Button>
    </div>
  </div>
</Layout>
