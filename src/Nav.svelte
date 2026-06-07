<script>
  import { Link } from 'svelte-routing';
  import { theme } from './stores.js';

  const toggleTheme = () => theme.update(t => t === 'dark' ? 'light' : 'dark');

  let mobileOpen = false;
  const toggleMobile = () => mobileOpen = !mobileOpen;
</script>

<style>
  nav {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    height: var(--nav-h);
    z-index: 100;
    display: flex;
    align-items: center;
    padding: 0 20px;
    background: var(--bg);
    border-bottom: 1px solid var(--border);
  }

  .brand {
    display: flex;
    align-items: center;
    gap: 9px;
    text-decoration: none;
    flex-shrink: 0;
  }

  .brand img {
    width: 24px;
    height: 24px;
    opacity: 0.85;
  }

  .brand-name {
    font-size: 0.875rem;
    font-weight: 600;
    color: var(--text);
    letter-spacing: -0.01em;
  }

  .divider {
    width: 1px;
    height: 20px;
    background: var(--border-strong);
    margin: 0 16px;
    flex-shrink: 0;
  }

  .badge {
    font-size: 0.7rem;
    font-weight: 600;
    color: var(--text-2);
    letter-spacing: 0.04em;
    text-transform: uppercase;
  }

  .spacer { flex: 1; }

  .nav-actions {
    display: flex;
    align-items: center;
    gap: 6px;
  }

  .nav-link {
    font-size: 0.8125rem;
    font-weight: 400;
    color: var(--text-2);
    text-decoration: none;
    padding: 5px 10px;
    border-radius: 6px;
    transition: color 0.12s, background 0.12s;
  }

  .nav-link:hover {
    color: var(--text);
    background: var(--bg-hover);
  }

  .btn-donate {
    display: inline-flex;
    align-items: center;
    gap: 5px;
    height: 30px;
    padding: 0 12px;
    background: var(--accent);
    color: #fff;
    font-size: 0.78rem;
    font-weight: 500;
    border-radius: 6px;
    text-decoration: none;
    transition: filter 0.12s, box-shadow 0.12s;
    box-shadow: 0 2px 8px var(--accent-glow);
    white-space: nowrap;
  }

  .btn-donate:hover {
    color: #fff;
    filter: brightness(1.1);
    box-shadow: 0 4px 14px var(--accent-glow);
  }

  .theme-btn {
    width: 32px;
    height: 32px;
    padding: 0;
    border-radius: 8px;
    background: transparent;
    border: 1px solid var(--border-strong);
    color: var(--text-2);
    font-size: 0.85rem;
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
    box-shadow: none;
    transition: background 0.12s, border-color 0.12s;
  }

  .theme-btn:hover {
    background: var(--bg-hover);
    color: var(--text);
  }

  @media (max-width: 640px) {
    .hide-mobile { display: none; }
  }
</style>

<nav>
  <Link to="/" class="brand">
    <img src="/img/tor.svg" alt="Tor onion logo" />
    <span class="brand-name">Run A Relay</span>
  </Link>

  <div class="divider"></div>
  <span class="badge">Guide</span>

  <div class="spacer"></div>

  <div class="nav-actions">
    <Link to="/about" class="nav-link hide-mobile">About</Link>

    <a
      class="nav-link hide-mobile"
      href="https://github.com/EffectiveAF/run-a-relay"
      target="_blank"
      rel="noopener noreferrer"
    >GitHub</a>

    <a
      class="btn-donate hide-mobile"
      href="https://donate.torproject.org/"
      target="_blank"
      rel="noopener noreferrer"
    >Donate ↗</a>

    <button class="theme-btn" on:click={toggleTheme} title="Toggle theme">
      {#if $theme === 'dark'}☀{:else}☾{/if}
    </button>
  </div>
</nav>
