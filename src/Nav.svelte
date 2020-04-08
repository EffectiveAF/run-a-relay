<script>
  import { onMount } from 'svelte';
  import { Link } from 'svelte-routing';

  import ExternalButton from './ExternalButton.svelte';

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

  const togglerLabelOpen = 'V';
  const togglerLabelClose = 'X';

  // Styling in public/relay.css
  const _toggleClass = 'show-right-nav-mobile';

  // Used in helpers below
  let _navRight, toggler;
  onMount(async () => {
    _navRight = document.getElementsByClassName('nav-right')[0];
    toggler = document.getElementsByClassName('toggle-nav-right')[0];
  })

  const collapseNav = () => {
    _navRight.style.visibility = 'hidden';
    _removeClass(_navRight, _toggleClass);
    toggler.textContent = togglerLabelOpen;
  }
  const showNav = () => {
    _navRight.style.visibility = 'visible';
    _addClass(_navRight, _toggleClass);
    toggler.textContent = togglerLabelClose;
  }

  const toggleNavRight = (e) => {
    if (_hasClass(_navRight, _toggleClass)) {
      collapseNav();
    } else {
      showNav();
    }
  }
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

  .toggle-nav-right, .nav-right {
    display: flex;
    margin-left: auto;
  }

  .toggle-nav-right {
    display: none;
  }

  nav img {
    display: block;
    max-width: 32px;
    max-height: 32px;
    margin-right: 8px;
  }

  @media (max-width: 640px) {
    nav {
      padding: 0 10px;
    }

    .toggle-nav-right {
      display: flex;
    }

    .nav-right {
      visibility: hidden;
      flex-direction: column;
      justify-content: center;
      align-items: center;
      z-index: 2;
    }
  }
</style>

<nav>
  <Link to="/">
    <img src="/img/tor.svg" alt="Tor-esque logo (onion)" />
  </Link>

  <Link to="/">
    <h4>Run A Relay</h4>
  </Link>

  <div class="nav-right">
    <Link to="/about">
      <button class="secondary hvr-ripple-out-sec">About</button>
    </Link>

    <ExternalButton href="https://github.com/EffectiveAF/run-a-relay">
      GitHub Repo
    </ExternalButton>

    <ExternalButton primary href="https://donate.torproject.org/">
      Donate to Tor
    </ExternalButton>
  </div>

  <button
    class="toggle-nav-right"
    on:click={toggleNavRight}
    on:blur={() => {
      // TODO: Remove this race condition hack
      setTimeout(collapseNav, 250);
    }}
  >
    {togglerLabelOpen}
  </button>
</nav>
