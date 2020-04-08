<script>
  export let href = '#';
  export let primary = false;

  import ExternalLink from './ExternalLink.svelte';

  const primaryClass = "primary hvr-ripple-out";
  const secondaryClass = "secondary hvr-ripple-out-sec";

  const btnClass = primary ? primaryClass : secondaryClass;
</script>

<style>
  button {
    display: flex;
    flex-direction: row;
  }

  .external-link-ctn {
    width: 20px;
    height: 20px;
    margin-left: 8px;
    mask: url(/img/open_external.svg) no-repeat center / contain;
    -webkit-mask: url(/img/open_external.svg) no-repeat center / contain;
  }

  .primary > .external-link-ctn {
    background-color: var(--content-inverted);
  }

  .secondary > .external-link-ctn {
    background-color: var(--content-default);
  }
</style>

<ExternalLink href={href}>
  <button class={btnClass}>
    <slot />
    <div class="external-link-ctn"></div>
  </button>
</ExternalLink>
