<script>
  export let slug = '';
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

  const _forward = { x: 200, duration: 400 };
  const _forwardMobile = { x: 20, duration: 400 };
  const _back = { x: -200, duration: 400 };
  const _backMobile = { x: -20, duration: 400 };

  const transition = () => {
    const mobile = _screenWidth() <= 640;

    if ($currentStepIndex >= previousStepIndices.prev) {
      if (mobile) {
        return _forwardMobile;
      } else {
        return _forward;
      }
    } else {
      if (mobile) {
        return _backMobile;
      } else {
        return _back;
      }
    }
  }
</script>

<style>
  .step {
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: left;

    padding: var(--margin-left-default);
    max-width: 800px;
  }

  @media (max-width: 640px) {
    .step {
      padding: 15px;
      max-width: 96%;
    }
  }

  .instructions {
    padding: 30px 0;
    min-height: 200px;
  }
</style>

{#if slugOrder[$currentStepIndex] === slug}
  <div class="step" in:fly={transition()}>
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
