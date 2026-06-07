<script>
  import { onMount } from 'svelte';
  import { currentPath } from './stores.js';

  export let location;
  onMount(() => currentPath.set('/privacy'));
</script>

<style>
  .page-header {
    padding-bottom: 24px;
    margin-bottom: 24px;
    border-bottom: 1px solid var(--border);
  }
</style>

<div class="page-content">
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
