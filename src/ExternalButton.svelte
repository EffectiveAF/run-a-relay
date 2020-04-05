<script>
  export let href = '#';

  import ExternalLink from './ExternalLink.svelte';
</script>

<ExternalLink href={href}>
  <button class="primary hvr-ripple-out">
    <slot />
  </button>
</ExternalLink>
