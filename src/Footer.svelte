<script>
</script>

<style>
  footer {
    display: flex;
    align-items: center;
    border-top: 1px solid var(--border);
    height: 56px;
  }

  /* Left zone — fixed width matching sidebar */
  .footer-left {
    display: flex;
    align-items: center;
    width: var(--sidebar-w);
    flex-shrink: 0;
    padding: 0 16px 0 20px;
  }

  .footer-left span {
    font-size: 0.8125rem;
    color: var(--text-2);
    white-space: nowrap;
  }

  /* Vertical divider, same as nav */
  .v-divider {
    width: 1px;
    height: 20px;
    background: var(--border-strong);
    flex-shrink: 0;
  }

  /* Center — links */
  .footer-center {
    flex: 1;
    display: flex;
    align-items: center;
    padding: 0 20px;
    gap: 4px;
    min-width: 0;
  }

  .footer-link {
    display: inline-flex;
    align-items: center;
    height: 28px;
    padding: 0 8px;
    font-size: 0.8125rem;
    color: var(--text-3);
    text-decoration: none;
    border-radius: 5px;
    transition: color 0.1s, background 0.1s;
    white-space: nowrap;
  }

  .footer-link:hover {
    color: var(--text-2);
    background: var(--bg-hover);
  }

  /* Right — icons */
  .footer-right {
    display: flex;
    align-items: center;
    gap: 6px;
    padding: 0 16px;
    flex-shrink: 0;
  }

  .icon-link {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    width: 28px;
    height: 28px;
    border-radius: 5px;
    transition: background 0.1s;
  }

  .icon-link:hover { background: var(--bg-hover); }

  .icon-github {
    display: inline-block;
    width: 15px;
    height: 15px;
    background-color: var(--text-3);
    mask: url(/img/github_logo.svg) no-repeat center / contain;
    -webkit-mask: url(/img/github_logo.svg) no-repeat center / contain;
    transition: background-color 0.1s;
  }

  .icon-link:hover .icon-github { background-color: var(--text-2); }

  .icon-tor {
    display: inline-block;
    width: 15px;
    height: 15px;
    background-color: var(--text-3);
    mask: url(/img/tor.svg) no-repeat center / contain;
    -webkit-mask: url(/img/tor.svg) no-repeat center / contain;
    transition: background-color 0.1s;
  }

  .icon-link:hover .icon-tor { background-color: var(--accent-light); }

  @media (max-width: 900px) {
    .footer-left { display: none; }
    .v-divider   { display: none; }
    .footer-center { padding: 0 16px; }
  }

  @media (max-width: 640px) {
    footer { height: auto; padding: 12px 0; flex-wrap: wrap; }
    .footer-center { flex-wrap: wrap; }
  }
</style>

<footer>
  <!-- Left: copyright — aligns with sidebar -->
  <div class="footer-left">
    <span>© 2020 EffectiveAF</span>
  </div>

  <div class="v-divider"></div>

  <!-- Center: nav links -->
  <div class="footer-center">
    <a class="footer-link" href="/terms">Terms</a>
    <a class="footer-link" href="/privacy">Privacy</a>
    <a class="footer-link" href="/about">About</a>
  </div>

  <!-- Right: icons -->
  <div class="footer-right">
    <a
      class="icon-link"
      href="https://github.com/EffectiveAF/run-a-relay"
      target="_blank"
      rel="noopener noreferrer"
      aria-label="GitHub"
    >
      <div class="icon-github"></div>
    </a>
    <a
      class="icon-link"
      href="http://vxw5mjyap4bgb4w7h7qkeknxszmgp5efjtg77mnwjtbczynihsij3myd.onion"
      target="_blank"
      rel="noopener noreferrer"
      aria-label="Visit on Tor"
    >
      <div class="icon-tor"></div>
    </a>
  </div>
</footer>
