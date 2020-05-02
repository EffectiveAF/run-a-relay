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
    margin-top: 104px;
  }
  h1, h2 {
    text-align: left;
  }
  h2 {
    margin: 24px 0px;
  }
  p {
    margin-bottom: 24px;
  }
</style>

<div class="container-fluid">

  <div class="home col-lg-offset-2">
    <div class="row col-lg-6"><h1>Thinking about running a Tor relay?</h1></div>
    
    <div class="row"><h2 class="col-lg-4">What is this?</h2></div>
    
    <div class="row"><p class="col-lg-4">It would make sense to introduce in this body of text what this is and how to get started. The rest of this text is lorem ipsum suspendisse ante urna, condimentum vel diam eget, condimentum convallis elit. Maecenas congue enim quis tincidunt fermentum. Sed sit amet neque sapien. Donec ac nunc tincidunt, sodales dolor sit amet, egestas diam.
    </p></div>
  </div>

  <Link to="/step/intro">
    <button class="primary hvr-ripple-out col-lg-offset-2">Get Started</button>
  </Link>
</div>