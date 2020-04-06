<script>
  import { navigate } from 'svelte-routing';

  import { currentStepIndex } from './stores.js';
  import slugOrder from './slugOrder.js';
</script>

<style>
  .prev-next {
    display: flex;
    flex-direction: row;
    justify-contents: space-between;
    min-width: 240px;
    margin-left: 30px;
  }

  .prev-ctn {
    min-width: 130px;
  }

  .next-ctn {
    min-width: 110px;
    margin-left: 10px;
  }
</style>

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
