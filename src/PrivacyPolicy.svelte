<script>
  import { onMount } from 'svelte';
  import { currentPath } from './stores.js';

  export let location;

  onMount(() => currentPath.set('/privacy'));
</script>

<style>
  .page {
    padding: 48px 48px 80px;
    max-width: 680px;
  }

  .page-header {
    margin-bottom: 32px;
    padding-bottom: 20px;
    border-bottom: 1px solid var(--border);
  }

  @media (max-width: 900px) { .page { padding: 32px 24px 64px; } }
  @media (max-width: 640px) { .page { padding: 24px 16px 48px; } }
</style>

<div class="page">
  <div class="page-header">
    <h1>Privacy Policy</h1>
  </div>

  <p>Run A Relay does not set cookies on your device.</p>

  <p>
    Only severe TCP, HTTP, and TLS errors are logged. These logs may
    inadvertently contain visitor IP addresses.
  </p>

  <p>
    SSH connections to port 22 on RunARelay.org are not logged at all.
    If we begin to receive abuse, this may change.
  </p>

  <p>
    In the future we may record <em>how many</em> visitors we receive,
    but not <em>who</em>. We will never use Google Analytics.
  </p>

  <p>Aside from the caveats above, we store zero information about our users.</p>
</div>
