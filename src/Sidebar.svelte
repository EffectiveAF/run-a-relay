<script>
  import { Link, navigate } from 'svelte-routing';
  import { currentStepIndex, currentPath } from './stores.js';
  import slugOrder from './slugOrder.js';

  const steps = [
    { slug: 'intro',             title: 'Getting started'  },
    { slug: 'relay-or-donate',   title: 'Run or donate?'   },
    { slug: 'choosing-a-host',   title: 'Choosing a host'  },
    { slug: 'install-tor-daemon',title: 'Install Tor'       },
    { slug: 'thank-you',         title: "You're done!"      },
  ];

  const goToStep = (slug, idx) => {
    $currentStepIndex = idx;
    currentPath.set(`/step/${slug}`);
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
    padding: 16px 0 48px;
    background: var(--bg);
    display: flex;
    flex-direction: column;
    scrollbar-width: thin;
    scrollbar-color: var(--border-strong) transparent;
  }

  aside::-webkit-scrollbar { width: 3px; }
  aside::-webkit-scrollbar-thumb { background: var(--border-strong); border-radius: 2px; }

  /* Section group */
  .group { margin-top: 4px; }
  .group:first-child { margin-top: 0; }

  /* Section header — collapsible row like Linear */
  .group-label {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 6px 12px 6px 16px;
    font-size: 0.8125rem;   /* 13px */
    font-weight: 500;
    color: var(--text-2);
    cursor: default;
    user-select: none;
  }

  /* Nav item — matches Linear's 32px-tall, 14px items */
  .item {
    display: flex;
    align-items: center;
    justify-content: flex-start;
    gap: 8px;
    width: 100%;
    height: 32px;
    padding: 0 12px 0 20px;
    font-size: 0.875rem;     /* 14px — Linear's exact size */
    font-weight: 400;
    color: var(--text-2);
    background: none;
    border: none;
    border-radius: 0;
    text-align: left;
    text-decoration: none;
    cursor: pointer;
    transition: color 0.1s, background 0.1s;
    box-shadow: none;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    position: relative;
  }

  .item:hover {
    color: var(--text);
    background: var(--bg-hover);
  }

  .item.active {
    color: var(--text);
    font-weight: 500;
    background: var(--bg-hover);
  }

  /* Accent left bar on active item */
  .item.active::before {
    content: '';
    position: absolute;
    left: 0;
    top: 6px;
    bottom: 6px;
    width: 2px;
    background: var(--accent);
    border-radius: 0 1px 1px 0;
  }

  /* Step icon — small circle with number, like Linear's page icon */
  .icon {
    flex-shrink: 0;
    width: 18px;
    height: 18px;
    border-radius: 4px;
    display: inline-flex;
    align-items: center;
    justify-content: center;
    font-size: 0.6875rem;
    font-weight: 600;
    background: var(--bg-subtle);
    color: var(--text-3);
    transition: background 0.1s, color 0.1s;
  }

  .item.active .icon {
    background: var(--accent);
    color: #fff;
  }

  /* External-link suffix */
  .ext {
    margin-left: auto;
    font-size: 0.7rem;
    color: var(--text-3);
    opacity: 0;
    transition: opacity 0.1s;
  }

  .item:hover .ext { opacity: 1; }

  /* Spacer between main groups and bottom links */
  .bottom-group {
    margin-top: auto;
    padding-top: 16px;
    border-top: 1px solid var(--border);
  }

  @media (max-width: 900px) {
    aside { display: none; }
  }
</style>

<aside>

  <!-- Overview -->
  <div class="group">
    <div class="group-label">Overview</div>

    <a class="item" class:active={isHome} href="/">
      <span class="icon">H</span>
      Home
    </a>
  </div>

  <!-- Guide steps -->
  <div class="group">
    <div class="group-label">Guide</div>

    {#each steps as step, i}
      <button
        class="item"
        class:active={isOnStep && $currentStepIndex === i}
        on:click={() => goToStep(step.slug, i)}
      >
        <span class="icon">{i + 1}</span>
        {step.title}
      </button>
    {/each}
  </div>

  <!-- More links at bottom -->
  <div class="group bottom-group">
    <a class="item" class:active={isAbout} href="/about">
      About
    </a>

    <a class="item" href="https://github.com/EffectiveAF/run-a-relay" target="_blank" rel="noopener noreferrer">
      GitHub
      <span class="ext">↗</span>
    </a>

    <a class="item" href="https://donate.torproject.org/" target="_blank" rel="noopener noreferrer">
      Donate to Tor
      <span class="ext">↗</span>
    </a>
  </div>

</aside>
