<script>
  import { onMount } from 'svelte';
  import { currentPath } from './stores.js';
  import ExternalLink from './ExternalLink.svelte';

  export let location;

  onMount(() => currentPath.set('/terms'));
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

  h2 { margin: 28px 0 10px; }

  @media (max-width: 900px) { .page { padding: 32px 24px 64px; } }
  @media (max-width: 640px) { .page { padding: 24px 16px 48px; } }
</style>

<div class="page">
  <div class="page-header">
    <h1>Terms of Service</h1>
  </div>

  <p>
    EffectiveAF is a proud supporter of the
    <ExternalLink href="https://ethicalsource.dev/">Ethical Source</ExternalLink>
    movement. Run A Relay is released under an Ethical Source license,
    namely The Hippocratic License 2.1.
  </p>

  <p>
    You may do as you wish with this website and
    <ExternalLink href="https://github.com/EffectiveAF/run-a-relay">its software</ExternalLink>
    provided that you do not use it to violate human rights as defined by the
    UN Universal Declaration of Human Rights. See
    <ExternalLink href="https://github.com/EffectiveAF/run-a-relay/blob/master/LICENSE.md">LICENSE.md</ExternalLink>
    or <ExternalLink href="https://firstdonoharm.dev/">The Hippocratic License 2.1</ExternalLink>
    for details.
  </p>
</div>
