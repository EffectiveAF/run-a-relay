<script>
  import { navigate } from 'svelte-routing';

  import { currentStepIndex } from './stores.js';
  import slugOrder from './slugOrder.js';

  const navPrevious = () => {
    $currentStepIndex--;
    const newSlug = slugOrder[$currentStepIndex];
    navigate(`/step/${newSlug}`);
  }

  const navNext = () => {
    $currentStepIndex++;
    const newSlug = slugOrder[$currentStepIndex];
    navigate(`/step/${newSlug}`);
  }
</script>

<style>
  .prev-next {
    display: flex;
    flex-direction: row;
    justify-contents: space-between;
    min-width: 240px;
    padding: var(--margin-left-default);
  }

  .prev-ctn {
    min-width: 130px;
  }

  .next-ctn {
    min-width: 110px;
    margin-left: 10px;
  }

  @media (max-width: 640px) {
    .prev-next {
      padding: var(--margin-left-default-mobile);
      max-width: 96%;
    }
    .next-ctn {
      margin-left: 0;
    }
  }
</style>

<svelte:window
  on:popstate={(e) => {
    // Pulling the stepName from the router isn't good enough because
    // it's not reactive...
    const maybeStepName = document.location.pathname
          .slice('/step/'.length)
          .match(/[\w\-]+/);
    if (!maybeStepName || maybeStepName.length !== 1) {
      return;
    }
    const stepName = maybeStepName[0];

    // Re-sync URL displayed and currentStepIndex
    const stepIndex = slugOrder.indexOf(stepName);
    if (stepIndex !== -1) {
      $currentStepIndex = stepIndex;
    }
  }}

  on:keydown={(e) => {
    if (e.target.tagName !== 'BODY') {
      // Don't trigger keyboard shortcuts in code <input>s
      return;
    }

    if (e.altKey) {
      // Don't trigger keyboard shortcuts on left/right if user is
      // typing alt+left or alt+right
      return;
    }

    if ($currentStepIndex > 0) {
      if (['ArrowLeft'].indexOf(e.key) !== -1) {
        navPrevious();
        return;
      }
    }
    if ($currentStepIndex < slugOrder.length - 1) {
      if (['ArrowRight'].indexOf(e.key) !== -1) {
        navNext();
        return;
      }
    }
  }}
/>

<div class="prev-next">
  <div class="prev-ctn">
    {#if $currentStepIndex > 0}
      <button class="secondary" on:click={navPrevious}>
        Previous
      </button>
    {/if}
  </div>

  <div class="next-ctn">
    {#if $currentStepIndex < slugOrder.length - 1}
      <button class="primary" on:click={navNext}>
        Next
      </button>
    {/if}
  </div>
</div>
