<script>
  import { onMount } from 'svelte';

  onMount(async () => {
    // Add remote state from metrics.torproject.org ?
  })
</script>

<style>
  .home {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }
</style>

<div class="home">
  <h3>Welcome!</h3>
  <div>

      <h1>This is an h1</h1>
      <h2>This is an h2</h2>
      <p>This is body Copy</p>


      <button class="primary hvr-ripple-out">Donate to Tor</button>
      <br>
      <br>
      <button class="secondary hvr-ripple-out-sec">Label</button>
    
  </div>
</div>
