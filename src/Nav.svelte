<script>
  import { Link } from 'svelte-routing';
  import { theme, currentPath, currentStepIndex } from './stores.js';

  const toggleTheme = () => theme.update(t => t === 'dark' ? 'light' : 'dark');

  const stepTitles = {
    'intro':             'Introduction',
    'relay-or-donate':   'Run or donate?',
    'choosing-a-host':   'Choosing a host',
    'install-tor-daemon':'Install Tor',
    'thank-you':         "You're done!",
  };

  // Build breadcrumb from current path
  $: breadcrumb = (() => {
    const p = $currentPath;
    if (p === '/' || p === '')        return [{ label: 'Home' }];
    if (p.startsWith('/step/'))       return [{ label: 'Guide' }, { label: stepTitles[p.replace('/step/', '')] || 'Step' }];
    if (p === '/about')               return [{ label: 'About' }];
    if (p === '/terms')               return [{ label: 'Terms of Service' }];
    if (p === '/privacy')             return [{ label: 'Privacy Policy' }];
    return [{ label: 'Run A Relay' }];
  })();
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
    background: var(--bg);
    border-bottom: 1px solid var(--border);
  }

  /* ── Left zone — fixed width matching sidebar ── */
  .nav-left {
    display: flex;
    align-items: center;
    width: var(--sidebar-w);
    flex-shrink: 0;
    padding: 0 16px 0 20px;
    gap: 0;
  }

  .logo-link {
    display: flex;
    align-items: center;
    gap: 9px;
    text-decoration: none;
    flex-shrink: 0;
  }

  .logo-link img {
    width: 20px;
    height: 20px;
    opacity: 0.9;
  }

  .brand-name {
    font-size: 0.9375rem;
    font-weight: 510;
    color: var(--text);
    letter-spacing: -0.01em;
    white-space: nowrap;
  }

  /* ── Vertical divider between left + center ── */
  .v-divider {
    width: 1px;
    height: 20px;
    background: var(--border-strong);
    flex-shrink: 0;
  }

  /* ── Center zone — breadcrumb ── */
  .nav-center {
    flex: 1;
    display: flex;
    align-items: center;
    padding: 0 20px;
    min-width: 0;
  }

  .breadcrumb {
    display: flex;
    align-items: center;
    gap: 6px;
    font-size: 0.8125rem;
    color: var(--text-2);
    white-space: nowrap;
    overflow: hidden;
  }

  .breadcrumb-sep {
    color: var(--text-3);
    font-size: 0.75rem;
  }

  .breadcrumb span:last-child {
    color: var(--text);
    font-weight: 500;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  /* ── Right zone — actions ── */
  .nav-right {
    display: flex;
    align-items: center;
    gap: 4px;
    padding: 0 16px;
    flex-shrink: 0;
  }

  .nav-link {
    display: inline-flex;
    align-items: center;
    height: 30px;
    padding: 0 10px;
    font-size: 0.8125rem;
    font-weight: 400;
    color: var(--text-2);
    text-decoration: none;
    border-radius: 6px;
    transition: color 0.1s, background 0.1s;
  }

  .nav-link:hover {
    color: var(--text);
    background: var(--bg-hover);
  }

  /* Donate = filled button, like Linear's "Sign up" */
  .btn-cta {
    display: inline-flex;
    align-items: center;
    height: 30px;
    padding: 0 12px;
    background: var(--text);
    color: var(--bg);
    font-size: 0.8125rem;
    font-weight: 510;
    border-radius: 6px;
    text-decoration: none;
    transition: opacity 0.1s;
    white-space: nowrap;
    margin-left: 4px;
  }

  .btn-cta:hover { color: var(--bg); opacity: 0.88; }

  /* Theme toggle — icon-only, like Linear's sun/moon */
  .theme-btn {
    width: 30px;
    height: 30px;
    padding: 0;
    border-radius: 6px;
    background: transparent;
    border: none;
    color: var(--text-2);
    font-size: 0.9rem;
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
    box-shadow: none;
    transition: background 0.1s, color 0.1s;
  }

  .theme-btn:hover { background: var(--bg-hover); color: var(--text); }

  @media (max-width: 900px) {
    .nav-left { width: auto; }
    .breadcrumb { display: none; }
  }

  @media (max-width: 640px) {
    .hide-mobile { display: none; }
  }
</style>

<nav>
  <!-- Left: logo + brand — aligns with sidebar -->
  <div class="nav-left">
    <a class="logo-link" href="/">
      <img src="/img/tor.svg" alt="Tor onion logo" />
      <span class="brand-name">Run A Relay</span>
    </a>
  </div>

  <div class="v-divider"></div>

  <!-- Center: breadcrumb showing current location -->
  <div class="nav-center">
    <div class="breadcrumb">
      {#each breadcrumb as crumb, i}
        {#if i > 0}<span class="breadcrumb-sep">/</span>{/if}
        <span>{crumb.label}</span>
      {/each}
    </div>
  </div>

  <!-- Right: actions -->
  <div class="nav-right">
    <a class="nav-link hide-mobile" href="/about">About</a>

    <a
      class="nav-link hide-mobile"
      href="https://github.com/EffectiveAF/run-a-relay"
      target="_blank"
      rel="noopener noreferrer"
    >GitHub</a>

    <a
      class="btn-cta hide-mobile"
      href="https://donate.torproject.org/"
      target="_blank"
      rel="noopener noreferrer"
    >Donate</a>

    <button class="theme-btn" on:click={toggleTheme} title="Toggle theme">
      {#if $theme === 'dark'}☀{:else}☾{/if}
    </button>
  </div>
</nav>
