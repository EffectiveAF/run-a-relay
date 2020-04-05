<script>
  import { Router, Link, Route } from 'svelte-routing';

  import Home from './Home.svelte';
</script>

<style>
  nav {
    display: flex;
    flex-direction: row;
    justify-content: left;
    align-items: center;
  }
  .nav-right {
    margin-left: auto;
  }
  .nav-right a {
    background-color: purple;
  }
</style>

<Router>
  <nav>
    <h4>Run A Relay</h4>
    <div class="nav-right">
      <a href="https://github.com/EffectiveAF/run-a-relay#faq" target="_blank" rel="nofollow noreferrer noopener">
        <button>RaR FAQ</button>
      </a>
      <a href="https://donate.torproject.org/" target="_blank" rel="nofollow noreferrer noopener">
        <button>Donate to Tor</button>
      </a>
    </div>
  </nav>
  <div>
    <Route path="/"><Home /></Route>
  </div>
</Router>
