<script>
  export let slug = '';
  export let title = '';

  import { fly } from 'svelte/transition';

  import { currentStepIndex } from './stores.js';
  import slugOrder from './slugOrder.js';
</script>

<style>
  .step {
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: left;

    padding: 40px;
    max-width: 80%;
  }

  @media (max-width: 640px) {
    .step {
      padding: 10px;
      max-width: 96%;
    }
  }

  .instructions {
    padding: 30px 0;
    min-height: 200px;
  }
</style>

{#if slugOrder[$currentStepIndex] === slug}
  <div class="step" in:fly={{ x: 300, duration: 500 }}>
    <h3>
      Step {$currentStepIndex + 1}: {title}
    </h3>

    <div class="instructions">
      <slot>
        (User instructions go here)
      </slot>
    </div>
  </div>
{/if}
