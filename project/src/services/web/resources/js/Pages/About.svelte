<script>
  import Layout from "../Layouts/Layout.svelte";
  let { content } = $props();
</script>

<Layout>
  <div class="rounded-lg bg-white p-6 shadow-md border-l-4 border-blue-500">
    <h1 class="text-2xl font-bold mb-2">About Us</h1>
    <p class="text-gray-600">{content}</p>
    <p class="mt-4 text-sm italic">
      This is a separate route handled by Echo but rendered by Svelte.
    </p>
  </div>
</Layout>
