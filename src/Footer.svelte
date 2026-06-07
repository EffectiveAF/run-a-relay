<script>
  import { Link } from 'svelte-routing';
  import ExternalLink from './ExternalLink.svelte';
</script>

<style>
  footer {
    margin-left: var(--sidebar-w);
    border-top: 1px solid var(--border);
    padding: 20px 48px;
    display: flex;
    align-items: center;
    justify-content: space-between;
    flex-wrap: wrap;
    gap: 12px;
  }

  .left, .right {
    display: flex;
    align-items: center;
    gap: 20px;
    flex-wrap: wrap;
  }

  .copy {
    font-size: 0.78rem;
    color: var(--text-3);
  }

  a {
    font-size: 0.78rem;
    color: var(--text-3);
    text-decoration: none;
    transition: color 0.12s;
  }

  a:hover { color: var(--text-2); }

  .icon-row {
    display: flex;
    align-items: center;
    gap: 12px;
  }

  .icon-github {
    display: inline-block;
    width: 18px;
    height: 18px;
    background-color: var(--text-3);
    mask: url(/img/github_logo.svg) no-repeat center / contain;
    -webkit-mask: url(/img/github_logo.svg) no-repeat center / contain;
    transition: background-color 0.12s;
  }

  .icon-github:hover { background-color: var(--text-2); }

  .icon-tor {
    display: inline-block;
    width: 18px;
    height: 18px;
    background-color: var(--text-3);
    mask: url(/img/tor.svg) no-repeat center / contain;
    -webkit-mask: url(/img/tor.svg) no-repeat center / contain;
    transition: background-color 0.12s;
  }

  .icon-tor:hover { background-color: var(--accent-light); }

  @media (max-width: 900px) {
    footer { margin-left: 0; padding: 16px 24px; }
  }

  @media (max-width: 640px) {
    footer { padding: 16px; flex-direction: column; align-items: flex-start; }
  }
</style>

<footer>
  <div class="left">
    <span class="copy">© 2020 <ExternalLink href="https://effective.af/">EffectiveAF</ExternalLink></span>
    <Link to="/terms">Terms</Link>
    <Link to="/privacy">Privacy</Link>
    <ExternalLink href="https://github.com/elimisteve">Steve Phillips</ExternalLink>
    &amp;
    <ExternalLink href="https://github.com/timsully">Tim Sullivan</ExternalLink>
  </div>

  <div class="right">
    <div class="icon-row">
      <ExternalLink href="https://github.com/EffectiveAF/run-a-relay">
        <div class="icon-github" role="img" aria-label="GitHub"></div>
      </ExternalLink>
      <ExternalLink href="http://vxw5mjyap4bgb4w7h7qkeknxszmgp5efjtg77mnwjtbczynihsij3myd.onion">
        <div class="icon-tor" role="img" aria-label="Visit on Tor"></div>
      </ExternalLink>
    </div>
  </div>
</footer>
