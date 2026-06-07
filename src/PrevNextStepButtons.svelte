<script>
  import { navigate } from 'svelte-routing';
  import { currentStepIndex, currentPath } from './stores.js';
  import slugOrder from './slugOrder.js';

  const navPrevious = () => {
    $currentStepIndex--;
    const slug = slugOrder[$currentStepIndex];
    currentPath.set(`/step/${slug}`);
    navigate(`/step/${slug}`);
  };

  const navNext = () => {
    $currentStepIndex++;
    const slug = slugOrder[$currentStepIndex];
    currentPath.set(`/step/${slug}`);
    navigate(`/step/${slug}`);
  };
</script>

<style>
  .prev-next {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 48px 64px;
    max-width: 720px;
    border-top: 1px solid var(--border);
    margin-top: 16px;
    padding-top: 24px;
  }

  .nav-btn {
    display: flex;
    flex-direction: column;
    gap: 2px;
  }

  .nav-btn.next { align-items: flex-end; }

  .nav-label {
    font-size: 0.7rem;
    font-weight: 600;
    letter-spacing: 0.06em;
    text-transform: uppercase;
    color: var(--text-3);
  }

  .nav-title {
    font-size: 0.875rem;
    font-weight: 500;
    color: var(--accent-light);
    cursor: pointer;
    background: none;
    border: none;
    padding: 0;
    box-shadow: none;
    height: auto;
    border-radius: 0;
    transition: color 0.12s;
  }

  .nav-title:hover {
    color: var(--text);
    background: none;
  }

  .placeholder { width: 1px; }

  @media (max-width: 900px) {
    .prev-next { padding-left: 24px; padding-right: 24px; }
  }

  @media (max-width: 640px) {
    .prev-next { padding-left: 16px; padding-right: 16px; }
  }
</style>

<svelte:window
  on:popstate={() => {
    const maybeStep = document.location.pathname
      .slice('/step/'.length)
      .match(/[\w\-]+/);
    if (!maybeStep || maybeStep.length !== 1) return;
    const idx = slugOrder.indexOf(maybeStep[0]);
    if (idx !== -1) $currentStepIndex = idx;
  }}
  on:keydown={(e) => {
    if (e.target.tagName !== 'BODY' || e.altKey) return;
    if ($currentStepIndex > 0 && e.key === 'ArrowLeft') { navPrevious(); return; }
    if ($currentStepIndex < slugOrder.length - 1 && e.key === 'ArrowRight') { navNext(); return; }
  }}
/>

<div class="prev-next">
  {#if $currentStepIndex > 0}
    <div class="nav-btn prev">
      <span class="nav-label">← Previous</span>
      <button class="nav-title" on:click={navPrevious}>
        {slugOrder[$currentStepIndex - 1].replace(/-/g, ' ').replace(/\b\w/g, c => c.toUpperCase())}
      </button>
    </div>
  {:else}
    <div class="placeholder"></div>
  {/if}

  {#if $currentStepIndex < slugOrder.length - 1}
    <div class="nav-btn next">
      <span class="nav-label">Next →</span>
      <button class="nav-title" on:click={navNext}>
        {slugOrder[$currentStepIndex + 1].replace(/-/g, ' ').replace(/\b\w/g, c => c.toUpperCase())}
      </button>
    </div>
  {:else}
    <div class="placeholder"></div>
  {/if}
</div>
