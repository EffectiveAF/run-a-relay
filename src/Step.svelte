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

  const _forward = { x: 40, duration: 300 };
  const _back    = { x: -40, duration: 300 };

  const transition = () =>
    $currentStepIndex >= previousStepIndices.prev ? _forward : _back;
</script>

<style>
  .step {
    padding: 48px 48px 80px;
    max-width: 720px;
  }

  .step-meta {
    display: flex;
    align-items: center;
    gap: 10px;
    margin-bottom: 24px;
  }

  .step-badge {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    width: 22px;
    height: 22px;
    border-radius: 50%;
    background: var(--accent);
    color: #fff;
    font-size: 0.7rem;
    font-weight: 700;
    flex-shrink: 0;
  }

  .step-label {
    font-size: 0.75rem;
    font-weight: 600;
    letter-spacing: 0.06em;
    text-transform: uppercase;
    color: var(--text-3);
  }

  h2 {
    font-size: 1.625rem;
    font-weight: 700;
    letter-spacing: -0.03em;
    margin: 0 0 28px;
  }

  .step-body {
    min-height: 160px;
  }

  @media (max-width: 900px) {
    .step { padding: 32px 24px 64px; }
  }

  @media (max-width: 640px) {
    .step { padding: 24px 16px 48px; }
    h2 { font-size: 1.375rem; }
  }
</style>

{#if slugOrder[$currentStepIndex] === slug}
  <div class="step" in:fly={transition()}>
    <div class="step-meta">
      <span class="step-badge">{slugOrder.indexOf(slug) + 1}</span>
      <span class="step-label">Step {slugOrder.indexOf(slug) + 1} of {slugOrder.length}</span>
    </div>

    <h2>{title}</h2>

    <div class="step-body">
      <slot>(instructions go here)</slot>
    </div>
  </div>
{/if}
