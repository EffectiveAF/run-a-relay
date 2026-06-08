<script>
  export let slug  = '';
  export let title = '';

  import { fly } from 'svelte/transition';
  import { currentStepIndex, previousStepIndices } from './stores.js';
  import slugOrder from './slugOrder.js';

  function _screenWidth() {
    return Math.max(
      document.body.scrollWidth,
      document.documentElement.scrollWidth,
      document.body.offsetWidth,
      document.documentElement.offsetWidth,
      document.documentElement.clientWidth
    );
  }

  const transition = () =>
    $currentStepIndex >= previousStepIndices.prev
      ? { x: 32, duration: 260 }
      : { x: -32, duration: 260 };
</script>

<style>
  .step-header {
    margin-bottom: 28px;
  }

  .step-eyebrow {
    font-size: 0.75rem;
    font-weight: 500;
    letter-spacing: 0.04em;
    color: var(--text-3);
    margin-bottom: 10px;
    display: block;
  }

  h2 {
    font-size: 1.75rem;
    font-weight: 600;
    letter-spacing: -0.025em;
    line-height: 1.15;
    margin: 0;
    color: var(--text);
  }

  .step-body { }
</style>

{#if slugOrder[$currentStepIndex] === slug}
  <div in:fly={transition()}>
    <div class="step-header">
      <span class="step-eyebrow">
        Step {slugOrder.indexOf(slug) + 1} of {slugOrder.length}
      </span>
      <h2>{@html title}</h2>
    </div>

    <div class="step-body">
      <slot>(instructions go here)</slot>
    </div>
  </div>
{/if}
