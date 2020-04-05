<script>
  import { Link } from 'svelte-routing';

  import ExternalButton from './ExternalButton.svelte';
</script>

<style>
  nav {
    display: flex;
    flex-direction: row;
    justify-content: left;
    align-items: center;
    border-bottom: 1px solid var(--neutral-100);
    padding: 0 16px;
    height: 64px;
  }

  nav a:hover {
    text-decoration: none;
  }

  .nav-right {
    display: flex;
    margin-left: auto;
  }

  nav img {
    display: block;
    max-width: 32px;
    max-height: 32px;
    margin-right: 8px;
  }
</style>

<nav>
  <Link to="/"><img src="/img/tor.svg" alt="Tor-esque logo (onion)" /></Link>

  <h4>Run A Relay</h4>

  <div class="nav-right">
    <Link to="/about">
      <button class="secondary hvr-ripple-out-sec">About</button>
    </Link>

    <ExternalButton href="https://donate.torproject.org/">
      Donate to Tor
    </ExternalButton>
  </div>
</nav>
