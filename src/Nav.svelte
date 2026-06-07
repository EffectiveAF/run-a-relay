<script>
  import { onMount } from 'svelte';
  import { Link } from 'svelte-routing';

  import ExternalButton from './ExternalButton.svelte';
  import { theme } from './stores.js';

  // Derived from https://stackoverflow.com/a/28344281/197160
  function _hasClass(el, cls) {
    return !!el.className.match(new RegExp('(\\s|^)'+cls+'(\\s|$)'));
  }
  function _addClass(el, cls) {
    if (!_hasClass(el, cls)) {
      el.className += ' ' + cls;
    }
  }
  function _removeClass(el, cls) {
    if (_hasClass(el,cls)) {
      var reg = new RegExp('(\\s|^)'+cls+'(\\s|$)');
      el.className = el.className.replace(reg, '');
    }
  }

  // Styling in public/relay.css
  const _toggleClass = 'show-right-nav-mobile';

  // Used in helpers below
  let _navRight;
  onMount(async () => {
    _navRight = document.getElementsByClassName('nav-right')[0];
  })

  const collapseNav = () => {
    _navRight.style.visibility = 'hidden';
    _removeClass(_navRight, _toggleClass);
  }
  const showNav = () => {
    _navRight.style.visibility = 'visible';
    _addClass(_navRight, _toggleClass);
  }

  const toggleNavRight = (e) => {
    if (_hasClass(_navRight, _toggleClass)) {
      collapseNav();
    } else {
      showNav();
    }
  }

  const toggleTheme = () => {
    theme.update(t => t === 'dark' ? 'light' : 'dark');
  }
</script>

<style>
  nav {
    display: flex;
    flex-direction: row;
    justify-content: left;
    align-items: center;
    background-color: var(--bg-nav);
    border-bottom: 1px solid var(--border-color);
    padding: 0 20px;
    height: 64px;
    position: sticky;
    top: 0;
    z-index: 10;
    backdrop-filter: blur(8px);
  }

  nav a:hover {
    text-decoration: none;
  }

  .nav-brand {
    display: flex;
    align-items: center;
    gap: 10px;
  }

  .nav-title {
    font-size: 16px;
    font-weight: 700;
    letter-spacing: -0.01em;
    color: var(--text-primary);
  }

  .toggle-nav-right, .nav-right {
    display: flex;
    margin-left: auto;
    align-items: center;
    gap: 8px;
  }

  .toggle-nav-right {
    display: none;
  }

  nav img.logo {
    display: block;
    width: 28px;
    height: 28px;
  }

  .theme-toggle {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 38px;
    height: 38px;
    padding: 0;
    font-size: 16px;
    background-color: transparent;
    border: 1px solid var(--border-color);
    border-radius: 6px;
    color: var(--text-primary);
    cursor: pointer;
    box-shadow: none;
    transition: border-color 0.15s ease, background-color 0.15s ease;
    margin-left: 8px;
  }

  .theme-toggle:hover {
    border-color: var(--accent-primary);
    background-color: var(--button-secondary-bg);
  }

  @media (max-width: 640px) {
    nav {
      padding: 0 12px;
    }

    .toggle-nav-right {
      display: flex;
      height: 36px;
      width: 36px;
      padding: 0;
      border: none;
      box-shadow: none;
      background: transparent;
      cursor: pointer;
    }

    .nav-right {
      visibility: hidden;
      flex-direction: column;
      justify-content: center;
      align-items: center;
      z-index: 2;
    }

    .theme-toggle {
      margin-left: 0;
    }
  }
</style>

<nav>
  <Link to="/" class="nav-brand">
    <img class="logo" src="/img/tor.svg" alt="Tor-esque logo (onion)" />
    <span class="nav-title">Run A Relay</span>
  </Link>

  <div class="nav-right">
    <Link to="/about">
      <button class="secondary hvr-ripple-out-sec">About</button>
    </Link>

    <ExternalButton href="https://github.com/EffectiveAF/run-a-relay">
      GitHub
    </ExternalButton>

    <ExternalButton primary href="https://donate.torproject.org/">
      Donate to Tor
    </ExternalButton>

    <button class="theme-toggle" on:click={toggleTheme} title="Toggle theme">
      {#if $theme === 'dark'}☀{:else}☾{/if}
    </button>
  </div>

  <img
    class="toggle-nav-right"
    src="/img/hamburger_menu.svg"
    alt="Hamburger Menu"
    tabindex="1"
    on:click={toggleNavRight}
    on:blur={() => {
      // TODO: Remove this race condition hack
      setTimeout(collapseNav, 250);
    }}
  />
</nav>
