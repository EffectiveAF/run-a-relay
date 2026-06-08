<script>
  import { navigate } from 'svelte-routing';
  import { currentStepIndex, currentPath } from './stores.js';
  import slugOrder from './slugOrder.js';

  const stepTitles = [
    'Getting started',
    'Run or donate?',
    'Choosing a host',
    'Install Tor',
    "You're done!",
  ];

  const navTo = (idx) => {
    $currentStepIndex = idx;
    const slug = slugOrder[idx];
    currentPath.set(`/step/${slug}`);
    navigate(`/step/${slug}`);
  };
</script>

<style>
  .prev-next {
    display: flex;
    align-items: flex-start;
    justify-content: space-between;
    padding-top: 32px;
    margin-top: 40px;
    border-top: 1px solid var(--border);
    gap: 16px;
  }

  .nav-side { display: flex; flex-direction: column; gap: 4px; max-width: 48%; }
  .nav-side.right { align-items: flex-end; text-align: right; }

  .nav-label {
    font-size: 0.75rem;
    font-weight: 500;
    letter-spacing: 0.04em;
    color: var(--text-3);
  }

  .nav-btn {
    font-size: 0.875rem;
    font-weight: 500;
    color: var(--accent-light);
    background: none;
    border: none;
    padding: 0;
    box-shadow: none;
    height: auto;
    border-radius: 0;
    cursor: pointer;
    text-align: left;
    transition: color 0.1s;
    white-space: normal;
    line-height: 1.4;
  }

  .nav-side.right .nav-btn { text-align: right; }

  .nav-btn:hover { color: var(--text); background: none; }
  .placeholder { flex: 1; }
</style>

<svelte:window
  on:popstate={() => {
    const m = document.location.pathname.slice('/step/'.length).match(/[\w-]+/);
    if (m) {
      const idx = slugOrder.indexOf(m[0]);
      if (idx !== -1) $currentStepIndex = idx;
    }
  }}
  on:keydown={(e) => {
    if (e.target.tagName !== 'BODY' || e.altKey) return;
    if (e.key === 'ArrowLeft'  && $currentStepIndex > 0)                    navTo($currentStepIndex - 1);
    if (e.key === 'ArrowRight' && $currentStepIndex < slugOrder.length - 1) navTo($currentStepIndex + 1);
  }}
/>

<div class="prev-next">
  {#if $currentStepIndex > 0}
    <div class="nav-side left">
      <span class="nav-label">← Previous</span>
      <button class="nav-btn" on:click={() => navTo($currentStepIndex - 1)}>
        {stepTitles[$currentStepIndex - 1]}
      </button>
    </div>
  {:else}
    <div class="placeholder"></div>
  {/if}

  {#if $currentStepIndex < slugOrder.length - 1}
    <div class="nav-side right">
      <span class="nav-label">Next →</span>
      <button class="nav-btn" on:click={() => navTo($currentStepIndex + 1)}>
        {stepTitles[$currentStepIndex + 1]}
      </button>
    </div>
  {:else}
    <div class="placeholder"></div>
  {/if}
</div>
