<script>
  export let slug = '';
  export let title = '';

  import { fly } from 'svelte/transition';
  import { navigate } from 'svelte-routing';

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
    min-height: 150px;
  }

  .prev-next {
    display: flex;
    flex-direction: row;
    justify-contents: space-between;
    min-width: 200px;
  }

  .prev-ctn {
    min-width: 100px;
  }

  .next-ctn {
    min-width: 100px;
    margin-left: 10px;
  }
</style>

{#if slugOrder[$currentStepIndex] === slug}
  <div class="step" in:fly={{ x: 400, duration: 500 }}>
    <h3>
      Step {$currentStepIndex + 1}: {title}
    </h3>

    <div class="instructions">
      <slot>
        (User instructions go here)
      </slot>
    </div>

    <div class="prev-next">
      <div class="prev-ctn">
        {#if $currentStepIndex > 0}
          <button class="secondary" on:click={() => {
            $currentStepIndex--;
            const newSlug = slugOrder[$currentStepIndex];
            navigate(`/step/${newSlug}`);
          }}>
            Previous
          </button>
        {/if}
      </div>

      <div class="next-ctn">
        {#if $currentStepIndex < slugOrder.length - 1}
          <button class="primary" on:click={() => {
            $currentStepIndex++;
            const newSlug = slugOrder[$currentStepIndex];
            navigate(`/step/${newSlug}`);
          }}>
            Next
          </button>
        {/if}
      </div>
    </div>
  </div>
{/if}
