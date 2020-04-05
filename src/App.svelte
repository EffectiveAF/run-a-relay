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
    border-bottom: 1px solid var(--neutral-100);
    padding: 0px 16px;
    height: 64px;
  }

  nav a:hover {
    text-decoration: none;
  }
  
  .nav-right {
    display: flex;
    margin-left: auto;
  }

  .nav-right button {
    margin-left: 16px;
  }

  img {
    max-width: 32px;
    max-height: 32px;
    margin-right: 8px;
  }
</style>

<Router>
  <nav>

    <Link to="/"><img src="/img/tor.svg" alt="Tor-esque logo (onion)" /></Link>

    <h4>Run A Relay</h4>
    <div class="nav-right">
      <a href="https://github.com/EffectiveAF/run-a-relay#faq" target="_blank" rel="nofollow noreferrer noopener">
        <button class="secondary hvr-ripple-out-sec">RaR FAQ</button>
      </a>
      <a href="https://donate.torproject.org/" target="_blank" rel="nofollow noreferrer noopener">
        <button class="primary hvr-ripple-out">Donate to Tor</button>
      </a>
    </div>
  </nav>
  <div>
    <Route path="/"><Home /></Route>
  </div>
</Router>
