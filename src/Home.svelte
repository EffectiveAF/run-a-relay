<script>
  import { onMount } from 'svelte';
  import { Link } from 'svelte-routing';

  import { currentStepIndex } from './stores.js';

  onMount(async () => {
    $currentStepIndex = 0;

    // Add remote state from metrics.torproject.org ?
  })
</script>

<style>
  .home {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    margin-top: 30px;
  }
  h1, h2 {
    text-align: center;
  }
  h2 {
    margin-top: 40px;
    margin-bottom: 40px;
  }
</style>

<div class="home">
  <h1>Thinking about running a Tor relay?</h1>
  <h2>You're in the right place!</h2>

  <Link to="/step/intro">
    <button class="primary hvr-ripple-out">Get Started</button>
  </Link>
</div>
