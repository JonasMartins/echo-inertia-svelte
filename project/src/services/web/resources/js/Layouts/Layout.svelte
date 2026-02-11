<script>
  import { Link } from "@inertiajs/svelte";
  let { children } = $props();
</script>

<div class="flex min-h-screen flex-col bg-gray-50">
  <header class="bg-white shadow-sm">
    <nav class="container mx-auto flex items-center justify-between p-4">
      <div class="text-xl font-bold text-blue-600">My App</div>
      <div class="space-x-4">
        <Link href="/" class="text-gray-600 hover:text-blue-600">Home</Link>
        <Link href="/about" class="text-gray-600 hover:text-blue-600"
          >About</Link
        >
      </div>
    </nav>
  </header>

  <main class="container mx-auto flex-grow p-8">
    {@render children()}
  </main>

  <footer class="bg-white border-t p-4 text-center text-sm text-gray-500">
    &copy; 2026 Echo + Inertia + Svelte
  </footer>
</div>
