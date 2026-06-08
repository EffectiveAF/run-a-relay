<script>
  export let href = '#';
  export let primary = false;
  import ExternalLink from './ExternalLink.svelte';
</script>

<style>
  .ext-btn {
    display: inline-flex;
    align-items: center;
    gap: 6px;
    height: 36px;
    padding: 0 16px;
    border-radius: 8px;
    font-size: 0.8125rem;
    font-weight: 500;
    text-decoration: none;
    transition: filter 0.12s, box-shadow 0.12s, background 0.12s;
    white-space: nowrap;
    cursor: pointer;
  }

  .ext-btn.primary {
    background: var(--accent);
    color: #fff;
    box-shadow: 0 2px 10px var(--accent-glow);
  }

  .ext-btn.primary:hover {
    color: #fff;
    filter: brightness(1.1);
    box-shadow: 0 4px 16px var(--accent-glow);
  }

  .ext-btn.secondary {
    background: var(--bg-raised);
    color: var(--text);
    border: 1px solid var(--border-strong);
  }

  .ext-btn.secondary:hover {
    color: var(--text);
    background: var(--bg-subtle);
  }

  .arrow {
    font-size: 0.7rem;
    opacity: 0.7;
  }
</style>

<ExternalLink href={href}>
  <span class="ext-btn" class:primary class:secondary={!primary}>
    <slot />
    <span class="arrow">↗</span>
  </span>
</ExternalLink>
