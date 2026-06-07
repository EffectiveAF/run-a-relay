<script>
  import { Link, navigate } from 'svelte-routing';
  import { currentStepIndex, currentPath } from './stores.js';
  import slugOrder from './slugOrder.js';

  const steps = [
    { slug: 'intro',            title: 'Introduction' },
    { slug: 'relay-or-donate',  title: 'Run or Donate?' },
    { slug: 'choosing-a-host',  title: 'Choosing a Host' },
    { slug: 'install-tor-daemon', title: 'Install Tor' },
    { slug: 'thank-you',        title: "You're Done!" },
  ];

  const goToStep = (slug, idx) => {
    $currentStepIndex = idx;
    navigate(`/step/${slug}`);
  };

  $: isOnStep = $currentPath.startsWith('/step/');
  $: isHome   = $currentPath === '/';
  $: isAbout  = $currentPath === '/about';
</script>

<style>
  aside {
    position: fixed;
    top: var(--nav-h);
    left: 0;
    width: var(--sidebar-w);
    height: calc(100vh - var(--nav-h));
    overflow-y: auto;
    border-right: 1px solid var(--border);
    padding: 24px 0 40px;
    display: flex;
    flex-direction: column;
    background: var(--bg);
    scrollbar-width: thin;
    scrollbar-color: var(--border-strong) transparent;
  }

  aside::-webkit-scrollbar { width: 4px; }
  aside::-webkit-scrollbar-track { background: transparent; }
  aside::-webkit-scrollbar-thumb { background: var(--border-strong); border-radius: 2px; }

  .section-label {
    display: block;
    font-size: 0.7rem;
    font-weight: 600;
    letter-spacing: 0.08em;
    text-transform: uppercase;
    color: var(--text-3);
    padding: 16px 16px 6px;
    margin-top: 8px;
  }

  .section-label:first-child { margin-top: 0; }

  nav-item, a.nav-item {
    display: flex;
    align-items: center;
    gap: 8px;
    padding: 6px 16px;
    font-size: 0.85rem;
    font-weight: 400;
    color: var(--text-2);
    text-decoration: none;
    border-radius: 0;
    cursor: pointer;
    transition: color 0.1s, background 0.1s;
    border: none;
    background: none;
    width: 100%;
    text-align: left;
    line-height: 1.4;
  }

  a.nav-item:hover, button.nav-item:hover {
    color: var(--text);
    background: var(--bg-hover);
  }

  a.nav-item.active, button.nav-item.active {
    color: var(--text);
    font-weight: 500;
    background: var(--bg-active);
    position: relative;
  }

  a.nav-item.active::before, button.nav-item.active::before {
    content: '';
    position: absolute;
    left: 0;
    top: 4px;
    bottom: 4px;
    width: 2px;
    background: var(--accent);
    border-radius: 0 2px 2px 0;
  }

  .step-number {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    width: 18px;
    height: 18px;
    border-radius: 50%;
    font-size: 0.65rem;
    font-weight: 600;
    flex-shrink: 0;
    background: var(--border-strong);
    color: var(--text-3);
  }

  .step-number.active {
    background: var(--accent);
    color: #fff;
  }

  .step-number.done {
    background: var(--bg-active);
    color: var(--accent-light);
  }

  @media (max-width: 900px) {
    aside { display: none; }
  }
</style>

<aside>
  <span class="section-label">Overview</span>

  <Link to="/">
    <a class="nav-item" class:active={isHome} href="/">
      Getting Started
    </a>
  </Link>

  <span class="section-label">Guide</span>

  {#each steps as step, i}
    <button
      class="nav-item"
      class:active={isOnStep && $currentStepIndex === i}
      on:click={() => goToStep(step.slug, i)}
    >
      <span class="step-number" class:active={isOnStep && $currentStepIndex === i} class:done={isOnStep && $currentStepIndex > i}>
        {i + 1}
      </span>
      {step.title}
    </button>
  {/each}

  <span class="section-label">More</span>

  <Link to="/about">
    <a class="nav-item" class:active={isAbout} href="/about">About</a>
  </Link>

  <a class="nav-item" href="https://github.com/EffectiveAF/run-a-relay" target="_blank" rel="noopener noreferrer">
    GitHub ↗
  </a>

  <a class="nav-item" href="https://donate.torproject.org/" target="_blank" rel="noopener noreferrer">
    Donate to Tor ↗
  </a>
</aside>
