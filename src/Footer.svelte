<script>
  import { Link } from 'svelte-routing';
  import ExternalLink from './ExternalLink.svelte';
</script>

<style>
  footer {
    margin-left: var(--sidebar-w);
    border-top: 1px solid var(--border);
    padding: 16px 40px;
    display: flex;
    align-items: center;
    justify-content: space-between;
    flex-wrap: wrap;
    gap: 12px;
  }

  .left {
    display: flex;
    align-items: center;
    gap: 16px;
    flex-wrap: wrap;
  }

  footer a, footer span {
    font-size: 0.78rem;
    color: var(--text-3);
    text-decoration: none;
    transition: color 0.1s;
  }

  footer a:hover { color: var(--text-2); }

  .icons {
    display: flex;
    align-items: center;
    gap: 10px;
  }

  .icon-github {
    display: inline-block;
    width: 16px;
    height: 16px;
    background-color: var(--text-3);
    mask: url(/img/github_logo.svg) no-repeat center / contain;
    -webkit-mask: url(/img/github_logo.svg) no-repeat center / contain;
    transition: background-color 0.1s;
  }

  .icon-github:hover { background-color: var(--text-2); }

  .icon-tor {
    display: inline-block;
    width: 16px;
    height: 16px;
    background-color: var(--text-3);
    mask: url(/img/tor.svg) no-repeat center / contain;
    -webkit-mask: url(/img/tor.svg) no-repeat center / contain;
    transition: background-color 0.1s;
  }

  .icon-tor:hover { background-color: var(--accent-light); }

  @media (max-width: 900px) {
    footer { margin-left: 0; padding: 14px 24px; }
  }

  @media (max-width: 640px) {
    footer { padding: 14px 16px; flex-direction: column; align-items: flex-start; }
  }
</style>

<footer>
  <div class="left">
    <span>© 2020 <ExternalLink href="https://effective.af/">EffectiveAF</ExternalLink></span>
    <Link to="/terms">Terms</Link>
    <Link to="/privacy">Privacy</Link>
    <Link to="/about">About</Link>
  </div>

  <div class="icons">
    <ExternalLink href="https://github.com/EffectiveAF/run-a-relay">
      <div class="icon-github" role="img" aria-label="GitHub"></div>
    </ExternalLink>
    <ExternalLink href="http://vxw5mjyap4bgb4w7h7qkeknxszmgp5efjtg77mnwjtbczynihsij3myd.onion">
      <div class="icon-tor" role="img" aria-label="Visit on Tor"></div>
    </ExternalLink>
  </div>
</footer>
