<script>
  import Layout from "../Layouts/Layout.svelte";
  let { message } = $props();
</script>

<Layout>
  <div class="rounded-lg bg-white p-6 shadow-md">
    <h1 class="text-2xl font-bold mb-2">Home Page</h1>
    <p class="text-gray-600">{message}</p>
  </div>
</Layout>
